package main

import (
	"fmt"
	"strings"
)

// 让用户输入一句话，判断这句话中有没有“邪恶”，如果有“邪恶”就替换成“**”，然后输出

func main() {
	fmt.Println("请输入一句话：")
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(str, "邪恶") {
		str = strings.Replace(str, "邪恶", "**", -1)
	}

	fmt.Println(str)
}
