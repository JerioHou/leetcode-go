/*
 871. 最低加油次数
https://leetcode.cn/problems/minimum-number-of-refueling-stops/
*/
package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMinRefuelStops(t *testing.T) {
	stations := [][]int{{25, 27}, {36, 187}, {140, 186}, {378, 6}, {492, 202}, {517, 89}, {579, 234}, {673, 86}, {808, 53}, {954, 49}}
	result := minRefuelStops(1000, 83, stations)
	fmt.Println(result)
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {

	// 加油站的个数
	n := len(stations)
	// dp[i] 表示加油i次，能达到的最大距离
	dp := make([]int, n+1)
	// 加油0次，能达到的最近距离为初始油量值
	dp[0] = startFuel

	for i, v := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= v[0] {
				dp[j+1] = max871(dp[j+1], dp[j]+v[1])
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}
	return -1
}

func max871(a, b int) int {
	if a > b {
		return a
	}
	return b
}
