package main

import "fmt"

func main() {
	var arr [2][3]int
	arr[0][1] = 123
	arr[1][2] = 456

	fmt.Println("二维数组中第 1 行第 2 列的数是", arr[0][1])
	fmt.Println("二维数组中第 2 行第 3 列的数是", arr[1][2])
}
