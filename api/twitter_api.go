package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	cfg "github.com/twitter-bot-golang/config"
	"github.com/twitter-bot-golang/helpers"
	"github.com/twitter-bot-golang/models"
)

var apiKey = cfg.GetString("token.consumer_api_key")
var apiSecret = cfg.GetString("token.consumer_secret_key")
var accessToken = cfg.GetString("token.access_token_key")
var accessSecret = cfg.GetString("token.access_token_secret_key")

var config = oauth1.NewConfig(apiKey, apiSecret)
var token = oauth1.NewToken(accessToken, accessSecret)
var httpClient = config.Client(oauth1.NoContext, token)
var client = twitter.NewClient(httpClient)

var timeSinceReadyToTweet int64
var wg sync.WaitGroup

//CheckActivity to check activity received from webhook
func CheckActivity(activity *models.ActivityEvent) {
	if len(activity.DMEvents) > 0 {
		err := handleDM(activity)
		if err != nil {
			panic("Error Handling DM")
		}
		return
	}
}

func handleDM(activity *models.ActivityEvent) error {
	DM := activity.DMEvents[0]
	text := DM.MessageCreate.MessageData.Text
	sender := DM.MessageCreate.SenderID

	if sender == BotID {
		return nil
	}

	if strings.Contains(text, HelpTrigger) {
		_, _, err := client.DirectMessages.EventsNew(createDMReply(sender, HelpMessage))
		if err != nil {
			return errors.New("Failed to send help DM")
		}
	}

	if strings.Contains(text, SendTrigger) {
		_, _, err := client.DirectMessages.EventsNew(createDMReply(sender, AcceptedMessage))
		if err != nil {
			return errors.New("Failed to send help DM")
		}

		statusParams := twitter.StatusUpdateParams{}
		statusParams.Status = text

		attachmentID := getAttachmentFile(DM)
		if attachmentID > 0 {
			statusParams.Status = helpers.RemoveURLFromText(statusParams.Status)
			statusParams.MediaIds = []int64{attachmentID}
		}
		if isTimeToTweet() {
			doSendTweet(sender, 0, &statusParams)
		} else {
			wg.Add(1)
			go doSendTweet(sender, 60, &statusParams)
			wg.Wait()
		}
		if r := recover(); r != nil {
			_, _, err := client.DirectMessages.EventsNew(createDMReply(sender, FailedMessage))
			if err != nil {
				return errors.New("Failed to send DM")
			}
		}
		//Update timeSinceLastTweet
		timeSinceReadyToTweet = time.Now().Unix()
	}

	return nil
}

func getAttachmentFile(dm models.DMEvent) int64 {
	var err error
	//Check if dm event contains attachment
	attachment := dm.MessageCreate.MessageData.Attachment

	if attachment.Type == "" {
		return 0
	}

	//Get the attachment URL and request for the file
	path := attachment.Media.MediaURLHTTPS
	resp, err := httpClient.Get(path)
	if err != nil {
		return 0
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	request, err := helpers.UploadMedia(body, "filename", MediaParam, TwitterMediaUploadURL)
	if err != nil {
		return 0
	}

	uploadImageResponse, err := httpClient.Do(request)
	if err != nil {
		return 0
	}

	uploadMediaResponse, err := ioutil.ReadAll(uploadImageResponse.Body)
	var uploadMedia models.UploadMedia
	if err != nil {
		return 0
	}

	err = json.Unmarshal([]byte(uploadMediaResponse), &uploadMedia)
	if err != nil {
		return 0
	}
	return uploadMedia.MediaID
}

func doSendTweet(recipient string, waitTimeSecond int, params *twitter.StatusUpdateParams) {
	if waitTimeSecond > 0 {
		defer wg.Done()
	}

	//Sleep if bot is not ready to tweet
	time.Sleep(time.Second * time.Duration(waitTimeSecond))

	//Update status with message
	_, _, err := client.Statuses.Update(params.Status, params)
	if err != nil {
		panic("Failed to update status")
	}

	//Notify user that message has been sent
	_, _, err2 := client.DirectMessages.EventsNew(createDMReply(recipient, "Berhasil mengirim pesan!"))
	if err2 != nil {
		panic("Failed to send DM")
	}
}

func createDMReply(recipient string, message string) *twitter.DirectMessageEventsNewParams {
	return &twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Type: "message_create",
			Message: &twitter.DirectMessageEventMessage{
				Target: &twitter.DirectMessageTarget{
					RecipientID: recipient,
				},
				Data: &twitter.DirectMessageData{
					Text: message,
				},
			},
		},
	}
}

func isTimeToTweet() bool {
	currentTime := time.Now().Unix()
	if currentTime-timeSinceReadyToTweet < 60 {
		return false
	}
	return true
}
