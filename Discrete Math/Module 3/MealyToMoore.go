package main

import "fmt"

var inputAlphabet []string
var outputAlphabet []string

type Mealy struct {
	D        [][]int
	F        [][]string
	n, m, q1 int
}

var m, k, n int

func main() {

	fmt.Scan(&m)
	inputAlphabet = make([]string, m)
	for i := range inputAlphabet {
		fmt.Scan(&inputAlphabet[i])
	}

	fmt.Scan(&k)
	outputAlphabet = make([]string, k)
	for i := range outputAlphabet {
		fmt.Scan(&outputAlphabet[i])
	}

	fmt.Scan(&n)
	A1 := makeMealy(n, m, 0)
	readMealy(A1)
	mealyToMoore(A1)
}

func mealyToMoore(A *Mealy) {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	used := make(map[int]map[string]int)
	cnt := 1
	for i := 0; i < A.n; i++ {
		for j := 0; j < A.m; j++ {
			if used[A.D[i][j]] == nil || used[A.D[i][j]][A.F[i][j]] == 0 {
				if used[A.D[i][j]] == nil {
					used[A.D[i][j]] = make(map[string]int)
				}
				used[A.D[i][j]][A.F[i][j]] = cnt
				fmt.Printf("%d [label = \"(%d,%s)\"]\n", cnt, A.D[i][j], A.F[i][j])
				cnt++
			}
		}
	}

	for i := 0; i < A.n; i++ {
		for j := 0; j < A.m; j++ {
			for q := 0; q < k; q++ {
				if used[A.D[i][j]][A.F[i][j]] != 0 && used[i][outputAlphabet[q]] != 0 {
					fmt.Printf("%d->%d [label = \"%s\"]\n", used[i][outputAlphabet[q]], used[A.D[i][j]][A.F[i][j]], inputAlphabet[j])
				}
			}
		}
	}
	fmt.Println("}")
}

func readMealy(A *Mealy) {
	for i := 0; i < A.n; i++ {
		for j := 0; j < A.m; j++ {
			fmt.Scan(&(A.D[i][j]))
		}
	}

	for i := 0; i < A.n; i++ {
		for j := 0; j < A.m; j++ {
			fmt.Scan(&(A.F[i][j]))
		}
	}
}

func makeMealy(n, m, q1 int) *Mealy {
	res := &Mealy{make([][]int, n), make([][]string, n), n, m, q1}
	for i := 0; i < n; i++ {
		res.D[i] = make([]int, m)
		res.F[i] = make([]string, m)
	}
	return res
}
