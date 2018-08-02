package domains

import (
	"time"
)

type Tag struct {
	Id        int
	UserId    int
	Name      string
	CreatedAt time
}
