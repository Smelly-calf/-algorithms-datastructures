package main

import (
	"fmt"
)

func main() {
	a := []int{6, 1, 9, 9, 9, 10, 4, 5, 10, 8}
	QuickSort(a)
	fmt.Println(a)
}
// 基准必须为 low 或 high 么？以及为何要从距离 p 大的一边开始移动？
// 答：1) 必须为 low 或 high，如果以中间为基准，左右相等的时候无解.
// 2) 从那边开始移动决定了相遇位置的值是大于 p 还是小于 p; 假如 p 是 low, 从 low 开始移动, 相遇位置的值应该比 p 大, 那么相遇的位置就不是基准位置了.
func partition(a []int) int {
	low := 0
	high := len(a) - 1

	p := high
	fmt.Printf("以 %d 为基准开始移动\n", p)
	// 第一次排序结束的标志：low=high
	for low != high {
		// low 移动, 停在比 a[p] 大的点
		for a[low] <= a[p] && low < high {
			low += 1
		}
		// high 移动, 停在比 a[p] 小的点
		for a[high] >= a[p] && low < high {
			high -= 1
		}

		// 移动之后有可能出现 low=high 的情况，需要判断条件
		if low < high {
			fmt.Printf("交换 %d=%d 和 %d=%d, ", low, a[low], high, a[high])
			//需要交换 low 和 high 的值
			tmp := a[low]
			a[low] = a[high]
			a[high] = tmp
			fmt.Printf("交换后 a=%+v\n", a)
		}
	}

	// 判断如果 p==low，不交换; low=high 的位置为下一次分区的位置, p 的值和 low 的值交换
	if p != low {
		fmt.Printf("相遇: %d=%d 和 %d=%d 交换", p, a[p], low, a[low])
		tmp := a[p]
		a[p] = a[low]
		a[low] = tmp
		p = low
	}
	fmt.Printf("排序后, a=%+v, 新的分区点=%d \n\n", a, p)
	return p
}

// 递归
func QuickSort(a []int) {
	fmt.Printf("排序前: %+v\n", a)

	low := 0
	high := len(a) - 1
	// low==high 或者 high==1 表示数组只有1个元素, 无须排序
	if low != high {
		p := partition(a)

		// 递归结束条件: 数组只有1个元素，注意左闭右开
		if low < p-1 {
			fmt.Printf("左分区: %d %d ", low, p-1)
			QuickSort(a[low:p]) // 先递归排 p 的左边, 直到 low=p-1.
		}

		// 注意左闭右开, 递归结束条件: 数组只有1个元素
		if p < high-1 {
			fmt.Printf("右分区: %d %d ", p+1, high)
			QuickSort(a[p+1 : high+1]) // 再递归排 p 的右边, 直到 p=high-1
		}
	}
}
