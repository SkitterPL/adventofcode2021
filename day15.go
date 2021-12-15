package main

//https://adventofcode.com/2021/day/14

const MaxInt = int(^uint(0) >> 1)

type Item struct {
	key  int
	x    int
	y    int
	risk int
}

func newItem(x int, y int, risk int, key int) *Item {
	return &Item{key, x, y, risk}
}

func key(x int, y int) int {
	return ((x + y) * (x + y + 1) / 2) + x
}

func day15() (int, int) {
	data := fileTo2DIntArray("input/15/input.txt")
	//biggerData := prepareBiggerData(data)
	return djikstra(data), 0
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
	unvisited := make(map[int]*Item, length*length)
	distances := make(map[int]int, length*length)
	previous := make(map[int]*Item, length*length)
	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			key := key(x, y)
			unvisited[key] = newItem(x, y, data[y][x], key)
			distances[key] = MaxInt
		}
	}
	distances[0] = 0

	for len(unvisited) != 0 {
		nextItem := getElementWithSmallestDistToSource(distances, unvisited)

		if nextItem.key == key(length-1, length-1) {
			break
		}

		calculateDistanceForNeighbour(nextItem, nextItem.x+1, nextItem.y, distances, previous, unvisited)
		calculateDistanceForNeighbour(nextItem, nextItem.x, nextItem.y+1, distances, previous, unvisited)
		calculateDistanceForNeighbour(nextItem, nextItem.x-1, nextItem.y, distances, previous, unvisited)
		calculateDistanceForNeighbour(nextItem, nextItem.x, nextItem.y-1, distances, previous, unvisited)
	}

	return distances[key(length-1, length-1)]
}

func getElementWithSmallestDistToSource(distances map[int]int, unvisited map[int]*Item) *Item {
	min := MaxInt
	itemKey := 0
	for key := range unvisited {
		if distances[key] < min {
			min = distances[key]
			itemKey = key
		}
	}
	item := unvisited[itemKey]
	delete(unvisited, itemKey)
	return item
}

func calculateDistanceForNeighbour(item *Item, x int, y int, distances map[int]int, previous map[int]*Item, unvisited map[int]*Item) {
	key := key(x, y)
	if unvisited[key] == nil {
		return
	}
	distance := distances[item.key] + unvisited[key].risk
	if distance < distances[key] {
		distances[key] = distance
		previous[key] = item
	}
}
