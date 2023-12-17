package main

import (
	"fmt"
)

type Person struct {
	userName     string
	addressPhone map[string]string // key 表示电话的类型（公司/个人等），value：电话
}

var personList = make([]Person, 0) // 定义通讯录切片

// 界面操作提示
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

// 执行对应操作
func switchType(n int) {
	switch n {
	case 1:
		// 添加联系人
		addPerson()
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

func addPerson() {
	var name string
	var address string
	var phone string
	var addressPhone = make(map[string]string) // 保存电话的类型和电话，电话类型作为 key
	var exit string                            // 退出标志

	// 添加姓名
	fmt.Println("请输入姓名：")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// 保存电话类型
		fmt.Println("请输入电话类型：")
		_, err1 := fmt.Scan(&address)
		if err1 != nil {
			fmt.Println(err1)
			return
		}

		// 保存电话号码
		fmt.Println("请输入电话号码：")
		_, err2 := fmt.Scan(&phone)
		if err2 != nil {
			fmt.Println(err2)
			return
		}

		// 将电话号码及电话类型存储到 addressPhone 中
		addressPhone[address] = phone

		// 退出界面
		fmt.Println("如果结束电话的录入，请按 Q")
		_, err3 := fmt.Scan(&exit)
		if err3 != nil {
			fmt.Println(err3)
			return
		}
		if exit == "Q" {
			break
		} else {
			continue
		}
	}

	// 将联系人的信息存储到切片中
	personList = append(personList, Person{userName: name, addressPhone: addressPhone})
}

func main() {
	scanNum()
}
