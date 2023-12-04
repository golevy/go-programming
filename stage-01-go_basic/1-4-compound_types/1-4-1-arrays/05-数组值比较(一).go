package main

import "fmt"

func compareValue(arr1 [5]int, arr2 [5]int) {
	var areSame = true

	for i := 0; i < len(arr1); i++ {
		// 遍历比较每一项
		if arr1[i] == arr2[i] {
			continue // 如果相同就继续
		} else {
			areSame = false
			fmt.Printf("第 %d 个元素的值不同：arr1 的值为 %d，arr2 的值为 %d\n", i+1, arr1[i], arr2[i])
		}
	}

	if areSame {
		fmt.Println("两个数组中每一个元素的值都相同！！！")
	}
}

func main() {
	// 完成两个具有相同类型和长度的数组中元素的比较，判断其相同下标对应的元素是否完全一致
	var nums1 = [5]int{1, 2, 3, 4, 5}
	var nums2 = [5]int{3, 2, 3, 6, 5}

	compareValue(nums1, nums2)
}
