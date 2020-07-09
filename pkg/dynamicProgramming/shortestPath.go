package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n int
	)
	aa := make([][]int, n)

	fmt.Printf("矩阵:%+v\n", aa)

	//fmt.Printf("最短路径=%d", minPathSum(aa))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || grid == nil {
		return 0
	}

	n := len(grid)
	m := len(grid[0])

	var dp [][]int
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}
