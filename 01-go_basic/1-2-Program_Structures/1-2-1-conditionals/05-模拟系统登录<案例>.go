package main

import "fmt"

func main() {
	// Question：提示用户输入用户名，然后再提示输入密码，如果用户输入的用户名是“admin”并且密码是“88888”，则提示正确
	// 如果用户名是 admin 提示密码错误（用户名输入正确，密码输入错误）
	// 如果密码是“88888”，提示用户名错误
	var userName string
	var userPwd string

	fmt.Println("请输入用户名：")
	_, err := fmt.Scan(&userName)
	if err != nil {
		return
	}
	fmt.Println("请输入密码：")
	_, err = fmt.Scan(&userPwd)
	if err != nil {
		return
	}

	if userName == "admin" && userPwd == "88888" {
		fmt.Println("用户名密码正确，登录成功！！！")
	} else if userName == "admin" {
		fmt.Println("密码输入错误！！！")
	} else if userPwd == "88888" {
		fmt.Println("用户名输入错误")
	} else {
		fmt.Println("用户名和密码都输入错误！！！")
	}
}
