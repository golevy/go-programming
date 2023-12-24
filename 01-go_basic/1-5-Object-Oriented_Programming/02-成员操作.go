package main

import "fmt"

type People struct {
	id   int
	name string
	age  int
}

type Pupil struct {
	People
	score float64
}

func main() {
	var pupil1 = Pupil{People{id: 001, name: "Mike", age: 9}, 88}
	var pupil2 = Pupil{People{id: 002, name: "Lily", age: 10}, 90}

	/// 获取成员值
	fmt.Println(pupil1.id)
	fmt.Println(pupil2.score)

	// 修改成员值
	pupil1.age = 11
	fmt.Println("pupil1 的年龄为：", pupil1.age)
	fmt.Println("pupil2 的年龄为：", pupil2.age) // 修改 pupil1 中父类 People 的成员值，对 pupil2 没有影响
}
