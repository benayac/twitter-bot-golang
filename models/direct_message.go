package models

//DMEvent struct
type DMEvent struct {
	Type          string        `json:"type"`
	ID            string        `json:"id"`
	Timestamp     string        `json:"created_timestamp"`
	MessageCreate MessageCreate `json:"message_create"`
}

//MessageCreate struct
type MessageCreate struct {
	Target      Target        `json:"target"`
	SenderID    string        `json:"sender_id"`
	MessageData DMMessageData `json:"message_data"`
}

//Target struct
type Target struct {
	RecipientID string `json:"recipient_id"`
}

//DMMessageData struct
type DMMessageData struct {
	Text       string     `json:"text"`
	Attachment Attachment `json:"attachment"`
}

//Attachment struct
type Attachment struct {
	Type  string `json:"type"`
	Media Media  `json:"media"`
}

//Media struct
type Media struct {
	ID            uint64 `json:"id"`
	IDStr         string `json:"id_str"`
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	URL           string `json:"url"`
	DisplayURL    string `json:"display_url"`
	ExpandedURL   string `json:"expanded_url"`
	Type          string `json:"type"`
}
