package main

import (
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/17

const StartXVelocityGuessing = 0
const StartYVelocityGuessing = 150
const InitialMaxYAchieved = -1000

type TargetArea struct {
	minX int
	maxX int
	minY int
	maxY int
}

type InitialPoint struct {
	x            int
	y            int
	xVelocity    int
	yVelocity    int
	maxAchievedY int
}

func day17() (int, int) {
	targetArea := newTargetArea(fileToStringArray("input/17/input.txt")[0])
	var point *InitialPoint
	maxYAchieved := InitialMaxYAchieved
	probesWithinArea := 0
	for i := StartXVelocityGuessing; i <= targetArea.maxX; i++ {
		for j := StartYVelocityGuessing; j >= targetArea.maxY; j-- {
			point = &InitialPoint{0, 0, i, j, InitialMaxYAchieved}
			if point.launchProbe(targetArea) {
				probesWithinArea++
				if point.maxAchievedY > maxYAchieved {
					maxYAchieved = point.maxAchievedY
				}
			}
		}
	}
	return maxYAchieved, probesWithinArea
}

func newTargetArea(input string) *TargetArea {
	data := strings.Split(input, "=")
	xRange := strings.Split(data[1], "..")
	minX, _ := strconv.Atoi(xRange[0])
	maxX, _ := strconv.Atoi(strings.Split(xRange[1], ",")[0])
	yRange := strings.Split(data[2], "..")
	minY, _ := strconv.Atoi(yRange[0])
	maxY, _ := strconv.Atoi(yRange[1])
	return &TargetArea{minX, maxX, maxY, minY}
}

func (point *InitialPoint) launchProbe(area *TargetArea) bool {
	landedInArea := false
	for !point.isBeyond(area) {
		point.step()
		if point.isWithin(area) {
			landedInArea = true
		}
	}
	return landedInArea
}

func (point *InitialPoint) step() {
	point.x += point.xVelocity
	point.y += point.yVelocity
	if point.y > point.maxAchievedY {
		point.maxAchievedY = point.y
	}
	if point.xVelocity > 0 {
		point.xVelocity--
	} else if point.xVelocity < 0 {
		point.xVelocity++
	}
	point.yVelocity--
}

func (point *InitialPoint) isWithin(area *TargetArea) bool {
	return point.x >= area.minX &&
		point.x <= area.maxX &&
		point.y <= area.minY &&
		point.y >= area.maxY
}

func (point *InitialPoint) isBeyond(area *TargetArea) bool {
	return point.x > area.maxX ||
		point.y < area.maxY
}
