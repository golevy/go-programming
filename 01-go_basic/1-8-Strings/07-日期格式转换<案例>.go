package main

import (
	"fmt"
	"strings"
)

// 让用户输入一个日期格式，如：2024-01-24，输出日期为：2024年1月24日

func main() {
	fmt.Println("请输入日期，格式：年-月-日")
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := strings.Split(str, "-") // 把字符串按照指定的分隔符 "-" 进行分割，返回slice(切片)，意味着函数会在每个连字符的位置将字符串分割开

	fmt.Println(s[0] + "年" + s[1] + "月" + s[2] + "日")
}
