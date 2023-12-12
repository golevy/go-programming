package main

import "fmt"

func main() {
	var a = 10

	// 一级指针
	var p *int
	p = &a // 一级指针指向 a 的内存地址

	// 二级指针（存储的是一级指针的内存地址）
	// *pp：这是对 pp 的一级解引用。*pp 给出的是 pp 指向的值，即 p 的值，由于 p 本身是一个指针，所以 *pp 相当于 p。
	// **pp：这是对 pp 的二级解引用，首先，*pp 给出 p 的值，然后再次解引用 (*p) 给出 p 指向的值，即 a 的值
	var pp **int
	pp = &p
	**pp = 200

	fmt.Println(a) // a
}
