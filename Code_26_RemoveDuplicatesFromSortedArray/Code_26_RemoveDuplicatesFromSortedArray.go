package main

func removeDuplicates(nums []int) int {
	/*
		思路：
		首先设置两个指p和q，p指向数组的第一个位置，q指向数组的第二个位置
		如果p和q所指的位置上的数相等，q向后移
		不相等，将q位置上的数复制到p+1的位置上，p和q同时后移
		重复上面的过程，直到q等于数组的长度
		最后返回p+1，即为不重复数组的长度
	*/

	p := 0
	q := 1

	for q < len(nums) {
		if nums[p] == nums[q] {
			q++
		} else {
			nums[p+1] = nums[q]
			p++
			q++
		}
	}
	return p + 1

}
