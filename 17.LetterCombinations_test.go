/**
 * 17. 电话号码的字母组合
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
 */
package leetcode_go

import (
	"fmt"
	"testing"
)

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var result []string

func letterCombinations(digits string) []string {
	result = []string{}
	if len(digits) == 0 {
		return result
	}
	backTrack17(digits, 0, "")
	return result
}

func backTrack17(digits string, index int, combination string) {
	if index == len(digits) {
		result = append(result, combination)
	} else {
		digit := string(digits[index])
		letter := phoneMap[digit]
		for i := 0; i < len(letter); i++ {
			backTrack17(digits, index+1, combination+string(letter[i]))
		}
	}
}

func TestLetter(t *testing.T) {
	result := letterCombinations("")
	fmt.Println(result)
}
