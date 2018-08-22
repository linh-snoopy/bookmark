package entities

import (
	"time"
)

type Page struct {
	Id          int       `json:"id"`
	SourceId    int       `json:"source_id"`
	BaseUrl     string    `json:"base_url"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	OriginalUrl string    `json:"original_url"`
	Author      string    `json:"author"`
	Summary     string    `json:"summary"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PageRepository interface {
	InsertPage(page Page) error
	UpdatePage(page Page) error
}