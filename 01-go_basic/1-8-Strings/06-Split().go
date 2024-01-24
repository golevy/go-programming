package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello@World@Go"
	s := strings.Split(str, "@") // 把字符串按照指定的分隔符 "@" 进行分割，返回slice(切片)
	fmt.Println(s)               // [Hello World Go]
}
