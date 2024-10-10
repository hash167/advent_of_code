package utils

import "testing"

func TestNewLineIterator(t *testing.T) {
	it, err := NewLineIterator("../inputs/input1.txt")
	if err != nil {
		t.Error(err)
	}
	for {
		line, ok := it.Next()
		if !ok {
			break
		}
		println(line)
	}
}
