package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(judgestr("([])"))
}

// 函数
func judgestr(str string) bool {
	var resultStr string = str
	for strings.Contains(resultStr, "()") || strings.Contains(resultStr, "{}") || strings.Contains(resultStr, "[]") {
		if strings.Contains(str, "(") {
			resultStr = strings.Replace(resultStr, "()", "", -1)
		}
		if strings.Contains(str, "{") {
			resultStr = strings.Replace(resultStr, "{}", "", -1)
		}
		if strings.Contains(str, "[") {
			resultStr = strings.Replace(resultStr, "[]", "", -1)
		}
	}
	return len(resultStr) == 0
}
