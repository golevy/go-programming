package main

import "fmt"

func main() {
	// copy(slice1, slice2)，将 slice2 的值拷贝给 slice1
	slice1 := []int{1, 2}
	slice2 := []int{3, 4, 5, 6, 7}
	slice3 := []int{8, 9}

	res1 := copy(slice1, slice2) // 拷贝的长度为两个切片中长度较小的长度值
	fmt.Println("copy() 函数的返回值为：：", res1)
	fmt.Println("将 slice2 拷贝给 slice1 的结果为：", slice1)
	println()

	res2 := copy(slice2, slice3) // 拷贝的长度为两个切片中长度较小的长度值
	fmt.Println("copy() 函数的返回值为：：", res2)
	fmt.Println("将 slice3 拷贝给 slice2 的结果为：", slice2)
}
