package main

import "fmt"

func main() {
	// 对学员的结业考试成绩评测
	// 1. 成绩 >= 90：A
	// 2. 90 > 成绩 >= 80：B
	// 3. 80 > 成绩 >= 70：C
	// 4. 70 > 成绩 >= 60：D
	// 5. 成绩 < 60：E
	var score float64
	fmt.Println("请输入考试成绩：")
	_, err := fmt.Scan(&score)
	if err != nil {
		return
	}

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
