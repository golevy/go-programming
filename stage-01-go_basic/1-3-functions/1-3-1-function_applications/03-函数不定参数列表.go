package main

import "fmt"

func TestNum1(args ...int) {
	// args (arguments) 参数可以认为是一个集合，并在集合中存储了传递过来的数据
	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])
}

func TestNum2(args ...int) {
	// 可以用循环遍历的方式来输出打印 args 里存储的数
	for i := 0; i < len(args); i++ {
		fmt.Printf("第 %d 个数是：%d\n", i+1, args[i])
	}
}

func TestNum3(args ...int) {
	// 因为 args 本质上是一个集合，所以可以用集合专用的 for...range 方法来输出打印
	for i, arg := range args { // 如果不需要索引，可以用匿名变量 _ 来代替 i
		fmt.Printf("第 %d 个数是：%d\n", i+1, arg)
	}
}

func TextNum4(num int, args ...int) { // 注意：确定个数的参数一定要放在不定参数前面
	fmt.Println(num)
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	fmt.Println("第一种方法的测试结果是：")
	TestNum1(3, 6, 7, 8)

	fmt.Println("第二种方法的测试结果是：")
	TestNum2(3, 6, 7, 8)

	fmt.Println("第三种方法的测试结果是：")
	TestNum3(3, 6, 7, 8)

	fmt.Println("第四种方法的测试结果是：")
	TextNum4(3, 6, 7, 8)
}
