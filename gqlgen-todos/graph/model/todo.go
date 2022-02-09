package model

type Todo struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
}
