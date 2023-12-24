package main

import "fmt"

type StudentInfo struct {
	id   int
	name string
	age  int
}

func InitData(stu []StudentInfo) {
	for i := 0; i < len(stu); i++ {
		fmt.Printf("请输入第 %d 个学生的详细信息\n", i+1)
		_, err := fmt.Scan(&stu[i].id, &stu[i].name, &stu[i].age)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func GetMax(stu []StudentInfo) {
	var maxAge = stu[0].age
	var maxIndex int

	for i := 0; i < len(stu); i++ {
		if stu[i].age > maxAge {
			maxAge = stu[i].age
			maxIndex = i
		}
	}

	fmt.Println("最大年龄学生的信息为：", stu[maxIndex])
}

func main() {
	// 输出年龄最大的学生的详细信息
	stu := make([]StudentInfo, 3)

	InitData(stu)

	GetMax(stu)
}
