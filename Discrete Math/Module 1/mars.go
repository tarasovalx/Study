package main

import (
	"fmt"
	"sort"

	"github.com/skorobogatov/input"
)

func bfs(graphTable [][]int, start, n int, used, part *([]bool), components *([]int), color int) bool {
	queue := make([]int, 0)
	queue = append(queue, start)
	(*used)[start] = true
	(*part)[start] = true
	(*components)[start] = color
	for len(queue) != 0 {
		var vertex int
		vertex, queue = queue[len(queue)-1], queue[:len(queue)-1]
		for i := 0; i < len(graphTable[vertex]); i++ {
			dest := graphTable[vertex][i]
			if !(*used)[dest] {
				(*components)[dest] = color
				(*used)[dest] = true
				(*part)[dest] = !(*part)[vertex]
				queue = append(queue, dest)
			} else if (*part)[dest] == (*part)[vertex] {
				return false
			}
		}
	}
	return true
}

func findParts(graphTable [][]int, n int) (bool, []bool, []int, int) {
	used := make([]bool, n)
	components := make([]int, n)
	componentsCount := 0
	part := make([]bool, n)
	for i, val := range used {
		if !val {
			if !bfs(graphTable, i, n, &used, &part, &components, componentsCount) {
				return false, part, components, componentsCount + 1
			}
			componentsCount++
		}
	}
	return true, part, components, componentsCount
}

func GetMax(old []int, val []int, n int) []int {
	sort.Ints(val)

	if n-len(val) < len(val) {
		return old
	} else if len(old) > len(val) {
		return old
	} else if len(old) < len(val) {
		return val
	}
	for i := 0; i < len(old); i++ {
		if old[i] > val[i] {
			return val
		}
	}
	return old
}

func getSolution(p1 [][]int, p2 [][]int, mask []bool, components, n int, result *[]int, k int) {
	if n == components {
		res := make([]int, 0)
		for i := 0; i < components; i++ {
			if mask[i] {
				for j := 0; j < len(p1[i]); j++ {
					res = append(res, p1[i][j])
				}
			} else {
				for j := 0; j < len(p2[i]); j++ {
					res = append(res, p2[i][j])
				}
			}
		}
		*result = GetMax(*result, res, k)
	} else {
		m := make([]bool, k)
		copy(m, mask)
		getSolution(p1, p2, m, components, n+1, result, k)
		mask[n] = true
		m1 := make([]bool, k)
		copy(m1, mask)
		getSolution(p1, p2, m1, components, n+1, result, k)
	}
}

func main() {
	var n int
	input.Scanf("%d ", &n)
	graphTable := make([][]int, n)

	for i := 0; i < n; i++ {
		str := input.Gets()
		for j := 0; j < n*2; j += 2 {
			if string(str[j]) == "+" {
				graphTable[i] = append(graphTable[i], j/2)
			}
		}
	}

	isParted, parts, components, componentsCount := findParts(graphTable, n)

	firstPart := make([][]int, componentsCount)
	for i := 0; i < componentsCount; i++ {
		firstPart[i] = make([]int, 0)
	}

	seconsPart := make([][]int, componentsCount)
	for i := 0; i < componentsCount; i++ {
		seconsPart[i] = make([]int, 0)
	}

	for i, val := range parts {
		if val {
			firstPart[components[i]] = append(firstPart[components[i]], i)
		} else {
			seconsPart[components[i]] = append(seconsPart[components[i]], i)
		}
	}

	if !isParted {
		fmt.Printf("No solution")
	} else {
		result := make([]int, 0)
		mask := make([]bool, componentsCount)
		getSolution(firstPart, seconsPart, mask, componentsCount, 0, &result, n)

		for _, val := range result {
			fmt.Printf("%d ", val+1)
		}
	}
}
