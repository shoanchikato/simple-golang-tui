package fun

import (
	"strings"
)

type UserIO interface {
	GetResponse(question string) string
}

type userIO struct {
	stringIO StringIO
}

func NewUserIO(stringIO StringIO) UserIO {
	return &userIO{
		stringIO,
	}
}

// GetResponse
func (u *userIO) GetResponse(question string) string {
	input := ""
	u.stringIO.WriteToString(question)
	u.stringIO.ReadFromString(&input)

	return strings.Trim(input, "")
}
