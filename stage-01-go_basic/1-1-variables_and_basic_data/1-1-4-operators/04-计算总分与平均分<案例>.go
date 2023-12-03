package main

import "fmt"

func main() {
	// 某学生三门课成绩为，语文：90，数学：89，英语：69，求总分与平均分
	var chinese = 90
	var math = 89
	var english = 95

	// 计算总分
	score := chinese + math + english

	// 平均分
	avg1 := float64(score) / 3 // 注意这里一定要先将总分转为 float64 后再进行平均分的计算，这里整数 3 被隐式地转换为浮点数来进行除法运算
	// 进行除法运算时，如果其中一个操作数是浮点数，另一个是整数，那么整数会先被隐式地转换为浮点数，然后执行浮点数除法。

	avg2 := float64(score / 3) // (score / 3) 是整型数据的运算，其结果一定是整数 91，再进行 float64 之后，就会将原来的整数变为 91.000000

	fmt.Printf("总分为：%d\n", score)
	fmt.Printf("正确平均分为：%.2f\n", avg1)
	fmt.Printf("错误平均分为：%.2f", avg2)
}
