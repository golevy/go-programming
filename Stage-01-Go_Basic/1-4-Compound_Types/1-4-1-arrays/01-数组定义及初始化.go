package main

import "fmt"

func main() {
	// 全部初始化
	var Numbers1 = [5]int{0, 1, 2, 3, 4}
	fmt.Println("Numbers1 数组中的第一个数是：", Numbers1[0])

	// 部分初始化
	Numbers2 := [5]int{1, 2}
	fmt.Println("Numbers2 数组中的第一个数是：", Numbers2[0])
	fmt.Println("Numbers2 数组中的第五个数是：", Numbers2[4]) // 未初始化的部分，每个位置上的值均为 0

	// 指定某个元素初始化
	Numbers3 := [5]int{2: 6, 4: 8}
	fmt.Println("Numbers3 数组中的第三个数是：", Numbers3[2])
	fmt.Println("Numbers3 数组中的第五个数是：", Numbers3[4])

	// 通过初始化确定数组长度
	Numbers4 := [...]int{8, 5, 7} // [...] 表示不定数组长度，类似于函数的不定参数 (args...)
	fmt.Println("Numbers4 数组的长度是：", len(Numbers4))
	fmt.Println("Numbers4 数组中的第一个数是：", Numbers4[0])
}
