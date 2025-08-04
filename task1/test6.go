package main

import "fmt"

/*
	两数之和
*/
func main() {
	nums := []int{3, 2, 4}
	arr := addnums(nums[:], 6)
	fmt.Println(arr)
}

func addnums(nums []int, target int) [2]int {
	var indexarr [2]int
outloop:
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indexarr[0] = i
				indexarr[1] = j
				break outloop
			}
		}
	}
	return indexarr
}
