package main

import (
	"math"
	"strings"
)

//https://adventofcode.com/2021/day/5

type Point struct {
	x   int
	y   int
	key int
}

func newPoint(x int, y int) Point {
	return Point{x, y, ((x+y)*(x+y+1))/2 + y}
}

func day5() (int, int) {
	return day5Task1(), day5Task2()
}

func day5Task1() int {
	return calculatOverlappedFiels(false)
}

func day5Task2() int {
	return calculatOverlappedFiels(true)
}

func calculatOverlappedFiels(withDiagonals bool) int {
	data := fileToStringArray("input/5/input.txt")
	coordsMap := make(map[int]int)
	overlappedFieldsCounter := 0
	for _, row := range data {
		coords := strings.Split(row, " -> ")
		startCoords := stringToIntArray(strings.Split(coords[0], ","))
		endCoords := stringToIntArray(strings.Split(coords[1], ","))
		if startCoords[0] != endCoords[0] && startCoords[1] != endCoords[1] {
			if !withDiagonals {
				continue
			}
			if math.Abs(float64(startCoords[0]-endCoords[0])) != math.Abs(float64(startCoords[1]-endCoords[1])) {
				continue
			}
		}

		startPoint := newPoint(startCoords[0], startCoords[1])
		endPoint := newPoint(endCoords[0], endCoords[1])

		overlappedFieldsCounter += fillCoordMap(coordsMap, &startPoint, &endPoint)
	}

	return overlappedFieldsCounter
}

func fillCoordMap(coordsMap map[int]int, startPoint *Point, endPoint *Point) int {
	points := calculatePointsBetween(startPoint, endPoint)
	overlappedFieldsCounter := 0
	for _, point := range points {
		coordsMap[point.key]++
		if coordsMap[point.key] == 2 {
			overlappedFieldsCounter++
		}
	}
	return overlappedFieldsCounter
}

func calculatePointsBetween(firstPoint *Point, secondPoint *Point) []Point {
	startingPoint, endingPoint := firstPoint, secondPoint
	var points []Point
	if firstPoint.x > secondPoint.x {
		startingPoint = secondPoint
		endingPoint = firstPoint
	}

	//Horizontals and verticals
	if startingPoint.x == endingPoint.x || startingPoint.y == endingPoint.y {
		if firstPoint.y > secondPoint.y {
			startingPoint = secondPoint
			endingPoint = firstPoint
		}
		for i := startingPoint.x; i <= endingPoint.x; i++ {
			for j := startingPoint.y; j <= endingPoint.y; j++ {
				points = append(points, newPoint(i, j))
			}
		}
		return points
	}

	//Diagonals
	distance := float64(startingPoint.y - endingPoint.y)
	var point Point
	for i := 0; i <= int(math.Abs(distance)); i++ {
		if distance < 0 {
			point = newPoint(startingPoint.x+i, startingPoint.y+i)
		} else {
			point = newPoint(startingPoint.x+i, startingPoint.y-i)
		}
		points = append(points, point)
	}

	return points
}
