package main

import "fmt"

func main() {
	// 将 [9, 8, 7, 6, 5, 4, 3, 2, 1, 0] ==> [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < len(slice)-1; i++ { // 外层控制第几趟比较
		var temp int
		for j := 0; j < len(slice)-1-i; j++ { // 里层控制交换多少次
			if slice[j] > slice[j+1] {
				temp = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}

	fmt.Println("冒泡排序的结果是：", slice)
}
