package main

import "fmt"

func main() {
	// for...len.. 遍历
	var Numbers1 = [5]int{8, 5, 7, 6, 1}
	fmt.Println("下面是 Numbers1 数组的遍历结果：")
	for i := 0; i < len(Numbers1); i++ {
		fmt.Println(Numbers1[i])
	}

	// for...range.. 遍历
	var Numbers2 = [5]int{1, 3, 7, 4, 0}
	fmt.Println("下面是 Numbers2 数组的遍历结果：")
	for _, value := range Numbers2 {
		fmt.Println(value)
	}
}
