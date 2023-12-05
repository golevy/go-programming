package main

import "fmt"

func main() {
	// var 形式创建切片的初始化
	var slice1 []int
	slice1 = append(slice1, 1, 2, 3, 4, 5, 6)
	fmt.Println("使用 append() 函数对切片 slice1 的初始化结果为：", slice1)

	slice1[1] = 66 // 通过这种方式可以完成对某个值的修改
	fmt.Println("对切片 slice1 中第 2 个数的修改结果为：", slice1[1])
	println()

	// 自动推导类型形式创建切片的初始化
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("切片 slice2 的初始化结果为：", slice2)

	slice2 = append(slice2, 56, 89) // 通过 append() 函数可以在切片原有值的后面继续追加值
	fmt.Println("使用 append() 函数对切片 slice2 的追加结果为：", slice2)

	slice2[2] = 88
	fmt.Println("对切片 slice2 中第 3 个数的修改结果为：", slice2[2])
	println()

	// make() 函数形式创建切片的初始化
	slice3 := make([]int, 3, 10)
	for i := 0; i < len(slice3); i++ {
		slice3[i] = i + 1
	}
	fmt.Println("切片 slice3 的初始化结果为：", slice3)

	// slice3[3] = 56 // 因为数组长度为 3，里面只有三个数，越界不可修改

	slice3 = append(slice3, 66, 99)
	fmt.Println("使用 append() 函数对切片 slice3 的追加结果为：", slice3)

	slice3[3] = 56
	fmt.Println("对切片 slice3 中第 4 个数的修改结果为：", slice3[3])
}
