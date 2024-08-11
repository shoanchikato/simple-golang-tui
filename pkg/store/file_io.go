package store

import (
	"fmt"
	"os"
)

type FileIO interface {
	CreateFile(fileName string) error
	WriteFile(fileName, contents string)
	ReadFile(fileName string) string
}

type fileIO struct{}

func NewFileIO() FileIO {
	return &fileIO{}
}

// CreateFile
func (f *fileIO) CreateFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create %s file %v\n", fileName, err)
		return err
	}

	defer file.Close()
	return nil
}

// WriteFile
func (f *fileIO) WriteFile(fileName string, contents string) {
	err := os.WriteFile(fileName, []byte(contents), 0666)
	if err != nil {
		fmt.Printf("Failed to write %s file %v\n", fileName, err)
		return
	}
}

// ReadFile
func (f *fileIO) ReadFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Failed to read %s file %v\n", fileName, err)
		return ""
	}

	fileText := string(bytes[:])

	return fileText
}
