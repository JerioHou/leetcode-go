/*
 899. 有序队列
 https://leetcode.cn/problems/orderly-queue/
*/
package leetcode_go

import "sort"

func orderlyQueue(s string, k int) string {
	if k == 1 {
		ans := s
		for i := 1; i < len(s); i++ {
			s = s[1:] + s[:1]
			if s < ans {
				ans = s
			}
		}
		return ans
	}
	t := []byte(s)
	// 当K>2时，经过若干次移动，总是可以互换任意两个字符的位置
	// 所以最终总是可以将字符串变为有序的
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}
