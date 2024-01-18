package main

import "fmt"

type Resulter interface {
	GetResult() int
}

type NumObject struct {
	numA int
	numB int
}

type AddObject struct {
	NumObject
}

type SubObject struct {
	NumObject
}

func (a *AddObject) GetResult() int {
	return a.numA + a.numB
}

func (s *SubObject) GetResult() int {
	return s.numA - s.numB
}

type OperatorFactory struct {
	id string
}

func (o *OperatorFactory) CreateOperator(op string, numA int, numB int) int {
	switch op {
	case "+":
		add := AddObject{NumObject{numA: numA, numB: numB}}
		return OperatorWho(&add)
	case "-":
		sub := SubObject{NumObject{numA: numA, numB: numB}}
		return OperatorWho(&sub)
	default:
		return 0
	}
}

func OperatorWho(h Resulter) int {
	return h.GetResult()
}

func main() {
	var operator OperatorFactory
	num1 := operator.CreateOperator("+", 20, 15)
	fmt.Println("加法计算结果为", num1)
	num2 := operator.CreateOperator("-", 20, 15)
	fmt.Println("减法计算结果为", num2)
}
