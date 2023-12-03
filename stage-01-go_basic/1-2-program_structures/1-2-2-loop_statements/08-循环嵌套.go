package main

import "fmt"

func main() {
	// 每天说 3 遍“我爱 Golang！！！”，一共说 7 天
	for j := 1; j <= 7; j++ {
		fmt.Printf("这是第 %d 天说的：\n", j)
		for i := 1; i <= 3; i++ {
			fmt.Println("我爱 Golang！！！")
		}
	}
	// 循环嵌套的本质就是，外层每循环一次，里层都要进行一遍完整的循环
}
