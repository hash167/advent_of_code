package aoc2023

import (
	"strconv"
	"strings"
)

func parseTextToMap(input string) map[string]interface{} {
	lines := strings.Split(input, "\n")
	result := make(map[string]interface{})

	var currentKey string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		switch {
		case strings.HasSuffix(line, "map:"):
			currentKey = strings.TrimSuffix(line, ":")
			result[currentKey] = [][]int{}
		case strings.HasPrefix(line, "seeds:"):
			currentKey = "seeds"
			result[currentKey] = parseSeedsLine(line)
		case currentKey != "":
			if _, ok := result[currentKey].([][]int); ok {
				result[currentKey] = append(result[currentKey].([][]int), parseMappingLine(line))
			}
		}
	}

	return result
}

func parseSeedsLine(line string) []SeedRange {
	valuesStr := strings.Fields(strings.TrimPrefix(line, "seeds:"))
	seedRanges := []SeedRange{}
	for i := 0; i < len(valuesStr); i += 2 {
		startValue, err1 := strconv.Atoi(valuesStr[i])
		rangeLength, err2 := strconv.Atoi(valuesStr[i+1])
		if err1 == nil && err2 == nil {
			seedRanges = append(seedRanges, SeedRange{start: startValue, length: rangeLength})
		}
	}
	return seedRanges
}

func parseMappingLine(line string) []int {
	valuesStr := strings.Fields(line)
	values := []int{}
	for _, valueStr := range valuesStr {
		value, err := strconv.Atoi(valueStr)
		if err == nil {
			values = append(values, value)
		}
	}
	return values
}
