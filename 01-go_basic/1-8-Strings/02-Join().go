package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"hello", "world", "go"}
	str := strings.Join(s, "|") // 以“|”对切片中的内容进行连接
	fmt.Println(str)            // hello|world|go
}
