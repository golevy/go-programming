package main

import "fmt"

// range 关键字用于遍历数组、切片、字符串、map 或通道（channel）
// 当用于遍历数组或切片时，它返回两个值：当前索引和该索引处的值
// 当用于遍历 map 时，它返回两个值：当前键和该键对应的值

func main() {
	arr := [3]int{1, 2, 3}

	for index, value := range arr { // 当使用 range 遍历数组时，它返回每个元素的索引（index）和该索引处的值的副本（value）
		fmt.Println("value：", value)
		value += 1     // 所以这里修改的是 value 变量的副本，这个修改不会影响原始数组 arr 中的元素（这里是 int 没有影响，但如果这个 value 是一个指针变量，就会有影响）
		arr[index] = 3 // 修改了原数组 arr 中的元素，这种修改会影响到数组 arr 本身，但不会影响 range 遍历过程中的 value 值
	}

	fmt.Println("遍历修改后的数组 arr 结果为：", arr) // [3 3 3]
}
