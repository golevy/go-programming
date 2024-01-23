package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "helloGo"
	n := strings.Index(str, "Go") // 判断“Go”在str中出现的位置，注意：位置是从0开始计算的
	fmt.Println(n)
}
