package leetcode_go

import (
	"testing"
)

/**
  5. 最长回文子串
  https://leetcode-cn.com/problems/longest-palindromic-substring/
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	// arr[i][j] = true 表示 从i到j的字串是回文字符串
	arr := CreateSlice(n, n)
	for i := range s {
		arr[i][i] = true
		i++
	}
	// l 表示截取的子串的长度
	l := 2
	begin := 0
	maxLen := 1
	for l <= n {
		for i := 0; i < n; i++ {
			// 从i开始，长度为j的字串的下表
			j := i + l - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				arr[i][j] = false
				continue
			}
			// 长度为 0 1 2
			if j-i < 3 {
				arr[i][j] = true
			} else {
				// s[i] = s[j] ,则当且仅当arr[i+1][j-1]为回文串时，arr[i][j]为回文串
				arr[i][j] = arr[i+1][j-1]
			}
			if arr[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
		l++
	}

	return s[begin : begin+maxLen]
}

func TestLongestPalindrome(t *testing.T) {

	t.Log(longestPalindrome("bb"))
}

/**
  创建一个动态的二位数组
*/
func CreateSlice(row, col int) [][]bool {
	flag := make([][]bool, row)
	for k := range flag {
		flag[k] = make([]bool, col)
	}
	return flag
}
