package aoc2023

import (
	"algos/aoc2023/utils"
	"regexp"
	"strconv"
)

var checkFor = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func getMaxColors(s string) map[string]map[string]int {
	// Regular expression to match game IDs and counts with colors
	reGame := regexp.MustCompile(`Game\s+(\d+):\s*([^;]+(?:;\s*[^;]+)*)`)
	reColor := regexp.MustCompile(`(\d+)\s+(blue|red|green)`)

	// Map to store the results for each game
	gamesMax := make(map[string]map[string]int)

	// Find all games
	gameMatches := reGame.FindAllStringSubmatch(s, -1)

	// Iterate over game matches
	for _, gameMatch := range gameMatches {
		gameID := gameMatch[1]
		gameData := gameMatch[2]

		// Initialize map for each game ID
		if _, exists := gamesMax[gameID]; !exists {
			gamesMax[gameID] = make(map[string]int)
		}

		// Find all color matches within the game data
		colorMatches := reColor.FindAllStringSubmatch(gameData, -1)

		// Iterate over color matches and update the maximum count for each color
		for _, colorMatch := range colorMatches {
			count, _ := strconv.Atoi(colorMatch[1])
			color := colorMatch[2]
			if count > gamesMax[gameID][color] {
				gamesMax[gameID][color] = count
			}
		}
	}

	return gamesMax
}

func SolveDay2(input string) {
	it, err := utils.NewLineIterator(input)
	if err != nil {
		panic(err)
	}

	var sum2a, sum2b int
	for {
		line, ok := it.Next()
		if !ok {
			break
		}
		gamesMax := getMaxColors(line)
		for gameID, colors := range gamesMax {
			addGame := true
			product := 1
			for color, count := range colors {
				product *= count
				if count > checkFor[color] {
					addGame = false

				}
			}
			if addGame {
				gameIDInt, _ := strconv.Atoi(gameID)
				sum2a += gameIDInt
			}
			sum2b += product
		}
	}
	println(sum2a)
	println(sum2b)
}
