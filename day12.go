package main

import (
	"strings"
)

//https://adventofcode.com/2021/day/12

type Cave struct {
	name        string
	connections []*Cave
	small       bool
}

func newCave(name string) *Cave {
	return &Cave{name, []*Cave{}, name == strings.ToLower(name)}
}

func (cave *Cave) addConnection(connectedCave *Cave) {
	cave.connections = append(cave.connections, connectedCave)
}

func (cave *Cave) move(alreadyVisited []string, connectedCave *Cave, possiblePaths int, smallCaveUsed bool) int {
	alreadyVisited = append(alreadyVisited, connectedCave.name)
	for _, connection := range connectedCave.connections {
		if connection.name == "end" {
			possiblePaths++
			continue
		}
		if connection.small && hasAlreadyVisitedTwice(alreadyVisited, connection, smallCaveUsed) {
			continue
		}
		caveUsed := smallCaveUsed || (connection.small && hasAlreadyVisited(alreadyVisited, connection))
		possiblePaths = connectedCave.move(alreadyVisited, connection, possiblePaths, caveUsed)
	}
	return possiblePaths
}

func countPossiblePaths(startCave *Cave, visitSingleSmallCaveTwice bool) int {
	paths := 0
	for _, connection := range startCave.connections {
		visited := []string{}
		if !visitSingleSmallCaveTwice {
			paths += startCave.move(visited, connection, 0, true)
		} else {
			paths += startCave.move(visited, connection, 0, false)
		}
	}
	return paths
}

func hasAlreadyVisited(visited []string, cave *Cave) bool {
	if cave.name == "end" || cave.name == "start" {
		return true
	}
	for _, caveName := range visited {
		if caveName == cave.name {
			return true
		}
	}
	return false
}

func hasAlreadyVisitedTwice(visited []string, cave *Cave, singleCaveUsed bool) bool {
	if cave.name == "end" || cave.name == "start" {
		return true
	}
	sum := 0
	for _, caveName := range visited {
		if caveName == cave.name {
			sum++
		}
	}
	return sum == 2 || (singleCaveUsed && sum == 1)
}

func day12() (int, int) {
	data := fileToStringArray("input/12/input.txt")
	caves := make(map[string]*Cave)
	for _, line := range data {
		connection := strings.Split(line, "-")
		if caves[connection[0]] == nil {
			caves[connection[0]] = newCave(connection[0])
		}
		if caves[connection[1]] == nil {
			caves[connection[1]] = newCave(connection[1])
		}
		caves[connection[0]].addConnection(caves[connection[1]])
		caves[connection[1]].addConnection(caves[connection[0]])
	}

	return countPossiblePaths(caves["start"], false), countPossiblePaths(caves["start"], true)
}
