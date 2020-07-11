package sorting

import "fmt"

// 摆动序列: 将数组排列成 nums[1]<nums[2]>nums[3]<nums[4]...
func Wobble(nums []int) {
	topK := GetTopK(nums, 0, len(nums)-1, 2)
	fmt.Println(topK)
}


