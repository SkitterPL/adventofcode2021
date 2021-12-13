package main

import (
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/13

type PaperPoint struct {
	x int
	y int
}

func day13() (string, int) {
	data := fileToStringArray("input/13/input.txt")
	points := make(map[string]*PaperPoint, len(data))
	instructionsBegin := false
	firstFoldResult := 0
	for _, line := range data {
		if line == "" {
			instructionsBegin = true
			continue
		}
		if !instructionsBegin {
			parsedPoint := strings.Split(line, ",")
			x, _ := strconv.Atoi(parsedPoint[0])
			y, _ := strconv.Atoi(parsedPoint[1])
			points[line] = &PaperPoint{x, y}
		} else {
			instruction := strings.Fields(line)[2]
			parsedInsruction := strings.Split(instruction, "=")
			middle, _ := strconv.Atoi(parsedInsruction[1])
			if parsedInsruction[0] == "x" {
				foldLeft(middle, points)
			} else {
				foldUp(middle, points)
			}
			if firstFoldResult == 0 {
				firstFoldResult = len(points)
			}
		}
	}
	//Part 2 drawing
	maxY, maxX := 0, 0
	for _, point := range points {
		if point.y > maxY {
			maxY = point.y
		}
		if point.x > maxX {
			maxX = point.x
		}
	}

	table := make([][]rune, maxY+1)
	for i := 0; i < maxY+1; i++ {
		row := make([]rune, maxX+1)
		for j := 0; j < maxX+1; j++ {
			row[j] = '.'
		}
		table[i] = row
	}

	for _, point := range points {
		table[point.y][point.x] = '#'
	}

	result := ""
	for _, row := range table {
		for _, col := range row {
			result += string(col)
		}
		result += "\n"
	}

	return result, firstFoldResult
}

func foldUp(middle int, points map[string]*PaperPoint) {
	for _, point := range points {
		if point.y < middle {
			continue
		}
		newY := middle - (point.y - middle)
		newKey := strconv.Itoa(point.x) + "," + strconv.Itoa(newY)
		if points[newKey] == nil {
			points[newKey] = &PaperPoint{point.x, newY}
		}
		delete(points, strconv.Itoa(point.x)+","+strconv.Itoa(point.y))
	}
}

func foldLeft(middle int, points map[string]*PaperPoint) {
	for _, point := range points {
		if point.x < middle {
			continue
		}
		newX := middle - (point.x - middle)
		newKey := strconv.Itoa(newX) + "," + strconv.Itoa(point.y)
		if points[newKey] == nil {
			points[newKey] = &PaperPoint{newX, point.y}
		}
		delete(points, strconv.Itoa(point.x)+","+strconv.Itoa(point.y))
	}
}
