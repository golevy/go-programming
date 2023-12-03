package main

import "fmt"

func main() {
	// 输出数字 1-5 个，数字 3 除外
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // 当 i == 3时，跳过当前循环中剩余的代码，并继续执行下一次循环
		}
		fmt.Println("当前数字是：", i)
	}
}
