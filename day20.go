package main

import (
	"fmt"
	"strconv"
)

const SIZE = 100

//https://adventofcode.com/2021/day/20

func day20() (int, int) {
	data := fileToStringArray("input/20/input.txt")
	pattern := data[0]
	matrix := make([][]int, len(data)+SIZE-2)
	standardRowLen := len(data[2])
	for i := 0; i < SIZE/2; i++ {
		matrix[i] = make([]int, standardRowLen+SIZE)
	}
	for i := standardRowLen + SIZE/2; i < standardRowLen+SIZE; i++ {
		matrix[i] = make([]int, standardRowLen+SIZE)
	}
	for key, line := range data[2:] {
		matrix[key+SIZE/2] = make([]int, standardRowLen+SIZE)
		for i := 0; i < SIZE/2; i++ {
			matrix[key+SIZE/2][i] = 0
		}
		for i := standardRowLen + SIZE/2; i < standardRowLen+SIZE; i++ {
			matrix[key+SIZE/2][i] = 0
		}
		for i := SIZE / 2; i < standardRowLen+SIZE/2; i++ {
			if line[i-SIZE/2] == '#' {
				matrix[key+SIZE/2][i] = 1
			} else {
				matrix[key+SIZE/2][i] = 0
			}
		}
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
