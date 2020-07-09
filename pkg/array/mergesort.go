package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 8, 3, 2}
	fmt.Printf("排序前的数组: %+v\n", a)

	tmp := make([]int, len(a))
	MergeSort(a, 0, len(a)-1, tmp)
	// 打印
	fmt.Printf("排序后的数组: %+v\n", a)
}
// 归并排序: left=0，right=len(a)-1；
// 先分再治，
// 分裂的终止条件: 数组只有 1 个元素.
// 分完之后开始合
func MergeSort(a []int, left int, right int, tmp []int) {
	// 递归成立的条件
	if left < right {
		mid := (left + right) / 2
		MergeSort(a, left, mid, tmp)
		MergeSort(a, mid+1, right, tmp)
		// 分裂终止
		fmt.Printf("分: %d:%d:%d, 左=%+v, 右=%+v\n", left, mid, right, a[left:mid+1], a[mid+1:right+1])
		merge(a, left, mid, right, tmp)
	}
}

func merge(a []int, left int, mid int, right int, tmp []int) {
	// a[left:mid] 和 a[mid+1, right] 为两个排序好的数组, merge 两个数组到 tmp，再拷贝到 a.
	i := left    //左数组的指针
	j := mid + 1 //右数组的指针
	t := i       // tmp 的指针：这里每次从 i 开始，已排好的部分应该在 i 的左边

	for i <= mid && j <= right {
		if a[i] <= a[j] {
			tmp[t] = a[i]
			i++
			t++
		} else {
			tmp[t] = a[j]
			j++
			t++
		}
	}
	// a[left:mid]剩余元素
	for i <= mid {
		tmp[t] = a[i]
		i++
		t++
	}
	// a[mid+1:right] 剩余元素
	for j <= right {
		tmp[t] = a[j]
		j++
		t++
	}

	fmt.Printf("排序从 %d 到 %d, tmp=%+v, t=%d\n", left, right, tmp[:t], t)
	// 拷贝 tmp 到 a, 注意只拷贝到 t 的部分, 不包括 t
	for k, e := range tmp[:t] {
		a[k] = e
	}
}
