/**

https://leetcode.cn/problems/valid-parentheses/
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

*/
package leetcode_go

import "testing"

func TestIsValid(t *testing.T) {
	s := "{}"
	t.Log(isValid(s))
}

func isValid(s string) bool {

	n := len(s)
	if n%2 != 0 {
		return false
	}

	pairMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}

	for i := 0; i < n; i++ {
		b := pairMap[s[i]]
		if b <= 0 {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		pop := stack[len(stack)-1]
		if b != pop {
			return false
		}
		stack = stack[0 : len(stack)-1]
	}
	return len(stack) == 0
}
