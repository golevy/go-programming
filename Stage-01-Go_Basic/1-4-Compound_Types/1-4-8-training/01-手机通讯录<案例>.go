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
		removePerson()
	case 3:
	// 查询联系人
	case 4:
	// 编辑联系人
	default:
		fmt.Println("请输入正确的指令！")
	}
}

func removePerson() {
	var name string
	var index = 1 // 记录要删除的联系人信息在切片中的下标
	fmt.Println("请输入要删除的联系人姓名：")
	_, err4 := fmt.Scan(&name)
	if err4 != nil {
		fmt.Println(err4)
		return
	}

	// 判断切片中是否存储了要删除的联系人信息
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			index = 1
			break
		}
	}

	if index != -1 {
		// 切片的截取操作，相当于将用户输入的 name 给排除了
		personList = append(personList[:index], personList[index+1:]...) // append() 函数第二个参数如果是切片，后面需要加...
		// 由于 append 的第二个参数需要单个元素，而不是切片，所以当你需要将一个切片的元素追加到另一个切片时，你需要使用 ... 来展开切片中的元素
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

	// 展示切片中存储的联系人信息
	showPersonList()
	println()
}

func showPersonList() {
	if len(personList) == 0 {
		fmt.Println("暂时没有联系人信息")
	} else {
		for _, value := range personList {
			fmt.Println("姓名：", value.userName)
			for k, v := range value.addressPhone {
				fmt.Println("电话类型：", k)
				fmt.Println("电话号码：", v)
			}
		}
	}
}

func main() {
	for {
		scanNum()
	}
}
