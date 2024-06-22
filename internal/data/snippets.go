package data

import "time"

type Snippet struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	ExpireAt time.Time `json:"expire_at"`
}
