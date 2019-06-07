package main

import (
	"fmt"
	"math"

	"github.com/skorobogatov/input"
)

func main() {
	var n, m, k, a, b int
	input.Scanf("%d", &n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = math.MaxInt32
		}
	}
	input.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		input.Scanf("%d%d", &a, &b)
		graph[a][b] = 1
		graph[b][a] = 1
	}
	input.Scanf("%d", &k)

	op := make([]int, k)
	for i := 0; i < k; i++ {
		input.Scanf("%d", &a)
		op[i] = a
	}

	dists := make([][]int, k)
	for i := 0; i < k; i++ {
		dists[i] = dijkstra(graph, op[i], n)
	}
	res := getEqDisted(dists, n)
	for _, val := range res {
		fmt.Printf("%d ", val)
	}
	if len(res) == 0 {
		fmt.Printf("-")
	}
}

func dijkstra(graph [][]int, start, n int) []int {
	dist := make([]int, n)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	visited[start] = true
	current := start
	nextIndex := 0
	for nextIndex != -1 {
		nextIndex = -1
		nextWeight := math.MaxInt32
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			if graph[current][i]+dist[current] < dist[i] {
				dist[i] = graph[current][i] + dist[current]
			}
		}
		for i, val := range visited {
			if !val && (dist[i] < nextWeight) {
				nextIndex = i
				nextWeight = dist[i]
			}
		}
		if nextIndex >= 0 {
			visited[nextIndex] = true
			current = nextIndex
		}
	}

	return dist
}

func getEqDisted(dists [][]int, n int) []int {
	res := make([]int, 0)
	mask := make([]bool, n)
	for i := 0; i < n; i++ {
		mask[i] = true
	}
	for i := 0; i < n; i++ {
		for j := 1; j < len(dists); j++ {
			mask[i] = mask[i] && (dists[j][i] == dists[j-1][i] && dists[j][i] != math.MaxInt32)
		}
	}
	for i, val := range mask {
		if val {
			res = append(res, i)
		}
	}
	return res
}
