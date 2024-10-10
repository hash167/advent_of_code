package aoc2023

import (
	"testing"
)

func TestGetFirstDigit(t *testing.T) {
	tests := []struct {
		input    string
		partB    bool
		expected string
		err      bool
	}{
		{"abc123", false, "1", false},
		{"one23", true, "1", false},
		{"onetwothree", true, "1", false},
		{"no digits here", false, "", true},
		{"onetwothree", false, "", true},
	}

	for _, tt := range tests {
		result, err := getFirstDigit([]rune(tt.input), tt.partB)
		if (err != nil) != tt.err {
			t.Errorf("getFirstDigit(%q, %v) error = %v, wantErr %v", tt.input, tt.partB, err, tt.err)
			continue
		}
		if result != tt.expected {
			t.Errorf("getFirstDigit(%q, %v) = %v, want %v", tt.input, tt.partB, result, tt.expected)
		}
	}
}

func TestGetLastDigit(t *testing.T) {
	tests := []struct {
		input    string
		partB    bool
		expected string
		err      bool
	}{
		{"abc123", false, "3", false},
		{"123one", true, "1", false},
		{"onetwothree", true, "3", false},
		{"no digits here", false, "", true},
		{"onetwothree", false, "", true},
	}

	for _, tt := range tests {
		result, err := getLastDigit([]rune(tt.input), tt.partB)
		if (err != nil) != tt.err {
			t.Errorf("getLastDigit(%q, %v) error = %v, wantErr %v", tt.input, tt.partB, err, tt.err)
			continue
		}
		if result != tt.expected {
			t.Errorf("getLastDigit(%q, %v) = %v, want %v", tt.input, tt.partB, result, tt.expected)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		input    string
		partB    bool
		expected int
		err      bool
	}{
		{"abc123", false, 13, false},
		{"one23", true, 13, false},
		{"onetwothree", true, 13, false},
		{"no digits here", false, 0, true},
		{"onetwothree", false, 0, true},
	}

	for _, tt := range tests {
		result, err := solve(tt.input, tt.partB)
		if (err != nil) != tt.err {
			t.Errorf("solve(%q, %v) error = %v, wantErr %v", tt.input, tt.partB, err, tt.err)
			continue
		}
		if result != tt.expected {
			t.Errorf("solve(%q, %v) = %v, want %v", tt.input, tt.partB, result, tt.expected)
		}
	}
}
