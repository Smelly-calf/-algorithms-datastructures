package main

import "fmt"

// 100个无差别的球，10个编号的篮子，每个篮子至少有一个，多少种方法？
// 100个球一字排开，有99个间隙，在99个间隙中放入9根筷子将100个球分成10份，有C（99，9）=99!/(90!9!)
// 递归公式：第 i 个篮子放入方法数 = 前面 i-1 个篮子放入个数*(100-前 i-1 )
func main() {
	num := putBall(100, 1)
	fmt.Println(num)
}

// 返回剩余个数，每次计算的结果
func putBall(n int, num int) int {
	a := 1
	if n < 10 {
		n = n - num
		res := putBall(n, num+1)
		a *= res
	}
	return a
}

//func levelOrder1(root *TreeNode) [][]int {
//	result := make([][]int, 0)
//	if root == nil {
//		return result
//	}
//	return dfsHelper(root, 0)
//}
//
//func dfsHelper(node *TreeNode, level int) [][]int {
//	result := make([][]int, 0)
//
//	if node == nil {
//		return result
//	}
//	if len(result) < level+1 {
//		result = append(result, make([]int, 0))
//	}
//	result[level] = append(result[level], node.Val)
//	dfsHelper(node.Left, level+1)
//	dfsHelper(node.Right, level+1)
//	return result
//}
