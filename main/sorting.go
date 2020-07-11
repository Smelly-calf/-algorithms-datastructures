package main

import (
	"fmt"
	"github.com/Smelly-calf/algorithms-datastructures/pkg/sorting"
)

func main() {
	// 快排
	//fmt.Println("========快速排序=======")
	//a := []int{6, 1, 9, 9, 9, 10, 4, 5, 10, 8}
	//sorting.QuickSort(a, 0, len(a)-1)
	//
	//sorting.ReverseArray(a)
	//fmt.Println("倒序：", a)

	fmt.Println("========topK=========")
	b := []int{6, 1, 9, 9, 9, 10, 4, 5, 10, 8}
	k := 4
	topK := sorting.GetTopK(b, 0, len(b)-1, k)
	fmt.Printf("top%d:%d\n", k, topK)

	// 归并排序
	//fmt.Println("========归并排序=======")
	//a = []int{2, 3, 5, 8, 3, 2}
	//fmt.Printf("排序前的数组: %+v\n", a)
	//tmp := make([]int, len(a))
	//sorting.MergeSort(a, 0, len(a)-1, tmp)
	//fmt.Printf("排序后的数组: %+v\n", a)

}
