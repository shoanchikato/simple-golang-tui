package main

import (
	"os"
	f "simple-golang-tui/pkg/fun"
	r "simple-golang-tui/pkg/repo"
	s "simple-golang-tui/pkg/store"
)

func main() {
	app := di()

	app.Run()
}

func di() f.App {
	stdin := os.Stdin
	stdout := os.Stdout

	stringIO := f.NewStringIO(stdin, stdout)
	userIO := f.NewUserIO(stringIO)

	fileIO := s.NewFileIO()
	postIO := r.NewPostIO()
	appIO := s.NewAppIO(fileIO, postIO)

	postOptions := f.NewPostOptions(postIO, userIO)
	app := f.NewApp(appIO, postOptions, userIO)

	return app
}
