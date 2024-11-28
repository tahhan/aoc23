package day05_2

import (
	"aoc23/utils"
	"math"
	"strconv"
	"strings"
)

type MapLookup struct {
	from string
	to   string
}
type LookupRange struct {
	sourceStart      int
	destinationStart int
	rangeLength      int
}

type SeedRange struct {
	rangeStart  int
	rangeLength int
}

func (lr LookupRange) getDestination(source int) int {
	if source >= lr.sourceStart && source <= lr.sourceStart+lr.rangeLength {
		return lr.destinationStart + (source - lr.sourceStart)
	}
	return -1
}

func Seeds2() int {
	var seedsRanges []SeedRange
	mapLookups := map[MapLookup][]LookupRange{}
	var currentMapLookup MapLookup
	utils.MapFileLines("day05-2/input05-2.txt", func(line string) {
		if seedsNumbersPart, found := strings.CutPrefix(line, "seeds: "); found {
			seedsNumbers := strings.Split(seedsNumbersPart, " ")

			seedRange := SeedRange{rangeStart: -1, rangeLength: -1}
			for _, seedNumber := range seedsNumbers {
				if seedRange.rangeStart == -1 {
					number, _ := strconv.Atoi(seedNumber)
					seedRange.rangeStart = number
				} else {
					number, _ := strconv.Atoi(seedNumber)
					seedRange.rangeLength = number
					seedsRanges = append(seedsRanges, seedRange)
					seedRange = SeedRange{rangeStart: -1, rangeLength: -1}
				}
			}

			if seedRange.rangeStart != -1 && seedRange.rangeLength != -1 {
				seedsRanges = append(seedsRanges, seedRange)
			}

		} else if mapTypePart, found := strings.CutSuffix(line, " map:"); found {
			mapTypeParts := strings.Split(mapTypePart, "-to-")
			mapLookup := MapLookup{from: mapTypeParts[0], to: mapTypeParts[1]}
			mapLookups[mapLookup] = []LookupRange{}
			currentMapLookup = mapLookup
		} else if line != "" {
			lookupLineParts := strings.Split(line, " ")
			destinationStart, _ := strconv.Atoi(lookupLineParts[0])
			sourceStart, _ := strconv.Atoi(lookupLineParts[1])
			rangeLength, _ := strconv.Atoi(lookupLineParts[2])
			mapLookups[currentMapLookup] = append(mapLookups[currentMapLookup], LookupRange{sourceStart: sourceStart, destinationStart: destinationStart, rangeLength: rangeLength})
		}
	})

	//fmt.Println(seedsRanges)

	seedLocation := math.MaxInt
	for _, seedsRange := range seedsRanges {
		for i := seedsRange.rangeStart; i < seedsRange.rangeStart+seedsRange.rangeLength; i++ {
			soilNumber := findInRanges(mapLookups, "seed", "soil", i)
			fertilizerNumber := findInRanges(mapLookups, "soil", "fertilizer", soilNumber)
			waterNumber := findInRanges(mapLookups, "fertilizer", "water", fertilizerNumber)
			lightNumber := findInRanges(mapLookups, "water", "light", waterNumber)
			temperatureNumber := findInRanges(mapLookups, "light", "temperature", lightNumber)
			humidityNumber := findInRanges(mapLookups, "temperature", "humidity", temperatureNumber)
			locationNumber := findInRanges(mapLookups, "humidity", "location", humidityNumber)

			//fmt.Println(i, soilNumber, fertilizerNumber, waterNumber, lightNumber, temperatureNumber, humidityNumber, locationNumber)
			seedLocation = int(math.Min(float64(seedLocation), float64(locationNumber)))
		}
	}

	return seedLocation
}

func findInRanges(mapLookups map[MapLookup][]LookupRange, from string, to string, needle int) int {
	sourceToDestinationRanges := mapLookups[MapLookup{from: from, to: to}]
	destinationNumber := -1
	for _, sourceToDestinationRange := range sourceToDestinationRanges {
		if destinationNumber = sourceToDestinationRange.getDestination(needle); destinationNumber != -1 {
			break
		}
	}

	if destinationNumber == -1 {
		destinationNumber = needle
	}

	return destinationNumber
}
