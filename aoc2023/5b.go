package aoc2023

import (
	"fmt"
	"os"
)

type Range struct {
	lower int
	upper int
}

type RangeMapping struct {
	sourceStart int
	destStart   int
	length      int
}

type SeedRange struct {
	start  int
	length int
}

func generateListFromMappings(parsedMap map[string]interface{}, key string) []RangeMapping {
	mappingSlice, ok := parsedMap[key].([][]int)
	if !ok {
		return nil
	}

	ranges := []RangeMapping{}
	for _, entry := range mappingSlice {
		if len(entry) == 3 {
			ranges = append(ranges, RangeMapping{
				sourceStart: entry[1],
				destStart:   entry[0],
				length:      entry[2],
			})
		}
	}
	return ranges
}

func divide(r Range, mappingRange Range) (Range, Range, Range) {
	// Divide the range into lower, intersecting, and upper parts
	lower := Range{lower: r.lower, upper: min(r.upper, mappingRange.lower-1)}
	intersect := Range{lower: max(r.lower, mappingRange.lower), upper: min(r.upper, mappingRange.upper)}
	upper := Range{lower: max(r.lower, mappingRange.upper+1), upper: r.upper}

	if lower.lower > lower.upper {
		lower = Range{} // Empty
	}
	if intersect.lower > intersect.upper {
		intersect = Range{} // Empty
	}
	if upper.lower > upper.upper {
		upper = Range{} // Empty
	}

	return lower, intersect, upper
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mapIntersect(mapping RangeMapping, intersect Range) Range {
	diffLower := intersect.lower - mapping.sourceStart
	intersectLen := intersect.upper - intersect.lower
	return Range{lower: mapping.destStart + diffLower, upper: mapping.destStart + diffLower + intersectLen}
}

func processMapSeeds(mappings []RangeMapping, seeds []Range, edgecases bool) []Range {
	ranges := []Range{}
	for len(seeds) > 0 {
		item := seeds[0]
		seeds = seeds[1:]
		lenBefore := len(ranges)
		for _, mapping := range mappings {
			mappingRange := Range{lower: mapping.sourceStart, upper: mapping.sourceStart + mapping.length - 1}
			lower, intersect, upper := divide(item, mappingRange)
			if intersect == (Range{}) {
				// No intersection found
				continue
			}
			ranges = append(ranges, mapIntersect(mapping, intersect))
			if intersect.lower == item.lower && intersect.upper == item.upper {
				// Full intersect
				break
			}
			if edgecases {
				// Continue processing the non-intersecting parts
				if upper != (Range{}) && upper.upper == item.upper {
					seeds = append(seeds, upper)
				} else if lower != (Range{}) && lower.lower == item.lower {
					seeds = append(seeds, lower)
				}
			}
		}
		if lenBefore == len(ranges) {
			ranges = append(ranges, item) // No mapping found, keep the original
		}
	}
	return ranges
}

func mapSeedsToLocationOptimized(parsedMap map[string]interface{}) (int, int) {
	// Generate mappings from parsed map
	seedToSoilMap := generateListFromMappings(parsedMap, "seed-to-soil map")
	soilToFertilizerMap := generateListFromMappings(parsedMap, "soil-to-fertilizer map")
	fertilizerToWaterMap := generateListFromMappings(parsedMap, "fertilizer-to-water map")
	waterToLightMap := generateListFromMappings(parsedMap, "water-to-light map")
	lightToTemperatureMap := generateListFromMappings(parsedMap, "light-to-temperature map")
	temperatureToHumidityMap := generateListFromMappings(parsedMap, "temperature-to-humidity map")
	humidityToLocationMap := generateListFromMappings(parsedMap, "humidity-to-location map")

	mappings := [][]RangeMapping{
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	}

	lowestSeed := -1
	lowestLocation := int(^uint(0) >> 1) // Max int value

	seedRanges, ok := parsedMap["seeds"].([]SeedRange)
	if !ok {
		return lowestSeed, lowestLocation
	}

	// Convert seed ranges to Range type
	seeds := []Range{}
	for _, seedRange := range seedRanges {
		seeds = append(seeds, Range{lower: seedRange.start, upper: seedRange.start + seedRange.length - 1})
	}

	for _, mappingList := range mappings {
		seeds = processMapSeeds(mappingList, seeds, true)
	}

	for _, r := range seeds {
		if r.lower < lowestLocation {
			lowestLocation = r.lower
			lowestSeed = r.lower
			fmt.Printf("[DEBUG] New lowest seed %d mapped to location %d\n", lowestSeed, lowestLocation)
		}
	}

	return lowestSeed, lowestLocation
}

func SolveDay5b(inputStr string) {
	inputBytes, err := os.ReadFile(inputStr)
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)
	parsedMap := parseTextToMap(input)

	fmt.Println("[DEBUG] Parsed map:", parsedMap)

	// Find the seed with the lowest location
	lowestSeed, lowestLocation := mapSeedsToLocationOptimized(parsedMap)
	fmt.Printf("\nSeed with the lowest location: %d (Location: %d)\n", lowestSeed, lowestLocation)
}
