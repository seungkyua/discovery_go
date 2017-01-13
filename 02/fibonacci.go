package main

import "fmt"

func main() {
	fibonacci(10)
	fmt.Println()
	fibonacci2(10)
}

var memo [100]int

func fibonacci(n int) int {
	if n <= 1 {
		if n == 0 {
			fmt.Print("0, 1")
		}
		return n
	} else {
		if memo[n] > 0 {
			return memo[n]
		}
		memo[n] = fibonacci(n-1) + fibonacci(n-2)
		fmt.Print(", ", memo[n])
		return memo[n]
	}
}

func fibonacci2(n int) {
	p, q := 0, 1
	fmt.Print("0")
	for i := 0; i < n; i++ {
		p, q = q, p+q
		fmt.Print(", ", p)
	}
}
