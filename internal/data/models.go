package data

import (
	"errors"

	"database/sql"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Snippets SnippetModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Snippets: SnippetModel{DB: db},
	}
}
