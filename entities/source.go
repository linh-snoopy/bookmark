package entities

import (
	"time"
)

type Source struct {
	Id         int       `json:"id"`
	SiteName   string    `json:"site_name"` // WordPress(https://example.wordpress.com/feed), Blogger(http://blogname.blogspot.com/feeds/posts/default), Medium.com(medium.com/feed/example-site), (Tumblr: http://example.tumblr.com/rss)
	FeedName   string    `json:"feed_name"`
	FeedFormat string    `json:"feed_format"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SourceRepository interface {
	InsertSource(source Source) error
	UpdateSource(source Source) error
}
