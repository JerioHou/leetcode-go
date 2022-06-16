/**

https://leetcode-cn.com/problems/pascals-triangle/

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/
package leetcode_go

import "testing"

func pascaksTriangle(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func TestPascaksTriangle(t *testing.T) {
	t.Log(pascaksTriangle(3))
}
