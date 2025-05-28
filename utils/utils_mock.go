package utils

import "github.com/stretchr/testify/mock"

type MockFileUtils struct {
	mock.Mock
}

func (utils *MockFileUtils) CreateFile(path string) error {
	args := utils.Called(path)
	return args.Error(0)
}
func (utils *MockFileUtils) WriteFile(path, content string) error {
	args := utils.Called(path, content)
	return args.Error(0)
}
func (utils *MockFileUtils) ReadFile(path string) (string, error) {
	args := utils.Called(path)
	return args.String(0), args.Error(1)
}
func (utils *MockFileUtils) CheckFile(path string) bool {
	args := utils.Called(path)
	return args.Bool(0)
}
