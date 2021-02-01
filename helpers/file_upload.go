package helpers

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

//UploadMedia to upload file using Form-Data
func UploadMedia(file []byte, filename string, paramName string, uri string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filename)
	if err != nil {
		fmt.Println(err)
	}
	part.Write(file)
	defer writer.Close()
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return &http.Request{}, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}
