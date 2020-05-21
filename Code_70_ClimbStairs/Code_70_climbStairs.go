package main

import "fmt"

func climbStairs(n int) int {
	/*
		懵逼的时候，两种方式：
			1. 暴力
			2. 先解决简单的问题
				最后找重复的子问题
	*/
	/*
		n=1: f(1) = 1
		n=2: f(2) = 1 + 1
		当n>3时，此时有两种可能：
			1. 从n-1层跨一个台阶上到n层
			2. 从n-2层跨两个台阶上到n层
		n=3: f(3) = f(2) + f(1)
		n=4: f(4) = f(3) + f(2)
		f(n) = f(n-2) + f(n-1)

	*/
	/*
		for i := 0; i <= n; i++ {
			a[i] = a[i-1] + a[i-2]
		}
	*/
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i < n+1; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3

}

func main() {
	n := 10
	fmt.Println(climbStairs(n))
}
