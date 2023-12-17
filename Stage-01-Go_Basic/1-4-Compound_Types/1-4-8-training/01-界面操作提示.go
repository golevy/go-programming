package main

import "fmt"

func scanNum() {
	fmt.Println("添加联系人信息，请按1")
	fmt.Println("删除联系人信息，请按2")
	fmt.Println("查询联系人信息，请按3")
	fmt.Println("编辑联系人信息，请按4")

	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
	switchType(num)
}

func switchType(n int) {
	switch n {
	case 1:
	// 添加联系人
	case 2:
	// 删除联系人
	case 3:
	// 查询联系人
	case 4:
	// 编辑联系人
	default:
		fmt.Println("请输入正确的指令！")
	}
}

func main() {
	scanNum()
}
