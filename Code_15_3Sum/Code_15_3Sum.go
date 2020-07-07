package main

import (
	"fmt"
	"sort"
)

func threeSum1(nums []int) [][]int {
	set := make(map[[3]int]bool)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j]+nums[k]+nums[i] == 0 {
					tmp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(tmp)
					newtmp := [3]int{tmp[0], tmp[1], tmp[2]}
					_, ok := set[newtmp]
					if !ok {
						set[newtmp] = true
						res = append(res, tmp)
					}
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target, L, R := -nums[i], i+1, len(nums)-1
		for L < R {
			sum := nums[L] + nums[R]
			if sum == target {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				// 判断左右边界是否和下一个位置相等，如果相等，左右边界移向下一个位置，去重
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum < target {
				L++
			} else {
				R--
			}
		}
	}
	return res
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))

}
