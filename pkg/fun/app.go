package fun

import (
	"fmt"
	s "simple-golang-tui/pkg/store"
	"strconv"
)

type App interface {
	showOptions() int
	clearScreen()
	Run()
}

type app struct {
	appIO       s.AppIO
	postOptions PostOptions
	userIO      UserIO
}

func NewApp(appIO s.AppIO, postOptions PostOptions, userIO UserIO) App {
	return &app{
		appIO,
		postOptions,
		userIO,
	}
}

// Run
func (a *app) Run() {
	a.appIO.OnLoad()

out:
	for {
		option := a.showOptions()

		switch option {
		case 0:
			break out
		case 1:
			a.postOptions.ShowPosts()
		case 2:
			a.postOptions.WritePost()
		case 3:
			a.clearScreen()
		case 4:
			a.postOptions.EditPost()
		case 5:
			a.postOptions.RemovePost()
		default:
			continue out
		}
	}

	a.appIO.OnEnd()
}

// clearScreen
func (a *app) clearScreen() {
	fmt.Printf("\x1B[2J\x1B[H")
}

// showOptions
func (a *app) showOptions() int {
	options := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"Choose an option:",
		"1 -> show posts",
		"2 -> add post",
		"3 -> clear",
		"4 -> edit",
		"5 -> remove",
		"0 -> exit",
	)

	input := a.userIO.GetResponse(options)

	option, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error parsing option %v \n", err)
		return 3
	}

	return option
}
