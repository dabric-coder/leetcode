package main

import "fmt"

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return smallSum(nums, 0, len(nums)-1)
}


func smallSum(slice []int, L, R int) (res int) {
	if L == R {
		return 0
	}
	mid := L + (R - L) >> 1

	return smallSum(slice, L, mid) + smallSum(slice, mid + 1, R) + merge(slice, L, mid, R)

}

func merge(slice []int, L, mid, R int) (res int) {
	var help []int = make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid +1

	left := L
	for right := mid+1; right<=R; right++ {
		for left <= mid && slice[left] <= slice[right] * 2 {
			left++
		}
		res += mid-left+1
	}

	for p1 <= mid && p2 <= R {

		if slice[p1] < slice[p2] {
			help[i] = slice[p1]
			p1++
		} else {
			help[i] = slice[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		help[i] = slice[p1]
		i++
		p1++
	}

	for p2 <= R {
		help[i] = slice[p2]
		i++
		p2++
	}

	for j := 0; j < len(help); j++ {
		slice[L+j] = help[j]
	}
	return
}


func main() {
	arr := []int{1,3,2,3,1}
	res := reversePairs(arr)
	fmt.Println(res)
}
