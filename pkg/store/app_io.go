package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	m "simple-golang-tui/pkg/model"
	r "simple-golang-tui/pkg/repo"
)

const FILE_NAME = "post_db.json"

type AppIO interface {
	OnLoad()
	OnEnd()
}

type appIO struct {
	fileIO FileIO
	postIO r.PostIO
}

func NewAppIO(fileIO FileIO, postIO r.PostIO) AppIO {
	return &appIO{
		fileIO,
		postIO,
	}
}

// OnEnd
func (a *appIO) OnEnd() {
	filePath := FILE_NAME

	writePosts := func() {
		posts := a.postIO.GetAll()

		bytes, err := json.Marshal(posts)
		if err != nil {
			fmt.Printf("Error serializing posts %v \n", err)
			return
		}

		stringPosts := string(bytes)
		a.fileIO.WriteFile(filePath, stringPosts)
	}

	if err := a.fileIO.CreateFile(filePath); err != nil {
		fmt.Printf("Error creating a new file %v \n", err)
		return
	}

	writePosts()
}

// OnLoad
func (a *appIO) OnLoad() {
	filePath := FILE_NAME

	loadPosts := func() {
		stringPosts := a.fileIO.ReadFile(filePath)

		posts := []m.Post{}

		err := json.Unmarshal([]byte(stringPosts), &posts)
		if err != nil {
			fmt.Printf("Error deserializing posts %v\n", err)
			return
		}

		a.postIO.AddAll(&posts)
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		err := a.fileIO.CreateFile(filePath)
		if err != nil {
			fmt.Printf("Error create a new file %v\n", err)
			return
		}
	}

	loadPosts()
}
