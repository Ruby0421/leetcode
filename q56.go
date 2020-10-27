package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	target := intervals[0]
	for i := 0; i < len(intervals); i++ {
		if target[1] < intervals[i][0] {
			res = append(res, intervals[i])
			target = intervals[i]
		} else if target[1] < intervals[i][1] {
			target[1] = intervals[i][1]
		}
	}
	return res
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(intervals)
	fmt.Println(merge(intervals))
}
