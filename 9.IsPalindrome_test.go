package leetcode_go

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	return rev == x || x == rev/10
}
