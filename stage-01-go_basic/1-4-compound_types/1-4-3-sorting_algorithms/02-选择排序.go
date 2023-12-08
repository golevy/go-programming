package main

import "fmt"

func main() {
	slice := []int{5, 9, 0, 2, 7}

	for i := 0; i < len(slice)-1; i++ {
		minNumber := slice[i]
		minIndex := i

		for j := i + 1; j < len(slice); j++ {
			if minNumber > slice[j] {
				minNumber = slice[j]
				minIndex = j
			}
		}

		if minIndex != i {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
		}
	}

	fmt.Println("选择排序的结果为：", slice)
}
