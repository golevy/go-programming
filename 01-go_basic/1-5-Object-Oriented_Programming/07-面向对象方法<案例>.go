package main

import "fmt"

// 定义一个学生
// 有姓名、性别、年龄、语文、数学、英语成绩 6 个属性
// 定义两个方法
// 一：介绍自己
// 二：计算总分与平均分

type StudentInfo struct {
	name    string
	gender  string
	age     int
	chinese float64
	math    float64
	english float64
}

func (s *StudentInfo) SayHello(userName string, userAge int, userGender string) {
	s.name = userName
	s.age = userAge
	s.gender = userGender

	if s.gender != "male" && s.gender != "female" {
		s.gender = "unknown"
	}

	if s.age < 1 || s.age > 100 {
		s.age = 18
	}

	fmt.Printf("我叫%s，年龄是%d，性别是%s\n", s.name, s.age, s.gender)
}

func (s *StudentInfo) GetScore(chinese float64, math float64, english float64) {
	s.chinese = chinese
	s.math = math
	s.english = english

	sum := s.chinese + s.math + s.english
	fmt.Printf("我叫%s，总分%.1f，平均分%.1f\n", s.name, sum, sum/3)
}

func main() {
	var stu StudentInfo
	stu.SayHello("Mike", 18, "Male")
	stu.GetScore(130, 140, 148)
}
