package main

import "fmt"

func main() {
	slice := []int{1, 5, 8, 9, 2, 4}

	// 切片指针的定义
	var p *[]int
	p = &slice
	fmt.Println("切片指针对应切片初始值为：", *p)           // [1 2 3 4 5 6]
	fmt.Println("切片指针对应切片的第一个元素的值为：", (*p)[0]) // 1
	//fmt.Println(p[0]) // 注意，在切片指针中不能像数组指针那样使用 p[0] 便捷访问切片中的值
	println()

	// 修改切片指针对应切片的值
	(*p)[0] = 666
	fmt.Println("修改后切片指针对应切片的值为：", *p)
	println()

	// 循环遍历
	for i := 0; i < len(*p); i++ {
		fmt.Printf("切片指针对应切片的第 %d 个元素的值为：%d\n", i+1, (*p)[i])
	}
	println()

	for k, v := range *p {
		fmt.Println("键：", k)
		fmt.Println("值：", v)
	}
}
