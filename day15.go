package main

//https://adventofcode.com/2021/day/15

const MaxInt = int(^uint(0) >> 1)

type Cavern struct {
	distances [][]int
	unvisited [][]int
	length    int
}

func day15() (int, int) {
	data := fileTo2DIntArray("input/15/input.txt")
	biggerData := prepareBiggerData(data)
	return djikstra(data), djikstra(biggerData)
}

func prepareBiggerData(data [][]int) [][]int {
	length := len(data)
	newData := make([][]int, 5*length)
	for i := 0; i < 5*length; i++ {
		newData[i] = make([]int, 5*length)
		for j := 0; j < 5*length; j++ {
			diffX := int(j / length)
			diffY := int(i / length)
			newData[i][j] = data[i-diffY*length][j-diffX*length] + diffX + diffY
			if newData[i][j] > 9 {
				newData[i][j] = newData[i][j] - 9
			}
		}
	}
	return newData
}

func djikstra(data [][]int) int {
	length := len(data)
	cavern := &Cavern{make([][]int, length), data, length}
	for y := 0; y < length; y++ {
		cavern.distances[y] = make([]int, length)
		for x := 0; x < length; x++ {
			cavern.distances[y][x] = MaxInt
		}
	}
	cavern.distances[0][0] = 0

	return cavern.calculateShortestPath()
}

func (cavern *Cavern) calculateShortestPath() int {
	maxIndex := cavern.length - 1
	for {
		y, x := cavern.getElementWithSmallestDistanceToSource()

		if y == maxIndex && x == maxIndex {
			break
		}

		cavern.calculateDistanceForNeighbour(x, y, x+1, y)
		cavern.calculateDistanceForNeighbour(x, y, x, y+1)
		cavern.calculateDistanceForNeighbour(x, y, x-1, y)
		cavern.calculateDistanceForNeighbour(x, y, x, y-1)
	}

	return cavern.distances[cavern.length-1][cavern.length-1]
}

func (cavern *Cavern) getElementWithSmallestDistanceToSource() (int, int) {
	min := MaxInt
	var x, y int
	for row, unvisitedRow := range cavern.unvisited {
		for col, unvisitedItem := range unvisitedRow {
			if unvisitedItem == -1 {
				continue
			}
			if cavern.distances[row][col] < min {
				min = cavern.distances[row][col]
				y, x = row, col
			}
		}
	}
	cavern.unvisited[y][x] = -1
	return y, x
}

func (cavern *Cavern) calculateDistanceForNeighbour(x int, y int, newX int, newY int) {
	if newX < 0 || newY < 0 || newY >= cavern.length || newX >= cavern.length || cavern.unvisited[newY][newX] == -1 {
		return
	}
	distance := cavern.distances[y][x] + cavern.unvisited[newY][newX]
	if distance < cavern.distances[newY][newX] {
		cavern.distances[newY][newX] = distance
	}
}
