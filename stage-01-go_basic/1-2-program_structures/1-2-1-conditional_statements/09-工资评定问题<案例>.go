package main

import "fmt"

func main() {
	// 李四的年终工作评定，评定标准如下：
	// A级：工资涨500元；B级：工资涨200元；C级：工资不变；D级：工资降200元，E级：工资降500元
	// 设李四的原工资为5000，请用户输入李四的评级，然后显示李四来年的工资
	var level string
	var salary = 5000
	var isCorrect = true
	fmt.Println("请输入评级信息：")
	_, err := fmt.Scan(&level)
	if err != nil {
		return
	}

	switch level {
	case "A":
		salary += 500
	case "B":
		salary += 200
	case "C":
	case "D":
		salary -= 200
	case "E":
		salary -= 500
	default:
		isCorrect = false
		fmt.Println("输入的评级信息有误！！！")
	}

	if isCorrect {
		fmt.Println("工资是：", salary)
	}
}
