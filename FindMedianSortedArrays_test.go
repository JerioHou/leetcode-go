/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/
package leetcode_go

import "testing"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	cur := -1
	i, j := 0, 0
	half := (n1 + n2) / 2
	merge := make([]int, half+1)
	for i < n1 && j < n2 && cur < half {
		cur++
		if nums1[i] < nums2[j] {
			merge[cur] = nums1[i]
			i++
		} else {
			merge[cur] = nums2[j]
			j++
		}
	}
	for i < n1 && cur < half {
		cur++
		merge[cur] = nums1[i]
		i++
	}
	for j < n2 && cur < half {
		cur++
		merge[cur] = nums2[j]
		j++
	}
	if (n1+n2)%2 == 0 {
		return float64((merge[cur] + merge[cur-1])) / 2
	} else {
		return float64(merge[cur])
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	//nums1 := []int{1,2}
	nums1 := []int{}
	nums2 := []int{3, 4}
	t.Log(findMedianSortedArrays(nums1, nums2))
}
