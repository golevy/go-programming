package main

import "fmt"

func main() {
	// 考试成绩大于等于 90 输出 A，大于等于 80 输出 B，大于大于 70 输出 C，大于等于 60 输出 D，不及格输出 E
	var score float64
	fmt.Println("请输入考试成绩：")

	_, err := fmt.Scan(&score)
	if err != nil {
		return
	}

	switch { // switch 后面的表达式可以省略不写，直接放在每一个 case 里来写
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
