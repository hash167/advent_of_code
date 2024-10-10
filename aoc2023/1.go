package aoc2023

import (
	"algos/aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Map and slice for number word to digit conversion.
var wordToNumberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// SolveDay1 processes the input for day 1, calculating and printing the sums for part 1 and part 2.
func SolveDay1(input string) {
	it, err := utils.NewLineIterator(input)
	if err != nil {
		panic(err)
	}

	var sum1a, sum1b int
	for {
		line, ok := it.Next()
		if !ok {
			break
		}
		res1a, err := solve(line, false)
		if err != nil {
			fmt.Println(err)
		}
		sum1a += res1a

		res1b, err := solve(line, true)
		if err != nil {
			fmt.Println(err)
		}
		sum1b += res1b
	}

	fmt.Printf("Sum for part 1 is: %d\n", sum1a)
	fmt.Printf("Sum for part 2 is: %d\n", sum1b)
}

// getFirstDigit finds the first digit or number word in the given rune slice.
func getFirstDigit(runeSlice []rune, partB bool) (string, error) {
	for index := 0; index < len(runeSlice); index++ {
		if unicode.IsDigit(runeSlice[index]) {
			return string(runeSlice[index]), nil
		}
		if partB {
			for _, word := range words {
				if strings.HasPrefix(string(runeSlice[index:]), word) {
					if num, ok := wordToNumberMap[word]; ok {
						return num, nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf("no digit or number word found")
}

// getLastDigit finds the last digit or number word in the given rune slice.
func getLastDigit(runeSlice []rune, partB bool) (string, error) {
	for index := len(runeSlice) - 1; index >= 0; index-- {
		if unicode.IsDigit(runeSlice[index]) {
			return string(runeSlice[index]), nil
		}
		if partB {
			for _, word := range words {
				if strings.HasPrefix(string(runeSlice[index:]), word) {
					if num, ok := wordToNumberMap[word]; ok {
						return num, nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf("no digit or number word found")
}

// solveDay1b processes a single line for part 2, extracting and summing the first and last digits or number words.
func solve(line string, partB bool) (int, error) {
	r := []rune(line)
	firstDigit, err := getFirstDigit(r, partB)
	if err != nil {
		return 0, err
	}
	lastDigit, err := getLastDigit(r, partB)
	if err != nil {
		return 0, err
	}
	res, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		return 0, err
	}
	return res, nil
}
