package utils

import (
	"fmt"
	"os"
)

type FileUtils struct {
}

type StoreFileInterface interface {
	ReadFile(filePath string) (string, error)
	WriteFile(filePath string, data string) error
	CheckFile(filePath string) bool
	CreateFile(path string) error
}

// CreateFile creates a new file if it does not already exist
func (utils *FileUtils) CreateFile(path string) error {
	// Check if the file already exists
	if _, err := os.Stat(path); err == nil {
		fmt.Println("File already exists.")
		return nil
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// fmt.Println("File created successfully:", path)
	return nil
}

// WriteFile writes content to a file, replacing existing content if any
func (utils *FileUtils) WriteFile(path, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	// fmt.Println("Successfully wrote to file:", path)
	return nil
}

// ReadFile reads the contents of a file and returns it as a string
func (utils *FileUtils) ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	// fmt.Println("File contents:", string(data))
	return string(data), nil
}

// CheckFile checks if a file exists at the given path
func (utils *FileUtils) CheckFile(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// fmt.Println("File found:", path)
		return true
	} else if os.IsNotExist(err) {
		// fmt.Println("File not found:", path)
		return false
	} else {
		// fmt.Println("Error checking file:", err)
		return false
	}
}
