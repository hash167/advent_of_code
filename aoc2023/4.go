package aoc2023

import (
	"algos/aoc2023/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func SolveDay4(input string) {
	iterator, err := utils.NewLineIterator(input)
	tp := 0
	if err != nil {
		panic(err)
	}
	gameWins := []int{}
	gameNum := 1
	for {
		line, ok := iterator.Next()
		if !ok {
			break
		}
		left, right, _ := parseLine(line)
		count := countItemsInLeftFromRight(left, right)
		points := powerOfTwoMinusOne(count)
		gameWins = append(gameWins, count)
		tp += points
		gameNum++

	}
	fmt.Println("Result of part A is", tp)
	fmt.Println("Result of part B is", SolveDay4PartB(gameWins))

}

func recursiveCount(input []int, index int, numCopies *map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that Done is called at the end of the function

	mu.Lock()
	(*numCopies)[index] += 1
	mu.Unlock()

	if input[index] == 0 {
		return
	}

	n := input[index]
	for j := 0; j < n; j++ {
		if index+j+1 < len(input) {
			wg.Add(1)
			go recursiveCount(input, index+j+1, numCopies, mu, wg)
		}
	}
}

func SolveDay4PartB(input []int) int {
	numCopies := make(map[int]int) // Initialize the map
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < len(input); i++ {
		wg.Add(1)
		go recursiveCount(input, i, &numCopies, &mu, &wg)

	}

	wg.Wait() // Wait for all goroutines to complete

	var ret int
	for _, v := range numCopies {
		ret += v
	}
	return ret
}

func powerOfTwoMinusOne(num int) int {
	// Calculate 2 to the power of (num - 1)
	result := math.Pow(2, float64(num-1))
	// Convert the result to an integer
	return int(result)
}

// countItemsInLeftFromRight counts the number of items from the right that appear in the left.
func countItemsInLeftFromRight(left []int, right []int) int {
	leftSet := make(map[int]struct{})
	for _, num := range left {
		leftSet[num] = struct{}{}
	}

	count := 0
	for _, num := range right {
		if _, found := leftSet[num]; found {
			count++
		}
	}

	return count
}

func parseLine(input string) ([]int, []int, error) {
	parts := strings.Split(input, ":")
	numParts := strings.Split(parts[1], "|")
	left := strings.Split(numParts[0], " ")
	right := strings.Split(numParts[1], " ")
	leftNums := make([]int, 0)
	rightNums := make([]int, 0)
	for _, num := range left {
		n, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		leftNums = append(leftNums, n)
	}
	for _, num := range right {
		n, err := strconv.Atoi(num)
		if err != nil {
			// not number, most likely whitespace
			continue
		}
		rightNums = append(rightNums, n)
	}
	return leftNums, rightNums, nil
}
