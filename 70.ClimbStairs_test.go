package leetcode_go

import "testing"

/**
 * https://leetcode-cn.com/problems/climbing-stairs/
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 */

func climbStairs(n int) int {
	x1, x2, result := 0, 0, 1
	for i := 1; i <= n; i++ {
		x2 = x1
		x1 = result
		result = x1 + x2
	}
	return result
}

func TestClimeStairs(t *testing.T) {
	t.Log(climbStairs(3))
	t.Log(climbStairs(4))
	t.Log(climbStairs(5))
}
