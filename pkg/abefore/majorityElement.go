package main

import (
	"fmt"
	"time"
)

func main() {
	a1 := []int{3, 2, 3}
	start := time.Now()
	num := majorityElementByQuickSort(a1)
	end := time.Now()
	fmt.Println("majorityElement1:", num, "elapsed:", end.Sub(start))
	start = time.Now()
	num = majorityElement(a1)
	end = time.Now()
	fmt.Println("majorityElement2:", num, "elapsed:", end.Sub(start))

	a2 := []int{2, 2, 1, 1, 1, 2, 2}
	start = time.Now()
	num = majorityElementByQuickSort(a2)
	end = time.Now()
	fmt.Println("majorityElement1:", num, "elapsed:", end.Sub(start))
	start = time.Now()
	num = majorityElement(a2)
	end = time.Now()
	fmt.Println("majorityElement2:", num, "elapsed:", end.Sub(start))

}

// 多数元素，出现次数> n/2 的元素：方案二：摩尔投票法
func majorityElement(nums []int) int {
	cnt := 1
	maj := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == maj {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				i++ // 从i+1开始遍历
				maj = nums[i]
				cnt = 1
			}
		}
	}
	return maj
}

// 多数元素，出现次数> n/2 的元素: 方案一：快排
func majorityElementByQuickSort(nums []int) int {
	fmt.Println("排序前：", nums)
	quickSort(nums)
	fmt.Println("排序后：", nums)

	return nums[len(nums)/2]
}

func quickSort(nums []int) {
	n := len(nums)
	if n > 1 {
		// 递推关系: 排完后左边元素全部小于基准元素，右边元素全部大于基准元素
		p := partition(nums, 0, n-1)
		if 0 < p-1 {
			// 递归排左边
			quickSort(nums[0:p])
		}
		if p+1 < n-1 {
			// 递归排右边
			quickSort(nums[p+1 : n])
		}
	}
}

func partition(nums []int, low int, high int) int {
	p := low //基准
	for low < high {
		for low < high && nums[high] >= nums[p] {
			high -= 1
		}
		for low < high && nums[low] <= nums[p] {
			low += 1
		}
		// 注意每次交换元素
		if low < high {
			tmp := nums[low]
			nums[low] = nums[high]
			nums[high] = tmp
		}
	}
	if low != p {
		tmp := nums[low]
		nums[low] = nums[p]
		nums[p] = tmp
	}
	return low
}
