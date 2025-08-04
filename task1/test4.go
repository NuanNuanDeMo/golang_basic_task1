package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4}
	lenght, newnums := repeatCheck(nums)
	fmt.Println(lenght)
	fmt.Println(newnums)
}

func repeatCheck(nums []int) (int, []int) {
	i, j := 0, 1
	for j < len(nums) {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}

	return i + 1, nums
}
