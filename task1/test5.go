package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ret := mergeArr(intervals)
	fmt.Println(ret)
}

func mergeArr(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//intervals元素第一个，然后第一个元素
	s := intervals[0][0]
	//intervals元素第一个，然后第二个元素
	e := intervals[0][1]
	size := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > e {
			intervals[size][0] = s
			intervals[size][1] = e
			size++
			s = intervals[i][0]
			e = intervals[i][1]
		} else {
			e = getMax(e, intervals[i][1])
		}
	}
	intervals[size][0] = s
	intervals[size][1] = e
	size++
	return intervals[0:size]
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
