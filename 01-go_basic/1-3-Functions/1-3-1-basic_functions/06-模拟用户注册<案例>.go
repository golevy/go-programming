package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Register() {
	reader := bufio.NewReader(os.Stdin) // 创建了一个新的读取器，用于从标准输入（例如键盘）读取数据

	fmt.Println("请输入用户名：")
	// ReadString 是 bufio.Reader 类型的一种方法，会读取用户的输入，直到用户按下回车键
	userName, _ := reader.ReadString('\n') // 读取并将输入字符串赋值给 userName 变量。下划线 _ 用于忽略错误返回值。
	// strings.TrimSpace 是 strings 包的一个函数，它移除字符串两端的空白字符，包括空格、制表符和换行符
	userName = strings.TrimSpace(userName) // 清理了 userName 字符串两端的任何额外空白，确保只保留用户实际想要输入的部分

	fmt.Println("请输入密码：")
	userPwd, _ := reader.ReadString('\n')
	userPwd = strings.TrimSpace(userPwd)

	fmt.Println("请输入邮箱：")
	userEmail, _ := reader.ReadString('\n')
	userEmail = strings.TrimSpace(userEmail)

	isSuccessful := CheckInfo(userName, userPwd, userEmail)

	if isSuccessful {
		SendEmail()
		fmt.Println("恭喜你，注册成功！！！")
	} else {
		fmt.Println("信息不完整，无法完成用户注册！！！")
	}
}

func CheckInfo(name string, pwd string, email string) bool {
	if name != "" && pwd != "" && email != "" {
		return true
	} else {
	}
	return false
}

func SendEmail() {
	fmt.Println("邮件发送成功！！！")
}

func main() {
	/*
		模拟用户注册，当用户输入完用户名、密码和邮箱后，进行校验。
		如果用户名、密码和邮箱均为空，则提示“信息不能为空，用户注册失败”
		否则，进行邮件发送，并提示“用户注册成功”
	*/
	// 思路：首先注册，然后校验信息，接着发送邮件，最后提示结果
	Register()
}
