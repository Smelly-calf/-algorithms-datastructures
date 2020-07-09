package main

import "fmt"

// 二叉树的按层遍历
// [3,9,20,null,null,15,7], 如果从1开始，节点i的左右孩子分别是:2i,2i+1；如果从0开始，节点i的左右孩子分别是:2j+1,2j+2
// 递归: 二维数组维护层数和层数组，先遍历根节点, 再遍历左子树，然后右子数。
// 递推式: 当前节点不为空, 如果二维数组长度<当前遍历的层数,追加一个切片, 把当前节点数据追加到这个切片的尾部。
func main() {
	// -1 表示节点为空
	a := []int{3, 9, 20, -1, -1, 15, 7}
	t := createBinaryTree(a)
	fmt.Printf("t:%+v\n", *t)

	// 前序遍历
	preA := make([]int, 0)
	preResult := preOrder(preA, t)
	fmt.Println("前序遍历：", preResult)

	// 按层遍历
	result := make([][]int, 0)
	r := dfsHelper(result, t, 0)
	for i := 0; i < len(r); i++ {
		fmt.Printf("第 %d 层: %+v\n", i, r[i])
	}
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

var i = -1

func createBinaryTree(a []int) *TreeNode {
	i += 1
	if i >= len(a) {
		return nil
	}
	if a[i] != -1 {
		fmt.Printf("i:%d, a[i]:%d\n", i, a[i])
		t := TreeNode{data: a[i]}
		t.left = createBinaryTree(a)
		t.right = createBinaryTree(a)
		return &t
	}
	return nil
}

func preOrder(a []int, t *TreeNode) []int {
	if t != nil {
		a = append(a, t.data)
		a = preOrder(a, t.left)
		a = preOrder(a, t.right)
	}
	return a
}

// level 从0开始
func dfsHelper(result [][]int, t *TreeNode, level int) [][]int {
	if t != nil {
		fmt.Printf("len:%d, level+1: %d\n", len(result), level+1)
		if len(result) < level+1 {
			result = append(result, make([]int, 0))
		}
		result[level] = append(result[level], t.data)
		fmt.Printf("result[%d]=%+v\n", level, result[level])
		result = dfsHelper(result, t.left, level+1)
		result = dfsHelper(result, t.right, level+1)
	}
	return result
}
