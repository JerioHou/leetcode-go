/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

有效括号组合需满足：左括号必须以正确的顺序闭合。

https://leetcode-cn.com/problems/generate-parentheses/
*/
package leetcode_go

import "testing"

func generateParenthesis(n int) []string {
	// 使用回溯算法
	result := make([]string, 0)
	var curr []byte
	backTrack(&result, curr, 0, 0, n)
	return result
}

/**
result :保存结果
curr: 当前遍历路径下以保存的括号
open: 当前 左括号的数量
close: 当前 右括号的数量
max: 括号组合的总数
*/
func backTrack(result *[]string, curr []byte, open int, close int, max int) {

	if len(curr) == 2*max {
		*result = append(*result, string(curr))
		return
	}
	if open < max {
		curr = append(curr, '(')
		backTrack(result, curr, open+1, close, max)
		curr = curr[:len(curr)-1]
	}
	// 右括号的数量 不能大于左括号，否则不合法
	if close < open {
		curr = append(curr, ')')
		backTrack(result, curr, open, close+1, max)
		curr = curr[:len(curr)-1]
	}
}

func TestCode(t *testing.T) {
	t.Log(generateParenthesis(1))
	t.Log(generateParenthesis(2))
	t.Log(generateParenthesis(3))
}
