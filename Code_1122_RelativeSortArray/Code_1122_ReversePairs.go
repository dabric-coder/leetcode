package main

import "fmt"

func relativeSortArray(arr1 []int, arr2 []int) []int {

	buckets := make([]int, 1002)
	res := make([]int, len(arr1))


	index := 0
	// 将arr1中的数出现的次数放入对应的桶中
	for i := 0; i < len(arr1); i++ {
		buckets[arr1[i]]++
	}
	//fmt.Println(buckets)

	// 处理arr2中的数，将arr2中的数对应的指定的桶，然后将桶中的计数复原

	for i := 0; i < len(arr2); i++ {
		for buckets[arr2[i]] > 0 {
			res[index] = arr2[i]
			index++
			buckets[arr2[i]]--
		}
	}

	// 最后再处理剩余桶中的计数，也就是处理arr2中未出现的元素
	for i := 0; i < len(buckets); i++ {
		for buckets[i] > 0 {
			res[index] = i
			index++
			buckets[i]--
		}
	}

	return res

}




func main() {
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	res := relativeSortArray(arr1, arr2)
	fmt.Println(res)
}
