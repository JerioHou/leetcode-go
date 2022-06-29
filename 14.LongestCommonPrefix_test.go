/**
14. 最长公共前缀
https://leetcode.cn/problems/longest-common-prefix/
*/
package leetcode_go

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	strs := []string{"b", "b", "blight"}
	t.Log("result = ", longestCommonPrefix2(strs))

}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var rows, cols = len(strs), len(strs[0])
	for i := 0; i < cols; i++ {
		var firstChar = strs[0][i]
		for j := 1; j < rows; j++ {
			if i == len(strs[j]) || firstChar != strs[j][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix(strs []string) string {

	shortest := len(strs[0])
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		if len(str) < shortest {
			shortest = len(str)
		}
	}
	if shortest == 0 {
		return ""
	}
	index := getIndex(strs, shortest)
	if index == 0 {
		return ""
	}
	return strs[0][:index]
}

func getIndex(strs []string, shortest int) int {
	cursor := 0
	for i := 0; i < shortest; i++ {
		var tmp byte
		for j := 0; j < len(strs); j++ {
			str := strs[j]
			if tmp == 0 {
				tmp = str[i]
			} else {
				if tmp != str[i] {
					return cursor
				}
			}
		}
		cursor++
	}
	return cursor
}
