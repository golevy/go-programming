package main

import "fmt"

func main() {
	// 使用 var 的形式创建
	var slice1 []int // 这种方式不会分配任何内存
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	println()

	// 使用自动推导类型的形式创建
	slice2 := []int{} // 通常不推荐用这种方式创建空切片，因为它会分配内存
	fmt.Println("创建的切片 slice2 为：", slice2)
	fmt.Println("创建的切片 slice2 的长度为：", len(slice2))
	println()

	// 使用 make() 函数创建
	slice3 := make([]int, 3, 5) // 其中 length 不能大于 capacity
	fmt.Println("make() 函数创建的切片 slice3 为：", slice3)
	fmt.Println("make() 函数创建的切片 slice3 的长度为：", len(slice3))
	fmt.Println("make() 函数创建的切片 slice3 的容量为：", cap(slice3))
	println()

	slice4 := make([]int, 3) // capacity 可以省略不写，直接等于 length 的长度
	fmt.Println("make() 函数创建的切片 slice4 为：", slice4)
	fmt.Println("make() 函数创建的切片 slice4 的长度为：", len(slice4))
	fmt.Println("make() 函数创建的切片 slice4 的容量为：", cap(slice4))
}
