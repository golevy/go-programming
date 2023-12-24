package main

import "fmt"

func main() {
	// 1. var 形式的初始化
	var map1 = map[int]string{1: "啦啦啦", 2: "哈哈哈", 3: "略略略"} // key 是唯一的
	fmt.Println("map1 的初始化结果为：", map1)

	// 2. 自动推导类型形式的初始化
	map2 := map[int]string{1: "啦啦啦", 2: "哈哈哈", 3: "略略略"} // key 是唯一的
	fmt.Println("map2 的初始化结果为：", map2)

	// 3. make() 函数形式的初始化
	map3 := make(map[string]int)
	map3["哈哈哈"] = 1
	map3["啦啦啦"] = 2
	map3["哈哈哈"] = 3 // 因为 key 是唯一的，所以使用重复的 key 就相当于修改原来的值
	fmt.Println("map3 的初始化结果为：", map3)
}
