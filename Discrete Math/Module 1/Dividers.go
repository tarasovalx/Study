package main

import (
	"fmt"
	"sort"
)

func BuildGraph(n int, N int) {
	var numbers []int
	var primes []int
	var divs []int

	for i := 0; i <= n; i++ {
		numbers = append(numbers, 1)
	}

	for i := 2; i < n; i++ {
		if N%i == 0 {
			if N/i > i {
				divs = append(divs, i, N/i)
			} else {
				divs = append(divs, i)
			}

		}
		if numbers[i] != 0 {
			primes = append(primes, i)
			for j := i + i; j <= n; j += i {
				numbers[j] = 0
			}
		}
	}

	var nextPrime = N
	for _, val := range primes {
		for {
			if nextPrime%val != 0 {
				break
			} else {
				nextPrime /= val
			}
		}
	}

	if nextPrime == N && N == 1 {
		divs = append(divs, N)
	} else if nextPrime != N {
		if nextPrime < n {
			nextPrime = 0
		}
		divs = append(divs, 1, N)
	} else {
		divs = append(divs, 1, N)
	}

	sort.Ints(divs)

	fmt.Printf("graph {\n")
	for _, val := range divs {
		fmt.Printf("    %d\n", val)
	}

	for i := 0; i < len(divs)-1; i++ {
		for j := i + 1; j < len(divs); j++ {
			if (divs[j]%divs[i] == 0) && ((divs[j]/divs[i] == nextPrime) || ((divs[j]/divs[i] < len(numbers)) && (numbers[divs[j]/divs[i]] == 1))) {
				fmt.Printf("    %d -- %d\n", divs[i], divs[j])
			}
		}
	}
	fmt.Printf("}")
}

/*
func main() {
	var n int
	fmt.Scan(&n)
	BuildGraph(int(math.Sqrt(float64(n)))+1, n)
}*/
