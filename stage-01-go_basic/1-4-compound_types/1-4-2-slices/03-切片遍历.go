package main

import "fmt"

func main() {
	// for...len.. 遍历
	slice1 := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("切片 slice1 中的第 %d 个元素的值是：%d\n", i+1, slice1[i])
	}
	println()

	// for...range... 遍历
	var slice2 = []int{1, 2, 3, 7, 8, 9}
	for i, v := range slice2 {
		fmt.Printf("切片 slice2 中的第 %d 个元素的值是：%d\n", i+1, v)
	}
}
