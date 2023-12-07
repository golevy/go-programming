package main

import "fmt"

func InitData(num []int) {
	// 2. 输入要求和的数据，并且存储到切片中
	for i := 0; i < len(num); i++ {
		fmt.Printf("请输入第 %d 个数\n", i+1)
		_, err := fmt.Scan(&num[i])
		if err != nil {
			return
		}
	}
}

func SumAdd(num []int) (sum int) {
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}

	return
}

func main() {
	// 计算出一组整型数据之和
	// 1. 明确要求的数据的个数
	var count int
	fmt.Println("请输入要求和的数据的个数：")
	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println("输入有误！")
		return
	}

	slice := make([]int, count)
	InitData(slice)

	// 3. 求和运算
	sum := SumAdd(slice)
	fmt.Println("求和的结果为：", sum)
}
