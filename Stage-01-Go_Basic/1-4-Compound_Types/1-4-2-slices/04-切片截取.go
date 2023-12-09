package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6} // length == capacity == 6

	// 1. slice[low:high:max]
	// 表示截取切片 [low, high-1] 部分的元素
	// low 代表的是截取的起始位置
	// high 代表的是截取的结束位置（最后结果不包含这个位置上的数，也即是一直截取到这个数的前一个数）
	// max 用来计算容量
	// 0 <= low <= high <= len(slice)
	// high <= max <= cap(slice)
	s1 := slice[1:3:5]
	fmt.Println("截取的切片 s1 的结果为：", s1)
	fmt.Println("截取的切片 s1 的长度为：", len(s1)) // length = high - low
	fmt.Println("截取的切片 s1 的容量为：", cap(s1)) // capacity = max - low
	fmt.Println()

	// 2. slice[:]
	// 表示截取切片中的全部元素
	s2 := slice[:]
	fmt.Println("截取的切片 s2 的结果为：", s2)
	fmt.Println("截取的切片 s2 的长度为：", len(s2)) // length = high - low
	fmt.Println("截取的切片 s2 的容量为：", cap(s2)) // capacity = max - low
	fmt.Println()

	// 3. slice[3:]
	// 表示截取切片 [3, len(slice)-1] 部分的元素
	// 也可以理解为截取切片中第 4 个数到最后
	s3 := slice[3:]
	fmt.Println("截取的切片 s3 的结果为：", s3)
	fmt.Println("截取的切片 s3 的长度为：", len(s3)) // length = high - low
	fmt.Println("截取的切片 s3 的容量为：", cap(s3)) // capacity = max - low
	fmt.Println()

	// 4. slice[:3]
	// 表示截取切片 [0，3-1] 部分的元素
	// 也可以直接理解为截取切片中前 3 个数
	s4 := slice[:3]
	fmt.Println("截取的切片 s4 的结果为：", s4)
	fmt.Println("截取的切片 s4 的长度为：", len(s4)) // length = high - low
	fmt.Println("截取的切片 s4 的容量为：", cap(s4)) // capacity = max - low
}
