package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20
	num1, num2 = num2, num1

	fmt.Println(num1, num2)
}
