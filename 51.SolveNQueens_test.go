package leetcode_go

import (
	"testing"
)

/**
  n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
*/

func TestNqueens(t *testing.T) {
	result := solveNQueens(4)
	t.Log(count)
	t.Log(result)
}

var count = 0

func solveNQueens(n int) [][]string {
	results := [][]string{}
	result := make([]int, n)
	calNQueen(&results, result, n, 0)
	return results
}

func calNQueen(finalResult *[][]string, result []int, n int, row int) {
	if row == n {
		convert(finalResult, result)
		count++
		return
	}
	for col := 0; col < n; col++ {
		if isOk(result, row, col) {
			result[row] = col
			calNQueen(finalResult, result, n, row+1)
		}
	}
}

// 判断row行column列是否符合规则
func isOk(result []int, row int, column int) bool {
	// 左上位置的列
	leftup := column - 1
	// 右上的列
	rightup := column + 1
	// 因为是从行的方向上，从上往下放置，所以校验的时候，行下方是空的，不用 校验
	// 所以从下往上校验
	for i := row - 1; i >= 0; i-- {
		// 意思是：row行上方的一行，Q也是在第column列，冲突了
		if result[i] == column {
			return false
		}
		if leftup >= 0 {
			// 意思是：row行上方的一行，Q也是在第column-1列，斜对角，冲突了
			if result[i] == leftup {
				return false
			}
		}
		if rightup < len(result) {
			// 意思是：row行上方的一行，Q也是在第column+1列，斜对角，冲突了
			if result[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

func convert(finalResult *[][]string, from []int) {
	n := len(from)
	board := []string{}
	for row := 0; row < len(from); row++ {
		arr := make([]byte, n)
		for column := 0; column < len(from); column++ {
			if from[row] == column {
				arr[column] = 'Q'
			} else {
				arr[column] = '.'
			}
		}
		board = append(board, string(arr))
	}
	*finalResult = append(*finalResult, board)
}
