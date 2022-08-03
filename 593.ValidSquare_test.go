package leetcode_go

import "sort"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var ans []int
	ans = append(ans, countDistance(p1, p2))
	ans = append(ans, countDistance(p1, p3))
	ans = append(ans, countDistance(p1, p4))
	ans = append(ans, countDistance(p2, p3))
	ans = append(ans, countDistance(p2, p4))
	ans = append(ans, countDistance(p3, p4))
	sort.Ints(ans)
	if ans[0] == 0 {
		return false
	}
	if ans[1] != ans[0] || ans[2] != ans[0] || ans[3] != ans[0] {
		return false
	}
	if ans[4] != ans[5] {
		return false
	}
	return ans[0]+ans[1] == ans[4]
}

func countDistance(p1 []int, p2 []int) int {

	x := p1[0] - p2[0]
	y := p1[1] - p2[1]
	return x*x + y*y
}
