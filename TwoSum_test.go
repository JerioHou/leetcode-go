package leetcode_go

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	// key得数组的值， value为数组下标
	var valueMap = make(map[int]int)
	for i, v := range nums {
		if index, ok := valueMap[target-v]; ok {
			return []int{i, index}
		} else {
			valueMap[v] = i
		}
	}
	return []int{}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
