package main

import "fmt"

func main() {
	var arr [2]*int
	var i = 10
	var j = 20

	arr[0] = &i
	arr[1] = &j
	fmt.Println("指针数组中全部元素为：", arr) // [0x140000a4018 0x140000a4020]
	fmt.Println("指针数组中第一个元素为：", *arr[0])
	fmt.Println("指针数组中第一个指针变量指向的变量值为：", *arr[0]) // 通过 * 这个符号即可通过内存地址获取到元素对应的值
	println()

	for i := 0; i < len(arr); i++ {
		fmt.Printf("指针数组中第 %d 个指针变量指向的值为：%d\n", i+1, *arr[i])
	}
	println()

	for k, v := range arr {
		fmt.Println("键：", k)
		fmt.Println("值：", *v)
	}

}
