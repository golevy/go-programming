package main

import "fmt"

func main() {
	// 输入公交卡当前的余额，只要超过 2 元，就可以上公交车；上车以后如果空座位的数量大于0，就可以坐下。
	// 1. 要求用户输入钱数
	var money float64
	fmt.Println("请输入钱数：")
	_, err := fmt.Scan(&money)
	if err != nil {
		return
	}
	// 2. 对钱数进行判断
	if money > 2 {
		// 3. 如果钱数大于 2 元，条件成立，接着判断座位数
		// 3.1. 要求用户输入座位数
		var count int
		fmt.Println("请输入座位数：")
		_, err = fmt.Scan(&count)
		if err != nil {
			return
		}

		// 3.2. 对输入的座位数进行判断
		if count > 0 {
			fmt.Println("你可以坐下！！！")
		} else {
			fmt.Println("你只能站着！！！")
		}
	} else {
		// 4. 如果钱数不满足，输出“不能上车”
		fmt.Println("你不能上车！！！")
	}
}
