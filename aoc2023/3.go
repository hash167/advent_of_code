package aoc2023

import (
	"algos/aoc2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

func SolveDay3(input string) {
	var grid [][]rune
	symbolIndices := make(map[rune][][]int)

	iterator, err := utils.NewLineIterator(input)
	if err != nil {
		panic(err)
	}
	rowIndex := 0
	for {
		line, ok := iterator.Next()
		if !ok {
			break
		}
		row := []rune(line)
		grid = append(grid, row)

		// Store the indices of symbols in the line
		for colIndex, char := range row {
			if char != '.' && !unicode.IsDigit(char) {
				symbolIndices[char] = append(symbolIndices[char], []int{rowIndex, colIndex})
			}
		}
		rowIndex++

	}
	// printSymbolIndices(symbolIndices)
	var res []int
	var isGear bool
	gearSum := 0
	for symbol, indices := range symbolIndices {
		if symbol == '*' {
			isGear = true
		} else {
			isGear = false
		}
		for _, index := range indices {
			gearProduct := 0
			seenNumbers := getAdjacentNumbers(grid, index[0], index[1])
			var addNums []int
			for numString := range seenNumbers {
				num, _ := strconv.Atoi(numString)
				addNums = append(addNums, num)
			}
			if isGear {
				if len(addNums) == 1 && seenNumbers[strconv.Itoa(addNums[0])] == 2 {
					gearProduct = addNums[0] * addNums[0]
				}
				if len(addNums) == 2 && seenNumbers[strconv.Itoa(addNums[0])] == 1 && seenNumbers[strconv.Itoa(addNums[1])] == 1 {
					gearProduct = addNums[0] * addNums[1]
				}
			}
			res = append(res, addNums...)
			gearSum += gearProduct
		}

	}
	sum := func(nums []int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}(res)
	fmt.Println("The sum for part A is", sum)
	fmt.Println("The sum for part B is", gearSum)
}

func capturedNumber(ni, nj int, grid [][]rune) string {
	start, end := nj, nj
	// Move left to capture the start of the number
	for start > 0 && unicode.IsDigit(grid[ni][start-1]) {
		start--
	}
	// Move right to capture the end of the number
	for end < len(grid[ni])-1 && unicode.IsDigit(grid[ni][end+1]) {
		end++
	}
	return string(grid[ni][start : end+1])
}

// getAdjacentCells returns the locations of all adjacent cells, including diagonals, for the given index (i, j).
func getAdjacentNumbers(grid [][]rune, i, j int) map[string]int {
	adjacentPositions := [][]int{
		{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1},
		{i, j - 1} /*   (i,j)   */, {i, j + 1},
		{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1},
	}

	seenNumbers := map[string]int{}
	// Filter out positions that are out of bounds
	for _, pos := range adjacentPositions {
		ni, nj := pos[0], pos[1]
		if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
			if unicode.IsDigit(grid[ni][nj]) {
				number := capturedNumber(ni, nj, grid)
				if _, exists := seenNumbers[number]; !exists {
					seenNumbers[number] += 1
				}
			}
		}
	}

	return seenNumbers
}
