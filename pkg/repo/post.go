package repo

import (
	"fmt"
	m "simple-golang-tui/pkg/model"
)

type PostIO interface {
	GetAll() *[]*m.Post
	GetOne(index uint) *m.Post
	Add(post *m.Post)
	AddAll(posts *[]m.Post)
	Remove(id uint)
	Edit(id uint, title, body string)
}

type postIO struct {
	posts []*m.Post
}

func NewPostIO() PostIO {
	return &postIO{
		posts: []*m.Post{},
	}
}

// Add
func (p *postIO) Add(post *m.Post) {
	p.posts = append(p.posts, post)
}

// AddAll
func (p *postIO) AddAll(posts *[]m.Post) {
	list := *posts
	for i := range list {
		post := list[i]
		p.posts = append(p.posts, &post)
	}
}

// Edit
func (p *postIO) Edit(id uint, title string, body string) {
	for _, post := range p.posts {
		if post.ID == id {
			if len(title) != 0 {
				post.Title = title
			}

			if len(body) != 0 {
				post.Body = body
			}

			return
		}
	}

	fmt.Printf("Post with id %d, not found\n", id)
}

// GetAll
func (p *postIO) GetAll() *[]*m.Post {
	return &p.posts
}

// GetOne
func (p *postIO) GetOne(index uint) *m.Post {
	//  index>=0 && int(index)<len(p.posts)
	if int(index) < len(p.posts) {
		return p.posts[index]
	}

	fmt.Printf("Post not found with index %d", index)
	return nil
}

// Remove
func (p *postIO) Remove(id uint) {
	for i, post := range p.posts {
		if post.ID == id {
			p.posts = append(p.posts[:i], p.posts[i+1:]...)
			return
		}
	}

	fmt.Printf("Post with id %d, not found\n", id)
}
