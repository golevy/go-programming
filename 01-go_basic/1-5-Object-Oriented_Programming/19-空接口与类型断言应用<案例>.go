package main

import "fmt"

type Outcomer interface {
	GetResult() int
	SetData(data ...interface{}) bool // 校验参与运算的数据的个数以及类型，"..."表示不定参数
}

type NumberObject struct {
	numA int
	numB int
}

type AddNumObject struct {
	NumberObject
}

type SubNumObject struct {
	NumberObject
}

type OperatorNumFactory struct {
	id string
}

func (a *AddNumObject) GetResult() int {
	return a.numA + a.numB
}

func (s *SubNumObject) GetResult() int {
	return s.numA - s.numB
}

func (a *AddNumObject) SetData(data ...interface{}) bool {
	var b = true

	if len(data) > 2 {
		fmt.Println("参数个数错误！！")
		b = false
	}

	value1, ok1 := data[0].(int)

	if !ok1 {
		fmt.Println("第一个数类型错误")
		b = false
	}

	value2, ok2 := data[1].(int)

	if !ok2 {
		fmt.Println("第二个数类型错误")
		b = false
	}

	a.numA = value1
	a.numB = value2

	return b
}

func (s *SubNumObject) SetData(data ...interface{}) bool {
	var b = true

	if len(data) > 2 {
		fmt.Println("参数个数错误！！")
		b = false
	}

	value1, ok1 := data[0].(int)

	if !ok1 {
		fmt.Println("第一个数类型错误")
		b = false
	}

	value2, ok2 := data[1].(int)

	if !ok2 {
		fmt.Println("第二个数类型错误")
		b = false
	}

	s.numA = value1
	s.numB = value2

	return b
}

func (o *OperatorNumFactory) CreateOperator(op string) Outcomer {
	switch op {
	case "+":
		return new(AddNumObject)
	case "-":
		return new(SubNumObject)
	default:
		return nil
	}
}

func Operator(h Outcomer) int {
	return h.GetResult()
}

func main() {
	var operator OperatorNumFactory

	obj1 := operator.CreateOperator("+")
	bool1 := obj1.SetData(30, 10)
	if bool1 {
		num1 := Operator(obj1)
		fmt.Println("加法计算结果为", num1)
	}

	obj2 := operator.CreateOperator("-")
	bool2 := obj2.SetData(30, 10)
	if bool2 {
		num2 := Operator(obj2)
		fmt.Println("减法计算结果为", num2)
	}
}
