package main

import "fmt"

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return mergeGo(nums, 0, len(nums)-1)

}


func mergeGo(nums []int, L, R int) int {
	if L == R {
		return 0
	}

	mid := L + (R-L) >> 1
	return mergeGo(nums, L, mid) + mergeGo(nums, mid+1, R) + merge(nums, L, mid, R)

}


func merge(nums []int, L, mid, R int) (res int) {
	help := make([]int, R-L+1)

	p1 := L
	p2 := mid+1
	i := 0

	//left := L
	//for right := mid+1; right <= R; right++ {
	//	for left <= mid && nums[left] <= nums[right] {
	//		left++
	//	}
	//	res += mid-left+1
	//}

	for p1 <= mid && p2 <= R {
		if nums[p1] > nums[p2] {
			res += mid-p1+1
		}

		if nums[p1] <= nums[p2] {
			help[i] = nums[p1]
			p1++
			i++
		} else {
			help[i] = nums[p2]
			i++
			p2++
		}
	}

	for p1 <= mid {
		help[i] = nums[p1]
		p1++
		i++
	}

	for p2 <= R {
		help[i] = nums[p2]
		p2++
		i++
	}

	for j := 0; j < len(help); j++ {
		nums[L+j] = help[j]
	}
	return res
}

func main() {
	arr := []int{1,3,2,3,1}
	res := reversePairs(arr)
	fmt.Println(res)
}
