package main

import "fmt"

func main() {
	// 全部初始化
	var arr1 = [2][3]int{{1, 2, 3}, {5, 6, 7}}
	fmt.Println("二维数组 arr1 中的数有：", arr1)

	// 部分初始化
	var arr2 = [2][3]int{{1, 2}, {5}}
	fmt.Println("二维数组 arr2 中的数有：", arr2) // 未初始化的位置上，值均为 0

	//指定元素初始化
	var arr3 = [2][3]int{0: {2, 5, 6}}
	fmt.Println("对 arr3 第 1 行数组进行初始化的结果为：", arr3)
	var arr4 = [2][3]int{1: {2: 5}}
	fmt.Println("对 arr4 第 2 行数组中的第 3 位进行初始化的结果为：", arr4)

	// 通过初始化确定二维数组行数
	arr5 := [...][3]int{{1, 2, 3}, {5, 6}} // 只有行数可以不定长度 [...] 初始化，列数则不行
	fmt.Println("对二维数组 arr5 行数做不定长度初始化处理的结果是： ", arr5)
	fmt.Println("二维数组 arr5 的行数为：", len(arr5))
	fmt.Println("二维数组 arr5 第 1 行数组的长度(列数) 为：", len(arr5[0]))
	// 因为这是一个 int 类型的数组，len() 方法只适用于获取数组、切片、字符串或映射的长度或大小
	// fmt.Println("二维数组 arr5 第 1 行第 2 列元素的长度为：", len(arr5[0][1]))
}
