package fun

import (
	"fmt"
	m "simple-golang-tui/pkg/model"
	r "simple-golang-tui/pkg/repo"
	"strconv"
	"strings"
)

type PostOptions interface {
	ShowPosts()
	WritePost()
	EditPost()
	RemovePost()
	GetAll() *[]*m.Post
	AddAll(posts *[]m.Post)
}

type postOptions struct {
	repo   r.PostIO
	userIO UserIO
}

func NewPostOptions(repo r.PostIO, userIO UserIO) PostOptions {
	return &postOptions{
		repo,
		userIO,
	}
}

// AddAll
func (p *postOptions) AddAll(posts *[]m.Post) {
	p.repo.AddAll(posts)
}

// EditPost
func (p *postOptions) EditPost() {
	input := p.userIO.GetResponse("Which post do you want to edit?")

	id, err := strconv.Atoi(strings.Trim(input, ""))
	if err != nil {
		fmt.Printf("Invalid id %s \n", input)
		return
	}

	message := fmt.Sprintf(
		"%s\n%s\n%s\n",
		"What do you want to edit?", "title", "body",
	)

	input = p.userIO.GetResponse(message)

	if input == "title" {
		input := p.userIO.GetResponse("Enter the new title")
		p.repo.Edit(uint(id), input, "")
	} else if input == "body" {
		input := p.userIO.GetResponse("Enter the new body")
		p.repo.Edit(uint(id), "", input)
	}
}

// GetAll
func (p *postOptions) GetAll() *[]*m.Post {
	return p.repo.GetAll()
}

// RemovePost
func (p *postOptions) RemovePost() {
	input := p.userIO.GetResponse("Which post do you want to remove?")

	id, err := strconv.Atoi(strings.Trim(input, ""))
	if err != nil {
		fmt.Printf("Invalid id %s \n", input)
		return
	}

	p.repo.Remove(uint(id))
}

// ShowPosts
func (p *postOptions) ShowPosts() {
	fmt.Println("\nPOSTS:")
	fmt.Printf("======\n\n")
	for _, post := range *p.GetAll() {
		fmt.Print(post)
	}
}

// WritePost
func (p *postOptions) WritePost() {
	questions := []string{"What's the title?", "What's the body?"}
	answers := []string{}

	for _, question := range questions {
		input := p.userIO.GetResponse(question)
		answers = append(answers, input)
	}

	appendNewPost := func(id uint) {
		post := m.NewPost(id, answers[0], answers[1])
		p.repo.Add(&post)
	}

	posts := *p.repo.GetAll()
	if len(posts) > 0 {
		id := posts[len(posts)-1].ID + 1
		appendNewPost(id)
	} else {
		appendNewPost(1)
	}
}
