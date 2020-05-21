package main

import "fmt"

func moveZeroes(nums []int) {
	k := -1 // [0..k]范围表示整个数组的非0范围，初始时这个范围不存在

	for i := 0; i < len(nums); i++ { // 保证[0...i]上为非0元素
		if nums[i] != 0 {
			nums[k+1] = nums[i]
			k++
		}
	}

	for i := k + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	/*
		使用[0,k]来表示nums中大于0的部分，起初大于0的部分不存在，另k=-1；
		遍历nums：
			如果当前cur指向的元素大于0则和大于0的区域的下一个元素进行交换，大于0的区域增加；
			如果当前的元素等于0，cur++
	*/
	k := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			swap(nums, k+1, i)
			k++
		}
	}
}

// 继续优化
// 对于nums中的元素都是大于0的情况下，上面的算法流程中所有的当前的元素都会和自己进行一个交换
// 于是进行下面的优化操作
func moveZeroes2(nums []int) {
	/*
		使用[0,k]来表示nums中大于0的部分，起初大于0的部分不存在，另k=-1；
		遍历nums：
			如果当前cur指向的元素大于0则和大于0的区域的下一个元素进行交换，大于0的区域增加；
			如果当前的元素等于0，cur++
	*/
	k := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if k+1 != i {
				swap(nums, k+1, i)
				k++
			} else {
				k++
			}
		}
	}
}

func swap(nums []int, x, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}

func main() {
	var nums []int = []int{10, 0, 11, 0, 1, 3, 3, 0}
	moveZeroes2(nums)
	fmt.Println(nums)
}
