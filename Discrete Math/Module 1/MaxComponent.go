package main

import (
	"fmt"
)

func main() {
	var n, m, a, b int
	fmt.Scanf("%d\n%d\n", &n, &m)

	graphTable := make([][]int, n)
	componentsLengths := make([]int, n)
	edges := make([]int, n)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		edges = append(edges, a, b)
		graphTable[a] = append(graphTable[a], b)
		graphTable[b] = append(graphTable[b], a)
	}
	compMask := make([]bool, n)
	comp, iMax := findMaxComponent(graphTable, &componentsLengths)

	for _, val := range comp {
		compMask[val] = true
	}
	var g int
	fmt.Printf("graph {\n")
	for i := 0; i < n; i++ {
		//if compMask[i] {
		if componentsLengths[i] == iMax {
			g = i
			fmt.Printf("    %d [color = red]\n", i)
		} else {
			fmt.Printf("    %d\n", i)
		}

		if len(graphTable) > 0 {
			for j := 0; j < len(edges); j += 2 {
				if iMax < graphTable[g][j/2] {
					continue
				}
				if componentsLengths[graphTable[g][j/2]] == iMax { //compMask[edges[i]] || compMask[edges[i+1]] {
					fmt.Printf("    %d -- %d [color = red]\n", edges[j], edges[j+1])
				} else {
					fmt.Printf("    %d -- %d\n", edges[j], edges[j+1])
				}
			}
		}
	}

	fmt.Printf("}")
}

func dfs(start int, graph [][]int, used []bool, component *[]int, edges *int, componentsLengths *[]int, ind int) {
	used[start] = true
	*component = append(*component, start)
	(*componentsLengths)[start] = ind

	for i := 0; i < len(graph[start]); i++ {
		to := graph[start][i]
		*edges++
		if !used[to] {
			dfs(to, graph, used, component, edges, componentsLengths, ind)
		}
	}
}

func findMaxComponent(graph [][]int, mc *[]int) ([]int, int) {
	used := make([]bool, len(graph))
	var maxComponent []int
	var maxVertexes int
	var indexOfMaxComp int
	for i := 0; i < len(graph); i++ {
		var comp []int
		if !used[i] {
			edges := 0
			dfs(i, graph, used, &comp, &edges, mc, i)
			edges /= 2
			if (len(comp) > len(maxComponent)) || ((len(comp) == len(maxComponent)) && (edges > maxVertexes)) {
				indexOfMaxComp = i
				maxVertexes = edges
				maxComponent = make([]int, len(comp))
				copy(maxComponent, comp)
			}
		}
	}
	return maxComponent, indexOfMaxComp
}
