package leetcode_go

func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	k := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			k = i
			break
		}
	}
	if k == 0 {
		return binarySearch(nums, target)
	}
	if nums[k] > target {
		return -1
	}
	if nums[0] > target {
		result := binarySearch(nums[k:], target)
		if result == -1 {
			return result
		} else {
			return result + k
		}
	} else {
		return binarySearch(nums[:k], target)
	}
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for high >= low {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
