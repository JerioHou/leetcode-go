/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
*/

package leetcode_go

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	if len(nums) < 3 {
		return nil
	}
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i <= n-3; i++ {
		left := i + 1
		right := n - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for right > left {
			if left != i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < n-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[i]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				right--
				left++
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	sums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(sums))
}
