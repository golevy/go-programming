package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	// 创建一个类型为结构体 Person 类型的切片，其中包含多个结构体元素
	var slice = []Person{
		{
			id:   1,
			name: "Mike",
			age:  18,
		},
		{
			id:   2,
			name: "Lily",
			age:  19,
		},
		{
			id:   3,
			name: "John",
			age:  25,
		},
	}
	fmt.Println(slice) // [{1 Mike 18} {2 Lily 19} {3 John 25}]
	println()

	// 修改切片中结构体成员的值
	slice[0].name = "None"
	fmt.Println("修改切片中结构体成员的值：", slice) // [{1 None 18} {2 Lily 19} {3 John 25}]
	println()

	// 循环遍历
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])     // 例如：{1 None 18}
		fmt.Println(slice[i].age) // 例如：18
	}
	println()

	for k, v := range slice {
		fmt.Printf("结构体切片 slice 中的第 %d 个结构体元素是：%v\n", k, v)
	}
	println()

	// append() 函数的使用
	slice = append(slice, Person{id: 6, name: "Haha", age: 20})
	fmt.Println("使用 append() 函数的结果为：", slice)
}
