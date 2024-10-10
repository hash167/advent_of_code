package aoc2023

import (
	"algos/aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseTextToMapA(input string) map[string]interface{} {
	iterator, _ := utils.NewLineIterator(input)
	result := make(map[string]interface{})

	var currentKey string
	for {
		line, ok := iterator.Next()
		if !ok {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			currentKey = strings.TrimSuffix(line, ":")
			result[currentKey] = [][]int{}
		} else if strings.HasPrefix(line, "seeds:") {
			currentKey = "seeds"
			values := parseValues(line[len("seeds:"):])
			result[currentKey] = values
		} else if currentKey != "" {
			values := parseValues(line)
			if _, ok := result[currentKey].([][]int); ok {
				result[currentKey] = append(result[currentKey].([][]int), values)
			}
		}
	}

	return result
}

func parseValues(input string) []int {
	valuesStr := strings.Fields(input)
	values := make([]int, len(valuesStr))
	for i, valueStr := range valuesStr {
		value, err := strconv.Atoi(valueStr)
		if err == nil {
			values[i] = value
		}
	}
	return values
}

func mapValue(mapping [][]int, key int) int {

	for _, entry := range mapping {
		if len(entry) == 3 && key >= entry[1] && key < entry[1]+entry[2] {
			mappedValue := entry[0] + (key - entry[1])
			return mappedValue
		}
	}

	return key
}

func mapSeedsToLocationLazy(parsedMap map[string]interface{}) map[int]int {
	maps := []string{"seed-to-soil map", "soil-to-fertilizer map", "fertilizer-to-water map", "water-to-light map", "light-to-temperature map", "temperature-to-humidity map", "humidity-to-location map"}
	seedToLocation := make(map[int]int)

	seeds := parsedMap["seeds"].([]int)
	for _, seed := range seeds {
		mappedValue := seed
		for _, mapKey := range maps {
			mapping := parsedMap[mapKey].([][]int)
			mappedValue = mapValue(mapping, mappedValue)
		}
		seedToLocation[seed] = mappedValue
	}

	return seedToLocation
}

func SolveDay5(inputStr string) {
	parsedMap := parseTextToMapA(inputStr)
	seedToLocation := mapSeedsToLocationLazy(parsedMap)

	lowestSeed, lowestLocation := findLowestLocation(seedToLocation)
	fmt.Printf("The seed %d is mapped to the lowest location %d\n", lowestSeed, lowestLocation)
}

func findLowestLocation(seedToLocation map[int]int) (int, int) {
	lowestSeed := 0
	lowestLocation := int(^uint(0) >> 1) // Max int value
	for seed, location := range seedToLocation {
		if location < lowestLocation {
			lowestSeed = seed
			lowestLocation = location
		}
	}
	return lowestSeed, lowestLocation
}
