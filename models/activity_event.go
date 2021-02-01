package models

//ActivityEvent struct
type ActivityEvent struct {
	ForUserID string    `json:"for_user_id"`
	DMEvents  []DMEvent `json:"direct_message_events"`
}
