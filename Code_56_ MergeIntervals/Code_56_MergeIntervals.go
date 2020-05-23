package main

import (
	"fmt"
	"sort"
)

type Ints [][]int

func (s Ints) Len() int {
	return len(s)
}

func (s Ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Ints) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(Ints(intervals))

	merged := make([][]int, 0)
	if len(intervals) < 1 {
		return merged
	}

	merged = append(merged, intervals[0])


	j := 0
	for i := 1; i < len(intervals); i++ {
		if merged[j][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
			fmt.Println("append", merged)
			j++
		} else {
			max := func(i, j int) int {
				if i > j {
					return i
				} else {
					return j
				}
			}(merged[j][1], intervals[i][1])
			merged[j][1] = max
		}
	}

	return merged
}


func main() {
	arr := [][]int{{1,4},{0,4}}
	res := merge(arr)
	fmt.Println(res)
}



