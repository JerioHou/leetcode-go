package leetcode_go

import (
	"fmt"
	"testing"
)

/*
  	最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：

	Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
	Paste（粘贴）：粘贴 上一次 复制的字符。
	给你一个数字 n ，你需要使用最少的操作次数，在记事本上输出 恰好 n 个 'A' 。返回能够打印出 n 个 'A' 的最少操作次数

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/2-keys-keyboard
*/
func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	k := findMaxYueShu(n)
	if k == n {
		return n
	}
	return minSteps(k) + n/k

}

// 获取n的最大公约数
func findMaxYueShu(n int) int {
	for i := n / 2; i >= 2; i-- {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func TestMin(t *testing.T) {
	fmt.Println(minSteps(12))
}
