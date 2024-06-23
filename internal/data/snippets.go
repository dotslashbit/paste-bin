package data

import (
	"time"

	"dev.dotslashbit.paste-bin/internal/validator"
)

type Snippet struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	ExpireAt time.Time `json:"expire_at"`
}

func ValidateSnippet(v *validator.Validator, snippet *Snippet) {
	v.Check(snippet.Title != "", "title", "must be provided")
	v.Check(len(snippet.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(snippet.Content != "", "content", "must be provided")

}
