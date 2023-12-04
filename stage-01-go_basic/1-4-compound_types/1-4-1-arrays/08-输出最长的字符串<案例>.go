package main

import "fmt"

func main() {
	// 有一个字符串数组：{"马龙", "马克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特"}
	// 请输出最长的字符串
	names := [...]string{"马龙", "马克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特"}
	var maxString = names[0]

	for i := 0; i < len(names); i++ {
		if len(names[i]) > len(maxString) {
			maxString = names[i]
		}
	}

	fmt.Println("最长的字符串是：", maxString)
}
