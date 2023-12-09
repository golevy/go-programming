package main

import "fmt"

func main() {
	// Question：有一个英文字符串，统计每个字母出现的次数
	var str = "helloGo"

	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		ch := str[i]      // ch = 'h'
		m[ch] = m[ch] + 1 // 这是 map 的初始化语句：m['h'] = m['h'] + 1，如果键值不存在就创建，存在就更新
		// 上述代码中，利用到了 map 处理不存在键的一个特性
		// 读取或更新一个映射中不存在的键时，Go 会自动使用该键对应值的类型的零值，在这里类型为 int，所以零值就为 0
	}

	for key, value := range m {
		fmt.Printf("字母 %c 出现 %d 次\n", key, value)
	}
}
