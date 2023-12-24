package main

import (
	"fmt"
)

type Contact struct {
	userName     string
	addressPhone map[string]string // key 表示电话的类型（公司/个人等），value：电话
}

var contactList = make([]Contact, 0) // 定义通讯录切片

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
		addContact()
	case 2:
		// 删除联系人
		removeContact()
	case 3:
		// 查询联系人
		findContact()
	case 4:
		// 编辑联系人
		editContact()
	default:
		fmt.Println("请输入正确的指令！")
	}
}

func addContact() {
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
	contactList = append(contactList, Contact{userName: name, addressPhone: addressPhone})

	// 展示切片中存储的联系人信息
	showContactList()
}

func removeContact() {
	var name string
	var index = -1 // 记录要删除的联系人信息在切片中的下标
	fmt.Println("请输入要删除的联系人姓名：")
	_, err4 := fmt.Scan(&name)
	if err4 != nil {
		fmt.Println(err4)
		return
	}

	// 判断切片中是否存储了要删除的联系人信息
	for i := 0; i < len(contactList); i++ {
		if contactList[i].userName == name {
			index = 1
			break
		}
	}

	if index != -1 {
		// 切片的截取操作，相当于将用户输入的 name 给排除了
		contactList = append(contactList[:index], contactList[index+1:]...) // append() 函数第二个参数如果是切片，后面需要加...
		// 由于 append 的第二个参数需要单个元素，而不是切片，所以当你需要将一个切片的元素追加到另一个切片时，你需要使用 ... 来展开切片中的元素
	}

	showContactList()
}

func findContact() *Contact {
	// 1. 输入要查询的联系人姓名
	var name string
	var index = -1 // 记录找到的联系人信息在切片中的下标
	fmt.Println("请输入要查询的联系人姓名：")
	_, err5 := fmt.Scan(&name)
	if err5 != nil {
		fmt.Println(err5)
		return nil
	}

	// 2. 根据输入的联系人姓名，查找对应的联系信息
	for key, value := range contactList {
		if value.userName == name {
			index = key
			fmt.Println("联系人姓名：", value.userName)

			for k, v := range value.addressPhone {
				fmt.Printf("电话类型：%s\n", k)
				fmt.Printf("电话号码：%s\n", v)
			}
		}
	}

	// 3. 打印输出结果
	if index == -1 {
		fmt.Println("没有找到该联系人信息！！！")
		return nil
	} else {
		return &contactList[index]
	}
}

func editContact() {
	// 1. 查找到要编辑的联系人信息
	var name string              // 存储新的联系人姓名
	var num int                  // 存储需要修改的数据的数字编号
	var menu = make([]string, 0) // 保存电话类型，方便后面修改
	var pNum int                 // 编辑的电话类型编号
	var phone string             // 存储新的电话号码
	var p *Contact
	p = findContact()

	// 2. 编辑联系人信息
	if p != nil {
		for {
			fmt.Println("编辑联系人姓名请按：5，编辑电话请按：6，退出请按：7")
			_, err := fmt.Scan(&num)
			if err != nil {
				fmt.Println(err)
				return
			}

			switch num {
			case 5:
				// 修改联系人姓名
				fmt.Println("请输入新的姓名：")
				_, err = fmt.Scan(&name)
				if err != nil {
					fmt.Println(err)
					return
				}
				p.userName = name
				showContactList()
			case 6:
				// 编辑联系人电话
				var j int
				for key, value := range p.addressPhone {
					fmt.Println("编辑（", key, "）", value, "请按：", j)
					menu = append(menu, key)
					j++
				}

				fmt.Println("请输入编辑号码的类型：")
				_, err = fmt.Scan(&pNum)
				if err != nil {
					fmt.Println(err)
					return
				}

				for index, value := range menu {
					if index == pNum {
						fmt.Println("请输入新的电话号码：")
						_, err = fmt.Scan(&phone)
						if err != nil {
							fmt.Println(err)
							return
						}
						p.addressPhone[value] = phone
					}
				}

			}

			if num == 7 {
				break
			}
		}
	} else {
		fmt.Println("没有找到要编辑的联系人信息！！！")
	}
}

func showContactList() {
	if len(contactList) == 0 {
		fmt.Println("暂时没有联系人信息")
	} else {
		for _, value := range contactList {
			fmt.Println("通讯录信息如下：")
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
