/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/
package leetcode_go

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	right, result := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(charMap, s[i-1])
		}
		for right < n && charMap[s[right]] == 0 {
			charMap[s[right]] = 1
			right++
		}
		result = max(result, right-i)
	}
	return result
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("111111"))
}
