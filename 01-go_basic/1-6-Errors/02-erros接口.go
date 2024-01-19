package main

import (
	"errors"
	"fmt"
)

func TestError(num1 int, num2 int) (result int, err error) {
	err = nil

	if num2 == 0 {
		err = errors.New("除数不能为0")
		return
	}

	result = num1 / num2
	return
}

func main() {
	num, err := TestError(10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
