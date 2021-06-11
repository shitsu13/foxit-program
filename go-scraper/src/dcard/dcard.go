package dcard

import (
	"time"
)

const (
	DCARD_URL = "www.dcard.tw"
	FORUM_API = "service/api/v2/forums/%s/posts"
	POST_API  = "service/api/v2/posts/%d"
)

type Posts struct {
	Post
	Categories []string  `json:"categories"`
	Topics     []string  `json:"topics"`
	Media      []Media   `json:"media"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Post struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Media struct {
	URL string `json:"url"`
}

/*
	new Posts
*/
func NewPosts() *Posts {
	p := new(Posts)
	p.Categories = make([]string, 0)
	p.Topics = make([]string, 0)
	p.Media = make([]Media, 0)

	return p
}
