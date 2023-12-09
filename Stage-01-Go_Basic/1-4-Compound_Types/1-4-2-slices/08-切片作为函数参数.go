package main

import "fmt"

func Iterate(num []int) {
	for i := 0; i < len(num); i++ {
		fmt.Printf("切片 num 的第 %d 个数为：%d\n", i, num[i])
	}
	println()
}

func Init(num []int) {
	for i := 0; i < len(num); i++ {
		num[i] = i + 1
	}
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	Iterate(slice1)

	slice2 := make([]int, 10)
	fmt.Println("切片 slice2 的初始结果为：", slice2)
	Init(slice2) // 在函数中修改切片的值也会对应修改原切片中的值，因为所有元素的内存地址指向是一致的
	fmt.Println("经处理后 slice2 的结果为：", slice2)
}
