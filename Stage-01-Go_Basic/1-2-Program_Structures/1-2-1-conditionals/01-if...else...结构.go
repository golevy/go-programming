package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄：")

	_, err := fmt.Scan(&age)
	if err != nil {
		return
	}

	if age >= 18 {
		fmt.Println("恭喜，你已经成年！！！")
	} else {
		fmt.Println("抱歉，你还未成年！！！")
	}
}
