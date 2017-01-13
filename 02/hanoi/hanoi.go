package main

import "fmt"

func Move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}

func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
	Move(n, 1, 2, 3)
}

func Move2(n, from, to int) {
	if n == 0 {
		return
	}
	Move2(n-1, from, 6-from-to)
	fmt.Println(from, "->", to)
	Move2(n-1, 6-from-to, to)
}

func Hanoi2(n int) {
	fmt.Println("Number of disks:", n)
	Move2(n, 1, 2)
}

func main() {
	Hanoi(3)
	Hanoi2(3)
}
