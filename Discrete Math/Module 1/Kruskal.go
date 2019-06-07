package main

import (
	"fmt"
	"math"

	"github.com/skorobogatov/input"
)

type Edge struct {
	a, b            *setItem
	priority, index int
}

type setItem struct {
	parent *setItem
	x      int
	y      int
	depth  int
}

type PriorityQueue struct {
	heap []*Edge
	cnt  int
}

func Less(pq *PriorityQueue, i, j int) bool {
	h := pq.heap
	return h[i].priority < h[j].priority
}

func Swap(pq *PriorityQueue, i, j int) {
	h := pq.heap
	h[i], h[j] = h[j], h[i]
}

func Up(pq *PriorityQueue, i int) {
	for ; i != 0 && Less(pq, i, (i-1)/2); i = (i - 1) / 2 {
		Swap(pq, i, (i-1)/2)
	}
}

func Down(pq *PriorityQueue, i int) {
	for i < pq.cnt/2 {
		mi := i*2 + 1
		if i*2+2 < pq.cnt && Less(pq, i*2+2, i*2+1) {
			mi = i*2 + 2
		}
		if Less(pq, i, mi) {
			return
		}
		Swap(pq, i, mi)
		i = mi
	}
}

func Pop(pq *PriorityQueue) *Edge {
	res := pq.heap[0]
	Swap(pq, 0, pq.cnt-1)
	Down(pq, 0)
	pq.cnt--
	return res
}

func Init(pq *PriorityQueue) {
	n := pq.cnt
	for i := n/2 - 1; i >= 0; i-- {
		Down(pq, i)
	}
}

func makeSet(x, y int) *setItem {
	t := &setItem{x: x, y: y, depth: 0}
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

func main() {
	var n, a, b int
	input.Scanf("%d", &n)
	set := make([](*setItem), n)
	for i := 0; i < n; i++ {
		input.Scanf("%d%d", &a, &b)
		set[i] = makeSet(a, b)
	}
	pq := PriorityQueue{make([]*Edge, (n*n-n)/2), 0}

	BuildGraph(set, &pq, n)
	Init(&pq)
	fmt.Printf("%.2f\n", spanningTree(n, &pq))
}

func BuildGraph(set [](*setItem), pq *PriorityQueue, n int) {
	cnt := 0
	var setI, setJ *setItem
	var a, b, dist int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			setI = set[i]
			setJ = set[j]
			a = (setJ.x - setI.x)
			b = (setJ.y - setI.y)
			dist = ((a * a) + (b * b))
			pq.heap[cnt] = &Edge{setI, setJ, dist, -1}
			cnt++
		}
	}
	pq.cnt = cnt
}

func spanningTree(n int, edges *PriorityQueue) float64 {
	totalDist := 0.0
	for i := 0; i < (n - 1); {
		elem := Pop(edges)
		if find(elem.a) != find(elem.b) {
			totalDist += math.Sqrt((float64)(elem.priority))
			union(elem.a, elem.b)
			i++
		}
	}
	return totalDist
}
