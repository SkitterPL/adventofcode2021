package main

import (
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/2

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

const (
	Forward string = "forward"
	Down    string = "down"
	Up      string = "up"
)

func day2() (int, int) {
	data := fileToStringArray("input/2/input.txt")
	position, positionWithAim := calculatePosition(data)
	return position.result(), positionWithAim.result()
}

func calculatePosition(data []string) (*Position, *Position) {
	var position, positionWithAim Position = Position{0, 0, 0}, Position{0, 0, 0}

	for _, line := range data {
		splittedDirectionAndValue := strings.Fields(line)
		direction := splittedDirectionAndValue[0]
		convertedValue, _ := strconv.Atoi(splittedDirectionAndValue[1])
		position.move(direction, convertedValue)
		positionWithAim.moveUsingAim(direction, convertedValue)
	}

	return &position, &positionWithAim
}

func (position *Position) move(direction string, value int) {
	switch direction {
	case Forward:
		position.Horizontal += value
	case Down:
		position.Depth += value
	case Up:
		position.Depth -= value
	}
}

func (position *Position) moveUsingAim(direction string, value int) {
	switch direction {
	case Forward:
		position.Horizontal += value
		position.Depth += position.Aim * value
	case Down:
		position.Aim += value
	case Up:
		position.Aim -= value
	}
}

func (position *Position) result() int {
	return position.Depth * position.Horizontal
}
