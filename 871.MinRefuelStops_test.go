/*
 871. 最低加油次数
https://leetcode.cn/problems/minimum-number-of-refueling-stops/
*/
package leetcode_go

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinRefuelStops(t *testing.T) {
	result := minRefuelStops2(100, 1, [][]int{{10, 1000}})
	fmt.Println(result)
	stations := [][]int{{25, 27}, {36, 187}, {140, 186}, {378, 6}, {492, 202}, {517, 89}, {579, 234}, {673, 86}, {808, 53}, {954, 49}}
	result = minRefuelStops2(1000, 83, stations)
	fmt.Println(result)
}

// 动态规划
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

// 贪心算法 + 优先级队列
func minRefuelStops2(target int, startFuel int, stations [][]int) (ans int) {
	// 剩余油量
	fuel := startFuel
	//
	prev := 0
	// 大顶堆
	hp := BigIntHeap{}
	// 加油站的个数
	n := len(stations)
	// 如果当然位置距离前一个加油站的距离 大于 剩余油量，则加油一次
	// 大顶堆存放的是 经过的加油站（未加油）的油量
	for i := 0; i <= n; i++ {
		// 下体个目标位置
		next := target
		if i < n {
			next = stations[i][0]
		}
		// 到达下一个加油站时所剩的油量
		fuel -= next - prev
		for fuel < 0 && hp.Len() > 0 {
			fuel += heap.Pop(&hp).(int)
			ans++
		}
		if fuel < 0 {
			return -1
		}
		if i < n {
			heap.Push(&hp, stations[i][1])
			prev = stations[i][0]
		}
	}
	return
}
