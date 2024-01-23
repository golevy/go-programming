package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	fmt.Println(strings.Replace(str, "l", "w", -1)) // 用新的字符串替换旧的字符串，第四个参数表示替换的次数，若小于0，则表示全部替换
}
