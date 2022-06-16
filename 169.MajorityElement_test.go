/**
 *给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/majority-element
 *
 */
package leetcode_go

import (
	"sort"
	"testing"
)

func majorityElement(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return nums[n/2]
}
func majorityElement2(nums []int) int {
	n := len(nums)
	countMap := make(map[int]int)
	for i := 0; i < n; i++ {
		count := countMap[nums[i]] + 1
		if count > n/2 {
			return nums[i]
		} else {
			countMap[nums[i]] = count
		}
	}
	return 0
}

/**
快排
*/
func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, 0, p-1)
	quickSort(nums, p+1, right)

}

func partition(nums []int, left int, right int) int {
	pivot := (nums)[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func TestQuickSort(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3, 2, 1, 3, 4}
	quickSort(nums, 0, len(nums)-1)
	t.Log(nums)
}

func TestMajorityElement(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 2, 3, 3, 4}
	t.Log(majorityElement(nums))
}

func TestMajorityElement2(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 2, 3, 3, 4}
	t.Log(majorityElement2(nums))
}
