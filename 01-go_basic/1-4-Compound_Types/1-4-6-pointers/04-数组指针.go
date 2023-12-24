package main

import "fmt"

func UpdateArr(p *[6]int) {
	p[5] = 66
}

func main() {
	nums := [6]int{1, 2, 3, 4, 5, 6}
	// 数组指针的定义与使用
	var p *[6]int
	p = &nums
	fmt.Println("通过数组指针获取数组全部元素：", *p) // [1 2 3 4 5 6]
	fmt.Println((*p)[3])               // [] 的优先级大于 *，所以需要加括号来提高优先级
	fmt.Println(p[3])                  // 这是通过数组指针获取数组中某个元素的简洁方法
	println()

	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
	println()

	UpdateArr(p)
	fmt.Println("经修改后数组 nums 的全部元素为：", nums)
}
