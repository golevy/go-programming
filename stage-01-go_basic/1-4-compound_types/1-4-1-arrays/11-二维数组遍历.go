package main

import "fmt"

func main() {
	// for..len... 遍历
	arr1 := [...][3]int{{1, 2, 3}, {5, 6, 8}}
	for j := 0; j < len(arr1); j++ { // 遍历行
		for i := 0; i < len(arr1[0]); i++ { // 遍历列
			fmt.Printf("二维数组 arr1 中第 %d 行数组中的第 %d 个元素是：%d\n", j+1, i+1, arr1[j][i])
		}
	}

	// for...range... 遍历
	arr2 := [...][3]int{{4, 2, 3}, {5, 6, 8}}
	for j, value := range arr2 { // 遍历行
		// fmt.Println(j, value) // 拿到每一行的数组
		for i, data := range value {
			// fmt.Println(data) // 拿到每一行数组中每一项的值
			fmt.Printf("二维数组 arr2 中第 %d 行数组中的第 %d 个元素是：%d\n", j+1, i+1, data)
		}
	}
}
