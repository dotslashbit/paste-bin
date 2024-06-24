package data

import (
	"database/sql"
	"errors"
	"time"

	"dev.dotslashbit.paste-bin/internal/validator"
)

type Snippet struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	ExpireAt time.Time `json:"expire_at"`
}

func ValidateSnippet(v *validator.Validator, snippet *Snippet) {
	v.Check(snippet.Title != "", "title", "must be provided")
	v.Check(len(snippet.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(snippet.Content != "", "content", "must be provided")

}

type SnippetModel struct {
	DB *sql.DB
}

func (m SnippetModel) Insert(snippet *Snippet) error {
	query :=
		`
			INSERT INTO snippets (title, content)
			VALUES($1, $2)
			RETURNING id,  expires_at
			`
	args := []interface{}{snippet.Title, snippet.Content}
	return m.DB.QueryRow(query, args...).Scan(&snippet.Id, &snippet.ExpireAt)
}

func (m SnippetModel) Get(id int) (*Snippet, error) {
	query :=
		`
			SELECT id, title, content, expires_at
			FROM snippets
			WHERE expires_at > now() AND id = $1
			`
	var snippet Snippet
	err := m.DB.QueryRow(query, id).Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.ExpireAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &snippet, nil
}

func (m SnippetModel) Update(snippet *Snippet) error {
	query := `
		UPDATE snippets
		SET title = $1, content = $2, expires_at = $3
		WHERE id = $4
		RETURNING id, expires_at
	`
	args := []interface{}{snippet.Title, snippet.Content, snippet.ExpireAt, snippet.Id}
	return m.DB.QueryRow(query, args...).Scan(&snippet.Id, &snippet.ExpireAt)
}

func (m SnippetModel) Delete(id int) error {
	query := `
		DELETE FROM snippets
		WHERE id = $1
		`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
