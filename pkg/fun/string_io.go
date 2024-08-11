package fun

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type StringIO interface {
	WriteToString(text string)
	ReadFromString(text *string)
}

type stringIO struct {
	reader io.Reader
	writer io.Writer
}

func NewStringIO(reader io.Reader, writer io.Writer) StringIO {
	return &stringIO{
		reader,
		writer,
	}
}

// ReadFromString
func (s *stringIO) ReadFromString(text *string) {
	reader := bufio.NewReader(s.reader)
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading from reader %v", err)
		*text = ""
	}

	*text = strings.Trim(content, "\n")
}

// WriteToString
func (s *stringIO) WriteToString(text string) {
	text = fmt.Sprintf("\n%s\n", text)
	if _, err := fmt.Fprint(s.writer, text); err != nil {
		fmt.Printf("Error writing to writer %v", err)
	}
}
