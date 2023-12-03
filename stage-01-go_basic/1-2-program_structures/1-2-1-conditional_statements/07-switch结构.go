package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入表示星期几的数字：")

	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	switch num {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("输入表示星期几的数字有误！！！")
	}
}
