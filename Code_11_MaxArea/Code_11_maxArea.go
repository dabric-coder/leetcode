package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	/*
		简单粗暴的方法：
		两层for循环，数组中的每一个数和其他的数进行面积的计算
	*/
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			max = int(math.Max(float64(max), float64(area)))
		}
	}
	return max
}

func maxArea1(height []int) int {
	/*
		如果左右边界选到最两边的话，肯定是有优势的，但是此时的高度不一定有优势
		此时两边的边界可以往里收缩，那边的边界小那边的往里收缩，直到两边的边界碰头
	*/
	i := 0
	j := len(height) - 1
	max := 0
	for {
		if i == j {
			return max
		}

		area := (j - i) * min(height[i], height[j])
		max = Max(area, max)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := maxArea1(height)
	fmt.Println(maxArea)
}
