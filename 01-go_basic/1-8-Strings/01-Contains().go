package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "HelloGo"
	b := strings.Contains(str, "Go") // 判断某个字符串是否在 str 中存在，若存在则返回 true，否则返回 false
	fmt.Println(b)                   // true
}
