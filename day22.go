package main

import (
	"math"
	"strconv"
	"strings"
)

//https://adventofcode.com/2021/day/22

type Cuboid struct {
	xMin     int
	xMax     int
	yMin     int
	yMax     int
	zMin     int
	zMax     int
	subtract bool
}

func (cuboid *Cuboid) calculateTurnedOn() float64 {
	result := (1 + math.Abs(float64(cuboid.xMax-cuboid.xMin))) * (1 + math.Abs(float64(cuboid.yMax-cuboid.yMin))) * (1 + math.Abs(float64(cuboid.zMax-cuboid.zMin)))
	if cuboid.subtract {
		return -1 * result
	} else {
		return result
	}
}

func day22() (int, int) {
	return calculateCuboids(true), calculateCuboids(false)
}

func calculateCuboids(hasInitializationProcedure bool) int {
	turnedOn := []*Cuboid{}
	data := fileToStringArray("input/22/input.txt")
	for _, line := range data {
		parsedData := strings.Fields(line)
		cuboidToModify := transformToCuboid(parsedData[1])
		if hasInitializationProcedure && !isInInitializationProcedure(cuboidToModify) {
			continue
		}
		if parsedData[0] == "on" {
			turnedOn = addCuboid(cuboidToModify, turnedOn, false)
		} else {
			turnedOn = addCuboid(cuboidToModify, turnedOn, true)
		}
	}
	sum := 0
	for _, cuboid := range turnedOn {
		sum += int(cuboid.calculateTurnedOn())
	}
	return sum
}

func addCuboid(cuboidToAdd *Cuboid, addedCuboids []*Cuboid, subtract bool) []*Cuboid {
	cuboidToAdd.subtract = subtract

	for _, turnedOnCuboid := range addedCuboids {
		potentialNewCuboid := turnedOnCuboid.findCommonCuboid(cuboidToAdd)
		if potentialNewCuboid != nil {
			addedCuboids = append(addedCuboids, potentialNewCuboid)
		}
	}
	if !subtract {
		addedCuboids = append(addedCuboids, cuboidToAdd)
	}

	return addedCuboids
}

func isInInitializationProcedure(cuboid *Cuboid) bool {
	initializationProcedure := &Cuboid{-50, 50, -50, 50, -50, 50, false}
	return cuboid.isWithin(initializationProcedure)
}

func (cuboid *Cuboid) findCommonCuboid(anotherCuboid *Cuboid) *Cuboid {

	if cuboid.xMin < anotherCuboid.xMin && cuboid.xMax < anotherCuboid.xMin {
		return nil
	}
	if anotherCuboid.xMin < cuboid.xMin && anotherCuboid.xMax < cuboid.xMin {
		return nil
	}

	minX := math.Max(float64(cuboid.xMin), float64(anotherCuboid.xMin))
	maxX := math.Min(float64(cuboid.xMax), float64(anotherCuboid.xMax))

	minY := math.Max(float64(cuboid.yMin), float64(anotherCuboid.yMin))
	maxY := math.Min(float64(cuboid.yMax), float64(anotherCuboid.yMax))

	minZ := math.Max(float64(cuboid.zMin), float64(anotherCuboid.zMin))
	maxZ := math.Min(float64(cuboid.zMax), float64(anotherCuboid.zMax))

	if minX <= maxX && minY <= maxY && minZ <= maxZ {
		return &Cuboid{int(minX), int(maxX), int(minY), int(maxY), int(minZ), int(maxZ), !cuboid.subtract}
	}
	return nil
}

func (cuboid *Cuboid) isWithin(anotherCuboid *Cuboid) bool {
	return cuboid.xMin >= anotherCuboid.xMin && cuboid.xMax <= anotherCuboid.xMax &&
		cuboid.yMin >= anotherCuboid.yMin && cuboid.yMax <= anotherCuboid.yMax &&
		cuboid.zMin >= anotherCuboid.zMin && cuboid.zMax <= anotherCuboid.zMax
}

func transformToCuboid(line string) *Cuboid {
	data := strings.Split(line, ",")
	xRange := strings.Split(strings.Split(data[0], "=")[1], "..")
	yRange := strings.Split(strings.Split(data[1], "=")[1], "..")
	zRange := strings.Split(strings.Split(data[2], "=")[1], "..")
	xMin, _ := strconv.Atoi(xRange[0])
	xMax, _ := strconv.Atoi(xRange[1])
	yMin, _ := strconv.Atoi(yRange[0])
	yMax, _ := strconv.Atoi(yRange[1])
	zMin, _ := strconv.Atoi(zRange[0])
	zMax, _ := strconv.Atoi(zRange[1])
	return &Cuboid{xMin, xMax, yMin, yMax, zMin, zMax, false}
}
