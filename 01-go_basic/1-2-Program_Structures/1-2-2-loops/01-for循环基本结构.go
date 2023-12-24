package main

import "fmt"

func main() {
	// 打印 10 遍 “我爱 Golang！！！”
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Println("我爱 Golang！！！")
	}
	fmt.Printf("从第 %d 次结束循环", i)
}
