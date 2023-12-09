package main

import "fmt"

func main() {
	// 单独赋值
	var Numbers1 [6]int
	Numbers1[0] = 2
	Numbers1[1] = 0
	Numbers1[2] = 2
	Numbers1[3] = 3
	Numbers1[4] = 1
	Numbers1[5] = 2
	fmt.Println("Numbers1 中的第一个数是：", Numbers1[0])
	fmt.Println("Numbers1 中的第五个数是：", Numbers1[4])

	// 遍历赋值
	var Numbers2 [6]int
	for i := 0; i < len(Numbers2); i++ {
		Numbers2[i] = i + 1
	}
	fmt.Println("Numbers2 中的第六个数是：", Numbers2[5])
}
