package domains

import (
	"time"
)

type User struct {
	Id        int
	Email     string
	Password  string
	CreatedAt time
}
