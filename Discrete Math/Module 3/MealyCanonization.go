package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

var n, m, q1 int
var D [][]int
var F [][]rune

var D_out [][]int
var F_out [][]rune
var used []bool
var new_num []int
var pos int

func dfs(q int) {
	if used[q] {
		return
	}
	used[q] = true
	old := pos
	new_num[q] = pos
	D_out = append(D_out, make([]int, m))
	F_out = append(F_out, make([]rune, m))
	pos++
	for i := 0; i < m; i++ {
		dfs(D[q][i])
		D_out[old][i] = new_num[D[q][i]]
		F_out[old][i] = F[q][i]
	}
}

func main() {
	input.Scanf("%d\n%d\n%d", &n, &m, &q1)
	used = make([]bool, n)
	new_num = make([]int, n)
	D = make([][]int, n)
	F = make([][]rune, n)
	D_out = make([][]int, 0)
	F_out = make([][]rune, 0)

	for i := 0; i < n; i++ {
		D[i] = make([]int, m)
		F[i] = make([]rune, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			input.Scanf("%d ", &D[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			input.Scanf("%c ", &F[i][j])
		}
	}

	dfs(q1)

	fmt.Printf("%d\n%d\n%d\n", len(D_out), m, 0)

	for i := 0; i < len(D_out); i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", D_out[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < len(F_out); i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c ", F_out[i][j])
		}
		fmt.Println()
	}
}
