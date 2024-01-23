package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello@World@Go"
	s := strings.Split(str, "@") // 把字符串按照sep(@)分割，返回slice(切片)
	fmt.Println(s)               // [Hello World Go]
}
