package main

import (
	"container/heap"
)

//https://adventofcode.com/2021/day/15

const MaxInt = int(^uint(0) >> 1)

type CavePosition struct {
	x        int
	y        int
	distance int
	index    int
}

type PriorityQueue []*CavePosition

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*CavePosition)
	item.index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) update(item *CavePosition, priority int) {
	item.distance = priority
	heap.Fix(pq, item.index)
}

type Cavern struct {
	heap      *PriorityQueue
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
	queue := make(PriorityQueue, length*length)
	cavern := &Cavern{&queue, make([][]int, length), data, length}
	for y := 0; y < length; y++ {
		cavern.distances[y] = make([]int, length)
		for x := 0; x < length; x++ {
			cavern.distances[y][x] = MaxInt
			(*cavern.heap)[y*length+x] = &CavePosition{x: x, y: y, distance: MaxInt, index: y*length + x}
		}
	}
	(*cavern.heap)[0].distance = 0
	cavern.distances[0][0] = 0
	heap.Init(&queue)

	return cavern.calculateShortestPath()
}

func (cavern *Cavern) calculateShortestPath() int {
	maxindex := cavern.length - 1
	for {
		cp := heap.Pop(cavern.heap).(*CavePosition)
		y := cp.y
		x := cp.x
		cavern.unvisited[y][x] = -1

		if y == maxindex && x == maxindex {
			break
		}

		cavern.calculateDistanceForNeighbour(x, y, x+1, y)
		cavern.calculateDistanceForNeighbour(x, y, x, y+1)
		cavern.calculateDistanceForNeighbour(x, y, x-1, y)
		cavern.calculateDistanceForNeighbour(x, y, x, y-1)
	}

	return cavern.distances[maxindex][maxindex]
}

func (cavern *Cavern) calculateDistanceForNeighbour(x int, y int, newX int, newY int) {
	if newX < 0 || newY < 0 || newY >= cavern.length || newX >= cavern.length || cavern.unvisited[newY][newX] == -1 {
		return
	}
	distance := cavern.distances[y][x] + cavern.unvisited[newY][newX]
	if distance < cavern.distances[newY][newX] {
		cavern.distances[newY][newX] = distance
		item := &CavePosition{x: newX, y: newY, distance: distance}
		heap.Push(cavern.heap, item)
		cavern.heap.update(item, distance)
	}
}
