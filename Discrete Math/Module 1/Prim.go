package main

import (
	"container/heap"
	"fmt"

	"github.com/skorobogatov/input"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	var n, m, a, b, w int
	input.Scanf("%d\n%d\n", &n, &m)

	graphTable := make([][]int, n)
	EDGES := make([][]int, n, n)

	for i, _ := range EDGES {
		EDGES[i] = make([]int, n, n)
	}

	for i := 0; i < m; i++ {
		input.Scanf("%d %d %d\n", &a, &b, &w)
		EDGES[a][b] = w
		EDGES[b][a] = w
		graphTable[a] = append(graphTable[a], b)
		graphTable[b] = append(graphTable[b], a)
	}

	fmt.Printf("%d", Prim(graphTable, 0, EDGES))
}

func Prim(graph [][]int, start int, edges [][]int) int {
	cnt := 0
	used := make([]bool, len(graph))
	dist := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		dist[i] = -1
	}
	dist[0] = 0
	pq := make(PriorityQueue, 0)
	var buf *Item
	var isFirst bool
	isFirst = true
	v := 0
	used[0] = true
	for {
		for _, u := range graph[v] {
			if (!used[u] && (edges[v][u] < dist[u]) || dist[u] == -1) || isFirst {
				dist[u] = edges[v][u]
				heap.Push(&pq, &Item{u, edges[v][u], -1})
			}
		}
		isFirst = false
		if len(pq) == 0 {
			break
		}
		buf = heap.Pop(&pq).(*Item)
		if !used[buf.value] {
			v = buf.value
			used[v] = true
			cnt += buf.priority
		}
	}
	return cnt
}
