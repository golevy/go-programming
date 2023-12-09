package main

import "fmt"

func main() {
	// 找出 100-999 之间的水仙花数
	// 水仙花数：是指一个三位数，它的每位数字的立方和等于其本身
	// 例如：153 = 3*3*3 + 5*5*5 + 1*1*1
	var h int     // 百位
	var t int     // 十位
	var u int     // 个位
	var count int // 计数器

	for i := 100; i <= 999; i++ {
		h = i / 100      // 跟 100 相除，有几个 100，百位就是几
		t = i % 100 / 10 // 跟 100 取余之后就只剩下十位数，除以 10，有几个 10，十位就是多少
		u = i % 10       // 跟 10 取余，10 取没了，剩下的就是个位的数字

		if i == h*h*h+t*t*t+u*u*u {
			count++
			fmt.Printf("第 %d 个水仙花数：%d\n", count, i)
		}
	}
	fmt.Printf("一共有 %d 个水仙花数", count)
}
