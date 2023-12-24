package main

import "fmt"

func main() {
	// 1. 通过 key 获取值
	var map1 = map[int]string{1: "Go"}
	fmt.Println("通过 key 获取 map 中的值的结果是：", map1[1])
	fmt.Println("通过 key 获取 map 中的值的结果是：", map1[2]) // key=2 在原 map 不存在对应的值，由于映射的值类型是 string，所以零值是空字符串 ""
	println()

	// 2. 通过 key 获取值时，判断值是否存在
	value1, isExisted := map1[1] // 第一个变量 value1 为键所对应的值，第二个变量 isExisted 为判断键所对应的值是否存在
	fmt.Println("通过 key 获取 map1 中对应的 value1 为：", value1)
	fmt.Println("通过 key 获取的 map1 中的 value1 是否存在：", isExisted)

	value2, isExisted := map1[2]
	fmt.Println("通过 key 获取 map1 中对应的 value2 为：", value2)
	fmt.Println("通过 key 获取的 map1 中的 value2 是否存在：", isExisted)
	println()

	// 3. 通过循环方式获取值
	// 注意：不能使用 for...len 的形式直接遍历 map
	// 这是因为 Go 中的 map 是一个无序的键值对集合，而 len 函数只能返回 map 中键值对的数量，并不能用于直接索引或遍历
	for key, value := range map1 { // 第一个参数 key 对应键，第二个参数 value 为 key 所对应的值
		fmt.Println("通过循环方式获取值的键为：", key)
		fmt.Println("通过循环方式获取的值为：", value)
	}

	// 4. 通过 key 删除某个值
	delete(map1, 1)
	fmt.Println("通过 key 删除 map1 中的值后 map1 为：", map1)
}
