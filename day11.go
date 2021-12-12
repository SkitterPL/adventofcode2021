package main

//https://adventofcode.com/2021/day/11

type OctopusCavern struct {
	octopuses [][]int
	flashed   [10][10]bool
	flashes   int
	length    int
}

func (cavern *OctopusCavern) makeStep() bool {
	for rowKey, row := range cavern.octopuses {
		for colKey := range row {
			cavern.octopuses[rowKey][colKey]++
			if !cavern.flashed[rowKey][colKey] && cavern.octopuses[rowKey][colKey] >= 10 {
				cavern.flash(rowKey, colKey)
			}
		}
	}
	allFlashed := true
	for rowKey, row := range cavern.octopuses {
		for colKey := range row {
			if cavern.flashed[rowKey][colKey] {
				cavern.octopuses[rowKey][colKey] = 0
				cavern.flashed[rowKey][colKey] = false
			} else {
				allFlashed = false
			}
		}
	}
	return allFlashed
}

func (cavern *OctopusCavern) flash(y int, x int) {
	cavern.flashes++
	cavern.flashed[y][x] = true
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if y+i < 0 || y+i >= cavern.length || x+j < 0 || x+j >= cavern.length {
				continue
			}
			cavern.octopuses[y+i][x+j]++
			if !cavern.flashed[y+i][x+j] && cavern.octopuses[y+i][x+j] >= 10 {
				cavern.flash(y+i, x+j)
			}
		}
	}
}
func day11() (int, int) {
	data := fileTo2DIntArray("input/11/input.txt")
	return getThroughCavern(data)
}

func getThroughCavern(data [][]int) (int, int) {

	flashed := [10][10]bool{}
	cavern := OctopusCavern{data, flashed, 0, 10}
	flashesNumber := 0
	firstAllFlashStep := 0
	stepCounter := 1
	for {
		if cavern.makeStep() {
			firstAllFlashStep = stepCounter
			break
		}
		if stepCounter == 100 {
			flashesNumber = cavern.flashes
		}

		stepCounter++
	}
	return flashesNumber, firstAllFlashStep
}
