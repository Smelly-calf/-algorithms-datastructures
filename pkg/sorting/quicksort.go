package sorting

import (
	"fmt"
)

/**
分区的时候初始基准必须为 low 或 high 的原因？以及为何要从相反的一边先开始移动？
答：
1) 必须为 low 或 high，如果以中间为基准，左右元素个数相等的时候无解. 如[1,3,2,9,8], 以数字2为基准, low停在3的地方，high也停在3的地方; 3与2交换，p 变为3
2) 从哪边开始移动决定了相遇位置的值是大于 p 还是小于 p; 以 low 为基准，我们要找比 p 小的元素与 p 交换，这样才能保证数组左边元素小于 p；假如我们反其道而行从 low 开始移动(在比low大的数停下), 相遇位置的值比 p 大，是不对的.
*/

func ReverseArray(a []int) {
	n := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		tmp := a[i]
		a[i] = a[n-i]
		a[n-i] = tmp
	}
}

func QuickSort(a []int, low int, high int) {
	// 数组元素>1 排序，否则递归结束
	if low < high {
		p := partition(a, low, high)
		fmt.Printf("分割:[%d, %d], 排序后: %+v, p=%d\n", low, high, a, p)
		// 左边递归结束条件: low < p-1
		if low < p-1 {
			QuickSort(a, low, p-1)
		}
		// 右边递归结束条件：p+1 < high
		if p+1 < high {
			QuickSort(a, p+1, high)
		}
	}
}

func partition(a []int, low int, high int) int {
	p := high // 以高位为基准，相遇位置的数字一定是比 p 大的数字才能与 p 交换，因此一定要从 low 先开始找比 a[p] 大的数字停下。
	for low < high {
		// low 移动, 在比 a[p] 大的位置停下
		for a[low] <= a[p] && low < high {
			low += 1
		}
		// high 移动, 在比 a[p] 小的位置停下
		for a[high] >= a[p] && low < high {
			high -= 1
		}
		// 此时 a[low]>a[p]>a[high], 交换低位和高位的数字
		if low < high {
			tmp := a[low]
			a[low] = a[high]
			a[high] = tmp
		}
	}
	// 判断如果 p==low，不交换; 否则相遇的位置一定比在 p 的左边，相遇位置即为下一次的分区下标。
	if low < p {
		tmp := a[p]
		a[p] = a[low]
		a[low] = tmp
	}
	return low
}

// 快排获取 topK: 递归每次快排找到一个pivot，比较 pivot 和 sub=n-k(排好序后 topK 对应的数组下标)：
//	如果pivot=(n-k), 那么 pivot 就是 topK, 递归结束；
//	如果pivot<(n-k), 那么 topK 在 pivot 右边的数组。
//	如果 pivot>(n-k), 那么 topK 在 pivot 左边的数组。
// 以上需在递归中进行 .

// 假设存在 topK
func GetTopK(nums []int, low int, high int, k int) int {
	if low > high {
		panic("非法输入")
	}
	p := partition(nums, low, high)
	fmt.Printf("分割:[%d, %d], %+v, p=%d\n", len(nums)-k, low, high, nums)
	sub := len(nums) - k
	result := 0
	if p == sub {
		result = nums[p]
	} else if p < sub {
		result = GetTopK(nums, p+1, high, k)
	} else {
		result = GetTopK(nums, low, p-1, k)
	}

	return result
}
