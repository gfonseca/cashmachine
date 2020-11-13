package test

import (
	"fmt"
	"io/ioutil"
	"os"
)

// MockFile represents a file with arbirtrary content for test porpouses
type MockFile struct {
	path       string
	descriptor *os.File
}

// NewMockFile mock a file for test porpouses
func NewMockFile(filepath, content string) *MockFile {
	descriptor, err := os.Create(filepath)

	if err != nil {
		panic(fmt.Errorf("[Error]: %s", err))
	}

	f := &MockFile{path: filepath, descriptor: descriptor}
	f.descriptor.WriteString(content)
	return f
}

// GetContent return mocked file content
func (f MockFile) GetContent() string {
	data, err := ioutil.ReadFile(f.path)
	if err != nil {
		panic(fmt.Errorf("[Error]: %s", err))
	}

	return string(data)
}

// ClearFile remove the file from file system
func (f MockFile) ClearFile() {
	os.Remove(f.path)
}
