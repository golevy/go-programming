package main

import "fmt"

func main() {
	// 完成两个具有相同类型和长度数组中元素的比较，判断其相同下标对应的元素是否完全一致
	var nums1 = [5]int{1, 2, 3, 4, 5}
	var nums2 = [5]int{3, 2, 3, 4, 5}

	if nums1 == nums2 {
		fmt.Println("两个数组中的每一个元素都相同！！！")
	} else {
		fmt.Println("两个数组中的值不完全相同！！！")
	}
}
