package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      string    `json:"likes,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.formatter()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("o título é obrigatório e não pode estar vazio")
	}

	if publication.Content == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar vazio")
	}

	return nil
}

func (publication *Publication) formatter() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
