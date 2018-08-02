package domains

import (
	"time"
)

type Note struct {
	Id               int
	UserId           int
	TagId            int
	NoteTypeId       int
	Title            string
	CreatedAt        time
	UpdatedAt        time
	OriginalLink     string
	Author           string
	OriginalSrouce   string
	Content          string
	ShortDescription string
	IsDeleted        bool
	IsFavorite       bool
}

type NoteInterface struct {
	
}
