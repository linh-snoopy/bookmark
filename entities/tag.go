package entities

import (
	"time"
)

type Tag struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
