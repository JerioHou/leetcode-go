package leetcode_go

import (
	"testing"
)

func nextPermutation(nums []int) {
	// 1 5 8 4 7 6 5 3 1
	// 从后向前找到第一个非递增的数，也就是 4
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		reverse31(nums)
		return
	}
	// 从后向前，找到第一个比4大的数, 也就是5
	j := n - 1
	for j > i && nums[j] <= nums[i] {
		j--
	}
	// 交换这两个数 1 5 8 5 7 6 4 3 1
	nums[i], nums[j] = nums[j], nums[i]
	// 反转5之后的数组
	reverse31(nums[i+1:])
}

func reverse31(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func TestNextPermutation(t *testing.T) {

	nums := []int{1, 3, 2}
	nextPermutation(nums)
	t.Log(nums)

}
