package entities

import (
	"time"
)

type NoteType struct {
	Id        int       `json:"id"`
	Name      int       `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// articles
// videos
// images

type NoteTypeRepository interface {
	InsertNoteType(name string) error
}
