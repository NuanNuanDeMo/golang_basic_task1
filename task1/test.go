package main

import "fmt"

func main() {
	// 出现过一次的数字
	var singlenum int
	// 数组
	arr1 := [...]int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	repeatMap := make(map[int]int)

	for i := 0; i < len(arr1); i++ {
		repeatMap[arr1[i]] += 1
	}

	for key, value := range repeatMap {
		if value == 1 {
			singlenum = key
		}
	}

	fmt.Println("只出现过一次的数字:", singlenum)

}
