package main

//https://adventofcode.com/2021/day/14

const MaxInt = int(^uint(0) >> 1)

type Item struct {
	x    int
	y    int
	risk int
}

func newItem(x int, y int, risk int) *Item {
	return &Item{x, y, risk}
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
	unvisited := make([][]*Item, length)
	distances := make([][]int, length)
	previous := make([][]*Item, length)
	for y := 0; y < length; y++ {
		unvisited[y] = make([]*Item, length)
		distances[y] = make([]int, length)
		previous[y] = make([]*Item, length)
		for x := 0; x < length; x++ {
			unvisited[y][x] = newItem(x, y, data[y][x])
			distances[y][x] = MaxInt
			previous[y][x] = nil
		}
	}
	distances[0][0] = 0

	for len(unvisited) != 0 {
		nextItem := getElementWithSmallestDistToSource(distances, &unvisited)

		if nextItem.x == length-1 && nextItem.y == length-1 {
			break
		}

		calculateDistanceForNeighbour(length, nextItem, nextItem.x+1, nextItem.y, &distances, &previous, unvisited)
		calculateDistanceForNeighbour(length, nextItem, nextItem.x, nextItem.y+1, &distances, &previous, unvisited)
		calculateDistanceForNeighbour(length, nextItem, nextItem.x-1, nextItem.y, &distances, &previous, unvisited)
		calculateDistanceForNeighbour(length, nextItem, nextItem.x, nextItem.y-1, &distances, &previous, unvisited)
	}

	return distances[length-1][length-1]
}

func getElementWithSmallestDistToSource(distances [][]int, unvisited *[][]*Item) *Item {
	min := MaxInt
	var item *Item
	for row, unvisitedRow := range *unvisited {
		for col, unvisitedItem := range unvisitedRow {
			if unvisitedItem == nil {
				continue
			}
			if distances[row][col] < min {
				min = distances[row][col]
				item = unvisitedItem
			}
		}
	}
	(*unvisited)[item.y][item.x] = nil
	return item
}

func calculateDistanceForNeighbour(length int, item *Item, x int, y int, distances *[][]int, previous *[][]*Item, unvisited [][]*Item) {
	if x < 0 || y < 0 || y >= length || x >= length || unvisited[y][x] == nil {
		return
	}
	distance := (*distances)[item.y][item.x] + unvisited[y][x].risk
	if distance < (*distances)[y][x] {
		(*distances)[y][x] = distance
		(*previous)[y][x] = item
	}
}
