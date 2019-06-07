package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func vis_mealy(d [][]int, f [][]rune, n int, m int, q1 int) {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	fmt.Println("dummy [label = \"\", shape = none]")
	for i := 0; i < len(d); i++ {
		fmt.Printf("%d [shape = circle]\n", i)
	}
	fmt.Printf("dummy -> %d\n", q1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d -> %d [label = \"%c(%c)\"]\n", i, d[i][j], j+97, f[i][j])
		}
	}

	fmt.Println("}")
}

func main() {
	var n, m, q1 int
	input.Scanf("%d\n%d\n%d", &n, &m, &q1)
	D := make([][]int, n)
	F := make([][]rune, n)

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
	vis_mealy(D, F, n, m, q1)
}
