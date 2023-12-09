package main

import "fmt"

func GetArray(arr [5]int) {
	fmt.Println("下面是 GetArray() 函数的输出结果：")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	arr[0] = 66 // 对接收的数组进行赋值不生效，原因是其指向内存地址不一样，无法赋值
}

func main() {
	Numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("下面是 Numbers 的每一项的初始值：")
	for _, value := range Numbers {
		fmt.Println(value)
	}

	GetArray(Numbers)
}
