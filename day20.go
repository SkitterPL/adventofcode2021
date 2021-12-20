package main

import (
	"fmt"
	"strconv"
)

//https://adventofcode.com/2021/day/20

func day20() (int, int) {
	data := fileToStringArray("input/20/input.txt")
	pattern := data[0]
	matrix := make([][]int, len(data)+4)
	standardRowLen := len(data[2])
	matrix[0] = make([]int, standardRowLen+6)
	matrix[1] = make([]int, standardRowLen+6)
	matrix[2] = make([]int, standardRowLen+6)
	matrix[standardRowLen+3] = make([]int, len(data[2])+6)
	matrix[standardRowLen+4] = make([]int, len(data[2])+6)
	matrix[standardRowLen+5] = make([]int, len(data[2])+6)
	for key, line := range data[2:] {
		matrix[key+3] = make([]int, standardRowLen+6)
		matrix[key+3][0] = 0
		matrix[key+3][1] = 0
		matrix[key+3][2] = 0
		for i := 3; i < standardRowLen+3; i++ {
			if line[i-3] == '#' {
				matrix[key+3][i] = 1
			} else {
				matrix[key+3][i] = 0
			}
		}
		matrix[key+3][standardRowLen+3] = 0
		matrix[key+3][standardRowLen+4] = 0
		matrix[key+3][standardRowLen+5] = 0
	}

	return enhance(matrix, pattern, 2), enhance(matrix, pattern, 50)
}

func enhance(matrix [][]int, pattern string, times int) int {
	newMatrix := matrix
	var counter int
	for i := 0; i < times; i++ {
		newMatrix, counter = enhanceImage(newMatrix, pattern, i+1)
	}
	return counter
}

func enhanceImage(image [][]int, pattern string, iteration int) ([][]int, int) {
	height := len(image)
	width := len(image[0])
	newImage := make([][]int, height)
	counter := 0
	var infiniteBit int
	if pattern[0] == '.' {
		infiniteBit = 0
		if iteration%2 == 0 {
			infiniteBit = 0
		}
	} else {
		infiniteBit = 0
		if iteration%2 == 0 {
			infiniteBit = 1
		}
	}

	for key, row := range image {
		newImage[key] = make([]int, width)
		for key2, _ := range row {
			number := calculateBinaryNumber(key, key2, image, height, width, infiniteBit)
			if pattern[number] == '.' {
				newImage[key][key2] = 0
			} else {
				newImage[key][key2] = 1
				counter++
			}
		}
	}

	return newImage, counter
}

func calculateBinaryNumber(y int, x int, image [][]int, height int, width int, infiniteBit int) int {
	number := ""
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			number += fmt.Sprint(calculatePixel(y+i, x+j, image, height, width, infiniteBit))
		}
	}
	num, _ := strconv.ParseInt(number, 2, 64)
	return int(num)
}

func calculatePixel(y int, x int, image [][]int, height int, width int, infiniteBit int) rune {
	if y < 0 || x < 0 || y >= height || x >= width {
		return rune(infiniteBit)
	} else {
		return rune(image[y][x])
	}
}
