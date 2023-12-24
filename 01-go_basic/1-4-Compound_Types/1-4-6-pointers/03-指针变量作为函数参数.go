package main

import "fmt"

func UpdateNum(p *int) {
	*p = 60
}

func main() {
	var num = 10
	UpdateNum(&num)
	fmt.Println(num)
}
