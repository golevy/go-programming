package main

import "fmt"

func AddData(num []int) {
	for i := 0; i < len(num); i++ {
		fmt.Printf("请输入第 %d 个数\n", i+1)
		_, err := fmt.Scan(&num[i])
		if err != nil {
			return
		}
	}
}

func compareValue(num []int) int {
	var maxNumber = num[0]
	for i := 0; i < len(num); i++ {
		if maxNumber < num[i] {
			maxNumber = num[i]
		}
	}

	return maxNumber
}

func main() {
	fmt.Println("请输入需要比较的数据的个数：")
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	slice := make([]int, count)
	AddData(slice)

	maxNumber := compareValue(slice)
	fmt.Println("切片中最大的数是：", maxNumber)
}
