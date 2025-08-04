package main

import "fmt"

/*
	最长公共前缀
*/
func main() {
	arr := [3]string{"flower", "flow", "floight"}
	arr1 := [3]string{"dog", "racecar", "car"}

	fmt.Println("公共前缀:", judge(arr[:]))
	fmt.Println("公共前缀:", judge(arr1[:]))
}

// 函数
func judge(arrSlice []string) string {
	// 返回参数
	var prestr string
	// 最短的字符串长度
	var minLength int = 0

	for _, item := range arrSlice {
		if minLength == 0 || minLength > len(item) {
			minLength = len(item)
		}
	}

	firstElement := arrSlice[0]
	for i := 0; i < minLength; i++ {
		var count int = 0
		singlechar := firstElement[i]
		for j := 1; j < len(arrSlice); j++ {
			if string(singlechar) == string(arrSlice[j][i]) {
				count++
			}
		}
		if count == len(arrSlice)-1 {
			prestr += string(singlechar)
		}
	}
	return prestr
}
