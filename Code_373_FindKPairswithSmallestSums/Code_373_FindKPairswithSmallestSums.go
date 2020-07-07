package main

import (
	"fmt"
	"math"
)


type Node struct {
	index1 int
	index2 int
	sum int
}

func comparator (i, j interface{}) int {
	return j.(Node).sum - i.(Node).sum
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	if (len(nums1) <= 0 || len(nums2) <=0 || k < 1) {
		return nil
	}

	h := NewHeap(comparator)
	k = int(math.Min(float64(k), float64(len(nums1)*len(nums2))))
	res := make([][]int, 0)  // 存放前k小的数组
	resIndex := 0
	i1 := 0
	i2 := 0

	set := make([][]bool, len(nums1))

	for i := range set {
		set[i] = make([]bool, len(nums2))
	}

	// 生成一个小根堆

	h.add(Node{i1,i2, nums1[i1]+nums1[i2]})
	set[i1][i2] = true


	for resIndex < k {
		curNode := h.poll()
		if curNode != nil {
			newCurNode := curNode.(Node)
			i1 = newCurNode.index1
			i2 = newCurNode.index2
		}


		res = append(res, []int{nums1[i1], nums2[i2]})
		resIndex++

		if i1+1 < len(nums1) {
			if set[i1+1][i2] == false {
				set[i1+1][i2] = true
				h.add(Node{i1+1,i2, nums1[i1+1]+nums2[i2]})
			}
		}

		if i2+1 < len(nums2) {
			if set[i1][i2+1] == false {
				set[i1][i2+1] = true
				h.add(Node{i1,i2+1, nums1[i1]+nums2[i2+1]})
			}
		}

	}


	return res
}



type Comparator func (i, j interface{}) int
// 自己实现堆结构
type Heap struct {
	data []interface{}
	size int
	comparator Comparator
}


func NewHeap(comparator Comparator) *Heap {
	return &Heap{nil, 0, comparator}
}

func (h *Heap) add(data interface{}) {
	h.data = append(h.data, data)
	h.size++
	h.siftUp(h.size-1)
}

func (h *Heap) siftUp(index int) {

	data := h.data[index]
	for index > 0 {
		parentIndex := (index - 1) >> 1
		parent := h.data[parentIndex]

		if h.comparator(data, parent) <= 0 {
			break
		}
		h.data[index] = parent
		index = parentIndex

	}
	h.data[index] = data
}

func (h *Heap) poll() interface{} {
	if h.size == 0 {
		return nil
	}

	res := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data = h.data[:h.size-1]
	h.size--

	h.siftDown(0)
	return res
}

func (h *Heap) siftDown(index int) {
	left := index * 2 + 1
	var largest int

	for left < h.size {
		if left + 1 < h.size && h.comparator(h.data[left+1], h.data[left]) > 0 {
			largest = left + 1
		} else {
			largest = left
		}

		if h.comparator(h.data[index], h.data[largest]) > 0 {
			largest = index
		}

		if index == largest {
			break
		}

		h.data[index], h.data[largest] =  h.data[largest], h.data[index]
		index = largest
		left = index * 2 + 1
	}

}




func main() {
	arr1 := []int{1, 7, 11}
	arr2 := []int{2, 4, 6}
	k := 3

	fmt.Println(kSmallestPairs(arr1, arr2, k))


}
