package models

//UploadMedia struct
type UploadMedia struct {
	MediaID    int64  `json:"media_id"`
	MediaIDStr string `json:"media_id_string"`
	Size       int64  `json:"size"`
	Expires    int64  `json:"expires_after_secs"`
}

//Image struct
type Image struct {
	ImageType string `json:"image_type"`
	Width     int    `json:"w"`
	Height    int    `json:"h"`
}
