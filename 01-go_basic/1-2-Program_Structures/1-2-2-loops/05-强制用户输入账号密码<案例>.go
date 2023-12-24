package main

import "fmt"

func main() {
	// 要求用户输入用户名和密码，只要不是 admin、88888，就一直提示用户：密码错误，请重新输入
	var userName string
	var userPwd string

	for { // for 循环后面不加任何条件，即表示是死循环（一直循环，无法终止）
		fmt.Println("请输入用户名：")
		_, err := fmt.Scan(&userName)
		if err != nil {
			return
		}
		fmt.Println("请输入密码")
		_, err = fmt.Scan(&userPwd)
		if err != nil {
			return
		}

		if userName == "admin" && userPwd == "88888" {
			fmt.Println("恭喜你，成功登录！！！")
			break
		}

		fmt.Println("用户名密码错误，请重新输入！！！")
	}
}
