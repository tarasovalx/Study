/*/*
package main

import (
	"fmt"
)

func bfs(graphTable [][]int, start int, nel int) (alist []int, blist []int) {
	var queue []int
	used := make([]bool, nel)
	part := make([]bool, nel)

	lastMark := false
	//used[start] = true
	check := true
	for check {
		///select next
		check = false
		cntNext := -1
		next := -1
		for j, val := range used {
			if !val {
				if len(graphTable[j]) > cntNext || next == -1 {
					cntNext = len(graphTable[j])
					next = j
					check = true
				}
			}
		}

		if check {
			queue = append(queue, next)
			used[next] = true
			///danger
			part[next] = !lastMark
			lastMark = !lastMark
			//part[next] = true
			for len(queue) != 0 {
				var vertex int
				vertex, queue = queue[len(queue)-1], queue[:len(queue)-1]
				for i := 0; i < len(graphTable[vertex]); i++ {
					dest := graphTable[vertex][i]
					if !used[dest] {
						used[dest] = true
						part[dest] = !part[vertex]
						///
						lastMark = part[dest]
						///
						queue = append(queue, dest)
					} else if part[dest] == part[vertex] {
						fmt.Printf("No solutions")
						return alist, blist
					}
				}
			}
		}
	}

	for i, val := range part {
		if val {
			alist = append(alist, i+1)
		} else {
			blist = append(blist, i+1)
		}
	}

	return alist, blist
}

func main() {
	var n int

	fmt.Scanf("%d ", &n)
	graphTable := make([][]int, n)
	count := make([]int, n)
	var str string

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&str)
			if str == "-" {

			} else {
				graphTable[i] = append(graphTable[i], j)
				count[i]++
			}
		}
	}

	alist, blist := bfs(graphTable, 0, n)

	println("1st group")
	for _, val := range alist {
		fmt.Printf("%d ", val)
	}
	println("")
	println("2nd group")
	for _, val := range blist {
		fmt.Printf("%d ", val)
	}
	/*
		for _, v1 := range graphTable {
			for _, v2 := range v1 {
				fmt.Printf("%d ", v2)
			}
		}*/

//}
