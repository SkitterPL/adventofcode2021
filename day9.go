package main

import (
	"sort"
	"strconv"
)

//https://adventofcode.com/2021/day/9

type HeightMap struct {
	width            int
	height           int
	numbers          [][]int
	visitedForBasins [][]bool
	lowestPoints     []*LowestPoint
}

type LowestPoint struct {
	x      int
	y      int
	height int
}

func (board *HeightMap) isLowestPoint(x int, y int) bool {
	current := board.numbers[y][x]
	if current < board.returnValueToCompare(x-1, y) &&
		current < board.returnValueToCompare(x+1, y) &&
		current < board.returnValueToCompare(x, y-1) &&
		current < board.returnValueToCompare(x, y+1) {
		return true
	}
	return false
}

func (board *HeightMap) returnValueToCompare(x int, y int) int {
	if y < 0 || y > board.height-1 {
		return 10
	}
	if x < 0 || x > board.width-1 {
		return 10
	}
	return board.numbers[y][x]
}

func (board *HeightMap) addLowestPoint(x int, y int, value int) {
	board.lowestPoints = append(board.lowestPoints, &LowestPoint{x, y, value})
}

func (board *HeightMap) calculateLowestPoints() {
	for rowKey := range board.numbers {
		for colKey, value := range board.numbers[rowKey] {
			if board.isLowestPoint(colKey, rowKey) {
				board.addLowestPoint(colKey, rowKey, value)
			}
		}
	}
}

func (board *HeightMap) calculateRisk() int {
	risk := 0
	for _, point := range board.lowestPoints {
		risk += point.height + 1
	}
	return risk
}

func (board *HeightMap) calculateLargestBasinsSizeMultiplier() int {
	basinSizes := []int{}
	for _, lowestPoint := range board.lowestPoints {
		size := board.calculateBasinSize(lowestPoint)
		basinSizes = append(basinSizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func (board *HeightMap) calculateBasinSize(point *LowestPoint) int {
	return board.walkUntilBasinBorder(0, point.y, point.x)
}

func (board *HeightMap) walkUntilBasinBorder(basinSize int, y int, x int) int {
	if y < 0 || y > board.height-1 {
		return basinSize
	}
	if x < 0 || x > board.width-1 {
		return basinSize
	}
	if board.visitedForBasins[y][x] {
		return basinSize
	}
	if board.numbers[y][x] == 9 {
		return basinSize
	}
	board.visitedForBasins[y][x] = true
	basinSize++
	return basinSize + board.walkUntilBasinBorder(0, y-1, x) +
		board.walkUntilBasinBorder(0, y+1, x) +
		board.walkUntilBasinBorder(0, y, x+1) +
		board.walkUntilBasinBorder(0, y, x-1)
}

func day9() (int, int) {
	board := prepareData(fileToStringArray("input/9/input.txt"))
	board.calculateLowestPoints()
	return board.calculateRisk(), board.calculateLargestBasinsSizeMultiplier()
}

func prepareData(data []string) *HeightMap {
	linesNumber := len(data)
	numbersInOneline := len(data[0])
	numbers := make([][]int, linesNumber)
	visited := make([][]bool, linesNumber)
	for i := 0; i < linesNumber; i++ {
		numbers[i] = make([]int, numbersInOneline)
		visited[i] = make([]bool, numbersInOneline)
		for key, character := range data[i] {
			numbers[i][key], _ = strconv.Atoi(string(character))
			visited[i][key] = false
		}
	}
	var lowestPointsArr []*LowestPoint
	return &HeightMap{numbersInOneline, linesNumber, numbers, visited, lowestPointsArr}
}
