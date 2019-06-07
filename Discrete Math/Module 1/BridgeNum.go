package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, m, a, b int
	input.Scanf("%d%d", &n, &m)

	graphTable := make([][]int, n)

	for i := 0; i < m; i++ {
		input.Scanf("%d%d\n", &a, &b)
		graphTable[a] = append(graphTable[a], b)
		graphTable[b] = append(graphTable[b], a)
	}
	cnt := countBridges(graphTable)
	fmt.Printf("%d", cnt)
}

func dfs(start int, graph [][]int, used *[]bool, timesIn *[]int, minTime *[]int, cnt *int, bridges *int, parent int) {
	(*used)[start] = true
	(*timesIn)[start] = *cnt
	(*minTime)[start] = *cnt
	(*cnt)++

	for i := 0; i < len(graph[start]); i++ {
		to := graph[start][i]
		if to == parent {
			continue
		}
		if (*used)[to] {
			(*minTime)[start] = min((*minTime)[start], (*minTime)[to])
		} else {
			dfs(to, graph, used, timesIn, minTime, cnt, bridges, start)
			(*minTime)[start] = min((*minTime)[start], (*minTime)[to])
			if (*minTime)[to] > (*timesIn)[start] {
				(*bridges)++
			}
		}
	}
}

func countBridges(graph [][]int) int {
	bridges := 0
	used := make([]bool, len(graph))
	timesIn := make([]int, len(graph))
	minTime := make([]int, len(graph))
	cnt := 0

	for i, val := range used {
		if !val {
			dfs(i, graph, &used, &timesIn, &minTime, &cnt, &bridges, -1)
		}
	}

	return bridges
}
