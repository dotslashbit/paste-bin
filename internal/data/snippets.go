package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"dev.dotslashbit.paste-bin/internal/validator"
)

type Snippet struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Format    string `json:"format"`
	ExpiresAt string `json:"expires_at"`
}

func ValidateSnippet(v *validator.Validator, snippet *Snippet) {
	v.Check(snippet.Title != "", "title", "must be provided")
	v.Check(len(snippet.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(snippet.Content != "", "content", "must be provided")
	v.Check(snippet.ExpiresAt != "", "expires_at", "must be provided")
	v.Check(snippet.Format != "", "format", "must be provided")

}

type SnippetModel struct {
	DB *sql.DB
}

func (m SnippetModel) Insert(snippet *Snippet) error {
	query :=
		`
            INSERT INTO snippets (title, content, format, expires_at)
            VALUES($1, $2, $3, $4)
            RETURNING id, format, expires_at
        `
	args := []interface{}{snippet.Title, snippet.Content, snippet.Format, snippet.ExpiresAt}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&snippet.Id, &snippet.Format, &snippet.ExpiresAt)
}

func (m SnippetModel) Get(id int) (*Snippet, error) {
	query :=
		`
			SELECT id, title, content, format, expires_at
			FROM snippets
			WHERE expires_at > now() AND id = $1
			`
	var snippet Snippet
	err := m.DB.QueryRow(query, id).Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.Format, &snippet.ExpiresAt)
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
        SET title = $1, content = $2, format = $3, expires_at = $4
        WHERE id = $5
        RETURNING id, format, expires_at
    `
	args := []interface{}{snippet.Title, snippet.Content, snippet.Format, snippet.ExpiresAt, snippet.Id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&snippet.Id, &snippet.Format, &snippet.ExpiresAt)
}

func (m SnippetModel) Delete(id int) error {
	query := `
		DELETE FROM snippets
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := m.DB.ExecContext(ctx, query, id)
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
