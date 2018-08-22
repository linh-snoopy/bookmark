package entities

import (
	"time"
)

type Note struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	NoteTypeId int       `json:"note_type_id"`
	PageId     int       `json:"page_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsDeleted  bool      `json:"is_deleted"`
	IsFavorite bool      `json:"is_favorite"`
	Tags       []string  `json:"tags"`
}

type NoteInterface struct {
}
