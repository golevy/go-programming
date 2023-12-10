package main

import "fmt"

type Student struct {
	id   int
	name string
	addr string
}

func main() {
	// 定义一个长度为 3 的结构体数组（包含三个 Student 实例的数组）
	// 该结构体数组中包含 3 个结构体，每个结构体类似于 JavaScript 中的对象
	var arr = [3]Student{
		{
			id:   100,
			name: "Hello",
			addr: "No",
		},
		{
			id:   101,
			name: "Hi",
			addr: "Yes",
		},
		{
			id:   102,
			name: "Nice",
			addr: "Oh",
		},
	}
	fmt.Println(arr)         // [{100 Hello No} {101 Hi Yes} {102 Nice Oh}]
	fmt.Println(arr[0].name) // Hello
	println()

	// 修改数组中结构体成员的值
	arr[1].id = 105
	fmt.Println(arr) // [{100 Hello No} {105 Hi Yes} {102 Nice Oh}]
	println()

	// 循环输出结构体数组中的内容
	for i := 0; i < len(arr); i++ {
		fmt.Printf("结构体数组 arr 中的第 %d 个结构体为：%v\n", i, arr[i])             // 例如：{100 Hello No}
		fmt.Printf("结构体数组 arr 中的第 %d 个结构体的 addr 为：%s\n", i, arr[i].addr) // 例如：Yes
	}
	println()

	for k, v := range arr {
		fmt.Println(k) // 例如：0
		fmt.Println(v) // 例如：{100 Hello No}
	}
}
