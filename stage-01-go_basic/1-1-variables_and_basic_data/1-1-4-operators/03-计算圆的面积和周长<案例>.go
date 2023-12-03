package main

import "fmt"

func main() {
	// 输入半径，计算圆的面积和周长并打印出来（PI 为 3.14），周长：2*PI*r，面积：PI*r*r
	// 1. 确定 π 的取值
	const PI = 3.14
	// 2. 接收用户输入的半径
	fmt.Println("请输入半径：")
	var r float64
	_, err := fmt.Scan(&r)
	if err != nil {
		return
	}
	// 3. 计算圆的面积
	area := PI * r * r
	// 4. 计算圆的周长
	p := 2 * PI * r
	// 5. 打印输出结果
	fmt.Printf("面积是%.2f\n", area)
	fmt.Printf("周长是%.2f", p)
}
