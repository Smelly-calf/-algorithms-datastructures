package main

import (
	"fmt"
	"time"
)

func main() {
	a1 := []int{2, 2, 1}
	start := time.Now()
	num1 := singleNumberWithExtraSpace(a1)
	end := time.Now()
	fmt.Println("singleNumWithSpace_time:", end.Sub(start))
	start = time.Now()
	num1 = singleNum(a1)
	end = time.Now()
	fmt.Println("singleNum_time:", end.Sub(start))
	fmt.Println("num1:", num1)
	fmt.Println("singleNum:num2:", num1)


	a2 := []int{4, 1, 2, 1, 2}
	start = time.Now()
	num2 := singleNumberWithExtraSpace(a2)
	end = time.Now()
	fmt.Println("singleNumWithSpace_time2:", end.Sub(start))
	start = time.Now()
	num2 = singleNum(a2)
	end = time.Now()
	fmt.Println("singleNum_time2:", end.Sub(start))
	fmt.Println("num2:", num2)
	fmt.Println("singleNum:num2:", num2)
}

/**
只出现一次的数字: 非空整数数组只有一个数组出现一次，其余均出现两次，找出出现一次的数字。要求：算法有线性时间复杂度。可以不使用额外空间吗？
同类题目：长度为 n 的数组，其中只有一个数字出现了奇数次，其他均出现偶数次，问如何使用优秀的时空复杂度快速找到这个数字。
*/

// 使用额外空间: 空间复杂度 O(n)，时间复杂度 O(n)
func singleNumberWithExtraSpace(nums []int) int {
	// 1. 使用哈希表存储数字和数字出现的次数。最后遍历hash表得到出现一次或奇数次的数字。
	// 2. 使用哈希表的keys作为集合存储数字，不在集合中就加入集合，在集合中则剔除，最后剩余的一个元素就是出现一次或奇数次的数字。
	extraMap := make(map[int]int, 10)

	for _, num := range nums {
		if value, ok := extraMap[num]; ok {
			extraMap[num] += value
			continue
		}
		extraMap[num] = 1
	}

	for k, v := range extraMap {
		if v == 1 || v%2 != 0 {
			return k
		}
		continue
	}
	return 0
}

// 不使用额外空间：空间复杂度 O(1)
func singleNum(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
