package main

import (
	"fmt"
)


func twoSum(nums []int, target int) []int {
	var i, j int
	for i = 0; i < len(nums) - 1; i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	var numMap map[int]int
	numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		_, ok := numMap[diff]
		if  ok && numMap[diff] != i{
			return []int{i, numMap[diff]}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var numMap map[int]int
	numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		_, ok := numMap[diff]
		if ok {
			return []int{numMap[diff], i}
		}
		numMap[nums[i]] = i
	}
	return nil
}


func main() {
	nums := []int{3, 2, 4}
	target := 7
	res := twoSum2(nums, target)
	fmt.Println(res)
}
