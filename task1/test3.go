package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(addnumber(nums[:]))
}

// 函数
func addnumber(slicearr []int) int64 {
	var resultStr string

	// int转string
	for _, item := range slicearr {
		resultStr += strconv.Itoa(item)
	}

	// string转int64
	i, err := strconv.ParseInt(resultStr, 10, 64)
	if err != nil {
		panic(err)
	}
	var number int64 = int64(i) + 1
	return number
}
