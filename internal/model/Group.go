package model

import "time"

type Group struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatorID int       `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
	Members   []int     `json:"members"`
}
