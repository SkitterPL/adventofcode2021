package main

import (
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/5

func day5() (int, int) {
	return day5task1(), 0
}

func day5task1() int {
	data := fileToStringArray("input/5/test_input.txt")
	coordsMap := make(map[int]map[int]int)
	overlappingFields := make(map[string]int)
	for _, row := range data {
		coords := strings.Split(row, " -> ")
		startCoords := stringToIntArray(strings.Split(coords[0], ","))
		endCoords := stringToIntArray(strings.Split(coords[1], ","))
		if startCoords[0] != endCoords[0] && startCoords[1] != endCoords[1] {
			continue
		}

		fillCoordMap(coordsMap, overlappingFields, startCoords, endCoords)
	}

	return len(overlappingFields)
}

func fillCoordMap(coordsMap map[int]map[int]int, overlappingFields map[string]int, startCoords []int, endCoords []int) {
	startFillCoords, endFillCoords := specifyFillDirection(startCoords, endCoords)
	for i := startFillCoords[0]; i <= endFillCoords[0]; i++ {
		for j := startFillCoords[1]; j <= endFillCoords[1]; j++ {
			if coordsMap[i] == nil {
				coordsMap[i] = make(map[int]int)
				coordsMap[i][j] = 1
			} else {
				coordsMap[i][j]++
				if coordsMap[i][j] >= 2 {
					overlappingFields[strconv.Itoa(i)+","+strconv.Itoa(j)] = 1
				}

			}
		}
	}

}

func specifyFillDirection(startCoords []int, endCoords []int) ([]int, []int) {
	startFillCoords, endFillCoords := startCoords, endCoords
	if startCoords[0] > endCoords[0] {
		startFillCoords = endCoords
		endFillCoords = startCoords
	}
	if startCoords[1] > endCoords[1] {
		startFillCoords = endCoords
		endFillCoords = startCoords
	}

	return startFillCoords, endFillCoords
}
