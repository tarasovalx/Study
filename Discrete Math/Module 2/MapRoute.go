package main

import (
	"fmt"
	"math"

	"github.com/skorobogatov/input"
)

type V struct {
	w    int
	dist int
	x, y int
}

type PriorityQueue struct {
	heap []*V
	cnt  int
}

func Less(pq *PriorityQueue, i, j int) bool {
	h := pq.heap
	return h[i].dist < h[j].dist
}

func Swap(pq *PriorityQueue, i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func Up(pq *PriorityQueue, i int) {
	for ; i != 0 && Less(pq, i, (i-1)/2); i = (i - 1) / 2 {
		Swap(pq, i, (i-1)/2)
	}
}

func Down(pq *PriorityQueue, i int) {
	for i < pq.cnt/2 {
		mi := i*2 + 1
		if i*2+2 < pq.cnt && Less(pq, i*2+2, i*2+1) {
			mi = i*2 + 2
		}
		if Less(pq, i, mi) {
			return
		}
		Swap(pq, i, mi)
		i = mi
	}
}

func Push(pq *PriorityQueue, elem *V) {
	pq.heap = append(pq.heap, elem)
	pq.cnt++
	Up(pq, pq.cnt-1)
}

func Pop(pq *PriorityQueue) *V {
	res := pq.heap[0]
	pq.cnt--
	pq.heap = pq.heap[1:]
	return res
}

var graph [][](*V)

func main() {
	var n int
	var w int
	input.Scanf("%d", &n)

	graph = make([][](*V), n)

	for i := 0; i < n; i++ {
		graph[i] = make([](*V), n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			input.Scanf("%d", &w)
			graph[i][j] = &V{w, math.MaxInt16, i, j}
		}
	}
	dist := dijkstra(n)
	fmt.Printf("%d", dist)
}

func dijkstra(n int) int {
	var next, current *V
	pq := &PriorityQueue{make([](*V), 0, n), 0}
	graph[0][0].dist = 0
	Push(pq, graph[0][0])
	dx := [4]int{0, 0, -1, 1}
	dy := [4]int{-1, 1, 0, 0}

	for pq.cnt != 0 {
		current = Pop(pq)
		x := current.x
		y := current.y
		for i := 0; i < 4; i++ {
			if x+dx[i] < n && x+dx[i] >= 0 && y+dy[i] < n && y+dy[i] >= 0 {
				nextX := x + dx[i]
				nextY := y + dy[i]
				next = graph[nextX][nextY]
				w := graph[nextX][nextY].w
				if w+current.dist < next.dist {
					next.dist = w + current.dist
					Push(pq, next)
				}
			}
		}
	}
	return graph[n-1][n-1].dist + graph[0][0].w
}
