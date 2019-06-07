package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

type Mealy struct {
	D        [][]int
	F        [][]rune
	n, m, q1 int
}

type setItem struct {
	parent *setItem
	self   int
	depth  int
}

func makeSet(q int) *setItem {
	t := &setItem{self: q, depth: 0}
	t.parent = t
	return t
}

func find(x *setItem) *setItem {
	if x.parent == x {
		return x
	} else {
		x.parent = find(x.parent)
		return x.parent
	}
}

func union(x *setItem, y *setItem) {
	rootY := find(y)
	rootX := find(x)
	if rootX.depth < rootY.depth {
		rootX.parent = rootY
	} else {
		rootY.parent = rootX
		if rootX.depth == rootY.depth && rootX != rootY {
			rootX.depth++
		}
	}
}

func makeMealy(n, m, q1 int) *Mealy {
	res := &Mealy{make([][]int, n), make([][]rune, n), n, m, q1}
	for i := 0; i < n; i++ {
		res.D[i] = make([]int, m)
		res.F[i] = make([]rune, m)
	}
	return res
}

func main() {
	var n, m, q1 int
	input.Scanf("%d\n%d\n%d", &n, &m, &q1)
	A1 := makeMealy(n, m, q1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			input.Scanf("%d ", &(A1.D[i][j]))
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			input.Scanf("%c ", &(A1.F[i][j]))
		}
	}
	vis_mealy(cannon(AufenkampHohn(A1)))
}

func split1(A *Mealy) (int, []int) {
	p := make([]*setItem, A.n)
	for i := 0; i < A.n; i++ {
		p[i] = makeSet(i)
	}
	cnt := A.n
	for i := 0; i < A.n; i++ {
		for j := i; j < A.n; j++ {
			if find(p[i]) != find(p[j]) {
				eq := true
				for k := 0; k < A.m; k++ {
					if A.F[i][k] != A.F[j][k] {
						eq = false
						break
					}
				}
				if eq {
					union(p[i], p[j])
					cnt--
				}
			}
		}
	}
	res := make([]int, A.n)
	for q, _ := range p {
		res[q] = find(p[q]).self
	}

	return cnt, res
}

func split(A *Mealy, p []int) (int, []int) {
	cnt := A.n
	q := make([]*setItem, A.n)
	for i := 0; i < A.n; i++ {
		q[i] = makeSet(i)
	}
	for i := 0; i < A.n; i++ {
		for j := i; j < A.n; j++ {
			if p[i] == p[j] {
				if find(q[i]) != find(q[j]) {

					eq := true
					for k := 0; k < A.m; k++ {
						w1 := A.D[i][k]
						w2 := A.D[j][k]
						if p[w1] != p[w2] {
							eq = false
							break
						}
					}
					if eq {
						union(q[i], q[j])
						cnt--
					}
				}
			}
		}
	}

	for i, _ := range p {
		p[i] = find(q[i]).self
	}

	return cnt, p
}

func AufenkampHohn(A *Mealy) *Mealy {
	m, p := split1(A)
	for {
		m1, _ := split(A, p)
		if m == m1 {
			break
		}
		m = m1
	}
	A1 := makeMealy(A.n, A.m, p[A.q1])

	used := make([]bool, A.n)
	for i := 0; i < A.n; i++ {
		for j := 0; j < A.m; j++ {
			A1.D[i][j] = -1
		}
	}

	for i := 0; i < A.n; i++ {
		q := p[i]
		if used[q] {
			continue
		}
		used[q] = true

		for j := 0; j < A.m; j++ {
			A1.D[q][j] = p[A.D[q][j]]
			A1.F[q][j] = A.F[q][j]
		}
	}
	return A1
}

func vis_mealy(A *Mealy) {
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	fmt.Println("dummy [label = \"\", shape = none]")
	for i := 0; i < len(A.D); i++ {
		fmt.Printf("%d [shape = circle]\n", i)
	}
	fmt.Printf("dummy -> %d\n", A.q1)
	for i := 0; i < A.n; i++ {
		for j := 0; j < A.m; j++ {
			fmt.Printf("%d -> %d [label = \"%c(%c)\"]\n", i, A.D[i][j], j+97, A.F[i][j])
		}
	}

	fmt.Println("}")
}

func cannon(A *Mealy) *Mealy {
	res := makeMealy(0, A.m, 0)
	used := make([]bool, A.n)
	new_num := make([]int, A.n)
	pos := 0
	dfs(A, res, used, &pos, A.q1, new_num)
	return res
}

func dfs(A *Mealy, out *Mealy, used []bool, pos *int, q int, new_num []int) {
	if q == -1 || used[q] {
		return
	}
	used[q] = true
	old := *pos
	new_num[q] = *pos
	out.D = append(out.D, make([]int, A.m))
	out.F = append(out.F, make([]rune, A.m))
	out.n++
	(*pos)++
	for i := 0; i < A.m; i++ {
		if A.D[q][i] == -1 {
			continue
		}
		dfs(A, out, used, pos, A.D[q][i], new_num)
		out.D[old][i] = new_num[A.D[q][i]]
		out.F[old][i] = A.F[q][i]
	}
}
