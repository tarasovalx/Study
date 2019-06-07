package main

import (
        "fmt"
        "github.com/skorobogatov/input" 
)

func main() {
    var n, m, a, b int
	var edges []int
	input.Scanf("%d%d", &n, &m)

	graphTable := make([][]int, n)

	for i := 0; i < m; i++ {
		input.Scanf("%d%d\n", &a, &b)
		edges = append(edges, a, b)
		graphTable[a] = append(graphTable[a], b)
		graphTable[b] = append(graphTable[b], a)
	}
	color, colors := findMaxComponent(&graphTable, n)

	fmt.Printf("graph {\n")
	for i := 0; i < n; i++ {
		if colors[i] == color {
			fmt.Printf("    %d [color = red]\n", i)
		} else {
			fmt.Printf("    %d\n", i)
		}
	}
	for i := 0; i < len(edges); i += 2 {
		if colors[edges[i]] == color || color == colors[edges[i+1]] {
			fmt.Printf("    %d -- %d [color = red]\n", edges[i], edges[i+1])
		} else {
			fmt.Printf("    %d -- %d\n", edges[i], edges[i+1])
		}
	}
	fmt.Printf("}")
}

func dfs(start int, graph *[][]int, used *[]int, vertexes *int, edges *int, color int) {
	(*used)[start] = color
	*vertexes++
	for i := 0; i < len((*graph)[start]); i++ {
		*edges++
		to := (*graph)[start][i]
		if (*used)[to] == -1 {
			dfs(to, graph, used, vertexes, edges,color)
		}
	}
}

func findMaxComponent(graph *[][]int, n int) (int, []int) {
	used := make([]int, n)
	for i := 0; i < n; i++ {
		used[i] = -1
	}
	maxComponentColor := -1
	var maxCompLength int
	var maxEdges int
	for i := 0; i < n; i++ {
		if used[i] == -1 {
			vertexes := 0
			edges := 0
			dfs(i, graph, &used, &vertexes, &edges, i)
			if (vertexes > maxCompLength) || ((vertexes == maxCompLength) && (edges > maxEdges)) {
				maxEdges = edges	
				maxCompLength = vertexes
				maxComponentColor = i
			}
		}
	}
	return maxComponentColor, used
}
