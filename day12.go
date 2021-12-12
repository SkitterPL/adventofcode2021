package main

import (
	"fmt"
	"strings"
)

//https://adventofcode.com/2021/day/12

type CaveMap struct {
	caves map[string]*Cave
}

type Cave struct {
	name        string
	connections map[string]*Cave
	small       bool
}

func newCave(name string) *Cave {
	small := false
	if name == strings.ToLower(name) {
		small = true
	}
	return &Cave{name, map[string]*Cave{}, small}
}

func (cave *Cave) addConnection(connectedCave *Cave) {
	cave.connections[connectedCave.name] = connectedCave
}

func (cave *Cave) move(alreadyVisited []string, connectedCave *Cave, possiblePaths int) int {
	alreadyVisited = append(alreadyVisited, connectedCave.name)
	for _, connection := range connectedCave.connections {
		if connection.name == "end" {
			possiblePaths++
			alreadyVisited = append(alreadyVisited, "end")
			fmt.Println(alreadyVisited)
			continue
		}
		if connection.small && hasAlreadyVisited(alreadyVisited, connection) {
			continue
		}
		possiblePaths = connectedCave.move(alreadyVisited, connection, possiblePaths)
	}
	return possiblePaths
}

func (caveMap *CaveMap) start() int {
	paths := 0
	startCave := caveMap.caves["start"]
	for _, connection := range startCave.connections {
		fmt.Println("start")
		visited := []string{}
		visited = append(visited, "start")
		paths += startCave.move(visited, connection, 0)
		fmt.Println("next")
	}
	return paths
}

func hasAlreadyVisited(visited []string, cave *Cave) bool {
	for _, caveName := range visited {
		if caveName == cave.name {
			return true
		}
	}
	return false
}

func day12() (int, int) {
	data := fileToStringArray("input/12/test_input.txt")
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
	caveMap := CaveMap{caves}

	return day12Task1(&caveMap), 0
}

func day12Task1(caveMap *CaveMap) int {
	return caveMap.start()
}
