package main

import "fmt"

func main() {
	// 从一个整数数组中取出最大的整数、最小的整数，并且求总和，求平均值
	var nums = [5]int{3, 6, 5, 8, 9}

	var minNum = nums[0]
	var maxNum = nums[0]
	var sum = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}

		if nums[i] < minNum {
			minNum = nums[i]
		}

		sum = sum + nums[i]
	}

	fmt.Println("最大的整数是：", maxNum)
	fmt.Println("最小的整数是：", minNum)
	fmt.Println("总和为：", sum)
	fmt.Printf("平均值为：%.2f", float64(sum/5))
}
