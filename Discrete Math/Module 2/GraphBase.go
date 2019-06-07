package main

import (
	"fmt"
	"sort"

	"github.com/skorobogatov/input"
)

type Graph [][]int

var G Graph
var H Graph
var used []bool
var colors []int
var order []int
var components [][]int

func dfs1(v int) {
	used[v] = true
	for _, u := range G[v] {
		if !used[u] {
			dfs1(u)
		}
	}
	order = append(order, v)
}

func dfs2(i, v int) {
	used[v] = true
	components[i] = append(components[i], v)
	colors[v] = i
	for _, u := range H[v] {
		if !used[u] {
			dfs2(i, u)
		}
	}
}

func main() {
	var n, k, a, b int
	input.Scanf("%d", &n)
	input.Scanf("%d", &k)

	G = make([][]int, n)
	H = make([][]int, n)
	order = make([]int, 0, n)
	components = make([][]int, 0)
	colors = make([]int, n)

	for i := 0; i < k; i++ {
		input.Scanf("%d%d", &a, &b)
		G[a] = append(G[a], b)
		H[b] = append(H[b], a)
	}

	used = make([]bool, n)

	for i := 0; i < n; i++ {
		if !used[i] {
			dfs1(i)
		}
	}

	used = make([]bool, n)

	for i := 0; i < n; i++ {
		v := order[n-i-1]
		if !used[v] {
			components = append(components, make([]int, 0))
			dfs2(len(components)-1, v)
		}
	}

	used = make([]bool, len(components))

	res := make([]int, 0)
	for i, component := range components {
		if !used[i] {
			used[i] = true
			res = append(res, components[i][0])
		}
		for _, v := range component {
			for _, u := range G[v] {
				used[colors[u]] = true
			}
		}

	}

	sort.Ints(res)

	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}
