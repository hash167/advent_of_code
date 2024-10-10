package main

import (
	"flag"
	"fmt"
	"strconv"

	"algos/aoc2023"
)

func main() {
	// Define a command-line flag for the AOC day number
	day := flag.Int("day", 0, "Advent of Code day number to solve (e.g., -day=1)")
	sample := flag.Bool("sample", false, "Run the sample input for the day")
	part := flag.String("part", "A", "Advent of Code part number to solve (e.g., -part=A)")

	// Parse command-line flags
	flag.Parse()

	// Print usage and exit if the day flag is not provided or is invalid
	if *day < 1 {
		fmt.Println("Usage: go run main.go -day=<day_number>")
		fmt.Println("Provide a valid AOC day number as argument (e.g., -day=1)")
		return
	}
	var input string
	if *sample {
		input = "./aoc2023/inputs/sample" + ".txt"
	} else {
		input = "./aoc2023/inputs/input" + strconv.Itoa(*day) + ".txt"
	}

	// Handle the specified day number
	switch *day {
	case 1:
		aoc2023.SolveDay1(input)
	case 2:
		aoc2023.SolveDay2(input)
	case 3:
		aoc2023.SolveDay3(input)
	case 4:
		aoc2023.SolveDay4(input)
	case 5:
		if *part == "A" {
			aoc2023.SolveDay5(input)
		} else {
			aoc2023.SolveDay5b(input)
		}
	default:
		fmt.Printf("Day %d not implemented yet.\n", *day)
	}
}
