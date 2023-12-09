package main

import "fmt"

func PrintMap(m map[int]string) {
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func DeleteMap(m map[int]string) {
	delete(m, 3)
}

func main() {
	var m = map[int]string{1: "Golang", 2: "JavaScript", 3: "Python"}

	PrintMap(m)
	DeleteMap(m) // 在函数中删除 map 的值后会影响原 map，类似于 slice，因为内存地址都是一样的，删除的都是同一内存指向的数据
	fmt.Println("删除后的 map 为：", m)
}
