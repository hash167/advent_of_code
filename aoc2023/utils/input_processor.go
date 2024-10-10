package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

// Write a funtion that takes a file path and returns a slice of strings
func ReadFileLines(path string) ([]string, error) {
	// Get absolute path
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Lets use an iterator for more efficiency

type LineIterator struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewLineIterator(path string) (*LineIterator, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return &LineIterator{file: file, scanner: scanner}, nil
}

func (li *LineIterator) Next() (string, bool) {
	if !li.scanner.Scan() {
		li.file.Close()
		return "", false
	}
	return li.scanner.Text(), true
}
