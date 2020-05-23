package main

import "fmt"

func isAnagram(s string, t string) bool {

	// 当两个字符串的长度不相等时，直接返回false
	if len(s) != len(t) {
		return false
	}
	// 创建26个桶
	buckets := make([]int, 26)

	// 同时遍历s和t，对于s中的每个字符出现后，在对应的桶中+1；t中的每个字符出现后，对应的桶中-1

	for i := 0; i < len(s); i++ {
		buckets[s[i] % 97]++
		buckets[t[i] % 97]--
	}

	count := 0
	for i := 0; i < len(buckets); i++ {
		if buckets[i] == 0 {
			count++
		}
	}

	if count == len(buckets) {
		return true
	} else {
		return false
	}
}




func main() {
	s := "rat"
	t := "car"
	fmt.Println(isAnagram(s, t))


}
