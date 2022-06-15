/**
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 *
 * 返回这三个数的和。
 *
 * 假定每组输入只存在恰好一个解。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 * 示例 2：
 *
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/3sum-closest
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package leetcode_go

import (
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n, best := len(nums), 10000
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for k > j {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			if myAbs(sum-target) < myAbs(best-target) {
				best = sum
			}
			if sum > target {
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for k > j && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return best
}

func myAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func TestThreeSumClosest(t *testing.T) {
	nums := []int{0, 2, 1, -3}
	t.Log(threeSumClosest(nums, 1))
}
