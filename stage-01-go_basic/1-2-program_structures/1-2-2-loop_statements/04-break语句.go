package main

import "fmt"

func main() {
	// 循环遍历输出数字 0-9，当数字等于 3 时终止循环
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println("当前数字是：", i) // 因为 i == 3 时程序就终止了，所以数字 3 不会输出打印
	}
}
