package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))


	// 先计算数组中除了当前元素外的，当前元素左边的乘积
	k := 1
	for i := 0; i < len(nums); i++ {
		res[i] = k     // k 为除了当前元素外，当前元素右边元素的乘积
		k *= nums[i]   // 更新k的值，k用来记录下一个元素左边元素的乘积
	}

	// 在计算数组中除了当前元素外，当前元素右边的乘积与左边乘积的乘积
	k = 1
	for i := len(nums) -1; i >= 0; i-- {
		res[i] = res[i] * k // k 为除了当前元素外，当前元素右边元素的乘积
		k *= nums[i]		// 更新k的值，k用来记录下一个元素右边元素的乘积
	}

	return res
}


func main() {
	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))
}