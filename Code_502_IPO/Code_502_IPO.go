package main

import (
	"fmt"
)

type Comparable func(i, j interface{}) int 

type BH struct {
	data []interface{}
	heapSize int
	comparable Comparable
}


func NewMaxHeap(comparable Comparable) *BH {
	return &BH{make([]interface{}, 0), 0, comparable}
}

func (bh *BH) siftUp(index int) {
	for bh.comparable(bh.data[index], bh.data[(index-1)/2]) > 0 {
		bh.data[index], bh.data[(index-1)/2] = bh.data[(index-1)/2], bh.data[index]
		index = (index-1)/2
	}
}

func (bh *BH) siftDown(index int) {
	left := index * 2 + 1
	var largest int
	for left < bh.heapSize {
		if left + 1 < bh.heapSize && bh.comparable(bh.data[left+1], bh.data[left]) > 0 {
			largest = left + 1
		} else {
			largest = left
		}

		if bh.comparable(bh.data[index], bh.data[largest]) > 0 {
			largest = index
		} else {
			largest = largest
		}

		if largest == index {
			break
		}
		bh.data[index], bh.data[largest] = bh.data[largest], bh.data[index]
		index = largest
		left = index * 2 + 1
	}

}

func (bh *BH) add(data interface{}) {
	bh.data = append(bh.data, data)
	bh.heapSize++
	bh.siftUp(bh.heapSize-1)
}


func (bh *BH) remove() interface{} {
	if bh.isEmpty() {
		fmt.Println("堆为空")
		return nil
	}
	data := bh.data[0]
	bh.data[0] = bh.data[bh.heapSize-1]
	bh.data = bh.data[0:bh.heapSize-1]
	bh.heapSize--

	bh.siftDown(0)

	return data
}

func (bh *BH) isEmpty() bool {
	if bh.heapSize == 0 {
		return true
	}
	return false
}

func (bh *BH) get() interface{} {
	if bh.isEmpty() {
		fmt.Println("heap is empty")
		return nil
	}
	return bh.data[0]
}

// func swap(arr []int, x, y int) {
// 	arr[x], arr[y] = arr[y], arr[x]
// }

type Project struct {
	cost int
	profit int
}


func maxHeap(i, j interface{}) int {
	return i.(Project).profit - j.(Project).profit
}

func minHeap(i, j interface{}) int {
	return j.(Project).cost - i.(Project).cost
}


func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	projects := make([]Project, 0)
	for k, _ := range Profits {
		projects = append(projects, Project{Capital[k], Profits[k]})
	}
	//fmt.Println(projects)

	minCapitalQ := NewMaxHeap(minHeap)
	maxProfitsQ := NewMaxHeap(maxHeap)

	for _, v := range projects {
		minCapitalQ.add(v)
	}
	//fmt.Println(minCapitalQ.data)

	for i := 0; i < k; i++ {
		for !minCapitalQ.isEmpty() && minCapitalQ.get().(Project).cost <= W {
			maxProfitsQ.add(minCapitalQ.remove())
		}
		//fmt.Println(maxProfitsQ.data...)

		if maxProfitsQ.isEmpty() {
			return W
		}
		//fmt.Println(maxProfitsQ.remove().(Project).profit)
		W += maxProfitsQ.remove().(Project).profit
	}

	return W

}


func main() {
	// k=2, W=0, Profits=[1,2,3], Capital=[0,1,1].
	profits := []int{1,2,3}
	capital := []int{0,1,1}
	w := findMaximizedCapital(2,0,profits, capital)
	fmt.Println(w)

}