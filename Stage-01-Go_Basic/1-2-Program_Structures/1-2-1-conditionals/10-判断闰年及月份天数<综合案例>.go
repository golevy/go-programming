package main

import "fmt"

func main() {
	// 请用户输入年份，再输入月份，并输出该月的天数（需要考虑闰年）
	// 闰年的判定（符合下述条件之一）
	// 年份能被 400 整除（如：2000）
	// 年份能被 4 整除但不能被 100 整除（2008）

	// 1. 定义变量存储年份月份和天数
	var year int
	var month int
	var day int
	// 2. 要求用户输入年份和月份
	fmt.Println("请输入年份：")
	_, err := fmt.Scan(&year)
	if err != nil {
		return
	}
	fmt.Println("请输入月份：")
	_, err = fmt.Scan(&month)
	if err != nil {
		return
	}
	// 3. 判断输入的月份是否正确
	if month >= 1 && month <= 12 {
		// 4. 如果月份是 1，3，5，7，8，10，12 输输出的天数是 31 天
		switch month {
		case 1:
			fallthrough // 这里用到了 case 穿透特性
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			day = 31
		case 2:
			// 5. 如果月份是 2，判断闰年，闰年 29 天，默认 28 天
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				day = 29
			} else {
				day = 28
			}
		default:
			// 6. 其他月份输出的是 30 天
			day = 30
		}
		fmt.Println("该月的天数是：", day)
	} else {
		fmt.Println("输入的月份有误！！！")
	}
}
