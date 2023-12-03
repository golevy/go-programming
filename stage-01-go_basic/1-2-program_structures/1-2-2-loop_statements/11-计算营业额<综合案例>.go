package main

import "fmt"

func main() {
	// 2006 年营业额为 80000 元，每年增长 25%，请问按此增长速度，到哪一年营业额将达到 20 万元
	year := 2006
	var money float64

	for money = 80000; money <= 200000; money = money * 1.25 {
		year += 1
	}

	fmt.Printf("营业额达到 20 万元的年份是 %d 年", year)
}
