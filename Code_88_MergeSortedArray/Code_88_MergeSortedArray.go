package main

import "fmt"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}


func main() {
	var nums1 []int
	nums1 = make([]int, 10, 10)
	copy(nums1, []int{2,5,6,7})

	fmt.Println(nums1)
	nums2 := []int{1,2,3,4}

	MergeSortedArray(nums1, 4, nums2, 4)
	fmt.Println(nums1)

}