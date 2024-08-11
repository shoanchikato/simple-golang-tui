package model

import "fmt"

type Post struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewPost(id uint, title, body string) Post {
	return Post{
		id,
		title,
		body,
	}
}

func (p Post) String() string {
	return fmt.Sprintf(
		"Post %d\ntitle: %s\nbody: %s\n\n",
		p.ID, p.Title, p.Body,
	)
}
