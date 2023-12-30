package main

import "fmt"

// 根据以下信息，实现对应的继承关系
// 演员：我叫Mike，我的爱好是Music，我的年龄是34，我是一个男演员
// 程序员：我叫Aubrey，我的年龄是23，我是男生，我的工作年限是3年
// 怎样抽取类、成员，考虑是否有父类，从而实现成员、方法继承

type PersonInfo struct {
	name   string
	age    int
	gender string
}

type Act struct {
	PersonInfo
	hobby string
}

type Pro struct {
	PersonInfo
	workYear int
}

func (p *PersonInfo) setValue(userName string, userAge int, userGender string) {
	p.name = userName
	p.age = userAge
	p.gender = userGender
}

func (a *Act) printInfo(hobby string) {
	a.hobby = hobby
	fmt.Printf("我叫%s，我的爱好是%s，我的年龄是%d，我是一个%s演员\n", a.name, a.hobby, a.age, a.gender)
}

func (p *Pro) printInfo(workYear int) {
	p.workYear = workYear
	fmt.Printf("我叫%s，我的年龄是%d，我是%s生，我的工作年限是%d年", p.name, p.age, p.gender, p.workYear)
}

func main() {
	var act Act
	act.setValue("Mike", 34, "男")
	act.printInfo("Music")

	var pro Pro
	pro.setValue("Aubrey", 23, "男")
	pro.printInfo(3)
}
