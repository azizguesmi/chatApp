package model

import "time"

type Message struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	CreatedAt  time.Time `json:"created_at"`
	Rec_type   string    `json:"rec_type"`
}
