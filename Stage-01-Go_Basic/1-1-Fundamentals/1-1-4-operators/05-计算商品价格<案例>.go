package main

import "fmt"

func main() {
	// 问题1：某商店 T 恤的价格为 35 元/件，裤子的价格为 120 元/条，小明在该店买了 3 件 T 恤和 2 条 裤子，并且打 8.8 折，请计算并显示小明应该付多少钱？
	// 问题2：如上题中打完 8.8 折后出现小数，商店为了结算方便，只收用户整数部分的钱，如本应收用户 206.3，现在只收用户 206 元，怎么处理？
	var shirt = 35
	var trousers = 120

	totalMoney := shirt*3 + trousers*2
	realMoney := float64(totalMoney) * 0.88

	fmt.Printf("支付：%.2f\n", realMoney)
	fmt.Printf("应支付：%d", int(realMoney))
}
