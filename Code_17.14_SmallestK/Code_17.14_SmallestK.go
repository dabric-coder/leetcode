package main

import (
	"fmt"
	"math/rand"
)


// 方法一 使用大根堆
func smallestK1(arr []int, k int) []int {

	for i := 0; i < len(arr); i++ {
		if i < k {
			sitUp(arr, i)
			//fmt.Println(arr)
		} else if arr[i] < arr[0] {
			replace(arr, arr[i], k)
			//fmt.Println(arr)
		}
	}
	fmt.Println(arr)
	result := arr[:k]
	fmt.Println(result)
	return result
}


func sitUp(arr []int, index int) {

	//for arr[index] > arr[(index-1)/2] {
	//	arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
	//	index = (index-1)/2
	//}
	data := arr[index]
	for index > 0 {
		parentIndex := (index-1) >> 1
		parent := arr[parentIndex]

		if data <= parent {
			break
		}
		arr[index] = parent
		index = parentIndex
	}
	arr[index] = data

}

func siftDown(arr []int, index int, heapSize int) {
	leftChild := index * 2 + 1
	var largest int
	for leftChild < heapSize {
		if leftChild+1 < heapSize && arr[leftChild+1] > arr[leftChild] {
			largest = leftChild+1
		} else {
			largest = leftChild
		}

		if arr[index] > arr[largest] {
			index = largest
		}

		if index == largest {
			break
		}

		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
		leftChild = index * 2 + 1
	}

}

func replace(arr [] int, data int, k int) {
	// 删除堆顶元素同时插入一个新元素
	arr[0] = data
	siftDown(arr, 0, k)
}

func smallestK(arr []int, k int) []int {
	// 使用快排的思想
	quickSortGo(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSortGo(arr []int, L, R, k int) {
	if L < R {
		swap(arr, L+ rand.Intn(R-L+1), R)
		p := partition(arr, L, R)
		if p[0] == k {
			return
		} else if p[0] > k {
			quickSortGo(arr, L, p[0]-1, k)
		} else {
			quickSortGo(arr, p[1]+1, R, k)
		}
	}

}


func partition(arr []int, L, R int) (p *[2]int){
	less := L - 1
	more := R + 1
	cur := L
	//povit := arr[R]
	for cur < more {
		if arr[cur] < arr[R] {
			swap(arr, cur, less+1)
			cur++
			less++
		} else if arr[cur] > arr[R] {
			swap(arr, cur, more - 1)
			more--
		} else {
			cur++
		}
	}
	return &[2]int{less+1, more-1}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func main() {
	arr := []int{0,0,0,1,2,2,3,7,6,1}
	result := smallestK(arr, 3)
	fmt.Println(result)
	//
}