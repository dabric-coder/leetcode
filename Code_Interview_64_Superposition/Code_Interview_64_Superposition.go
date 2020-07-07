package main

import "fmt"

func Superposition1(n int) int {
	res := n
	if n < 1 {
		return 0
	}
	res += Superposition1(n-1)
	return res
}

func Superposition(n int) int {
	var res int
	var sum func (n int) bool
	sum = func (n int) bool {
		res += n
		return n > 1 && sum(n-1)
	}
	sum(n)

	return res
}


func main() {
	fmt.Println(Superposition(10))
}