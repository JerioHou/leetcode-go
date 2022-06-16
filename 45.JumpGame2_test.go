/**
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。



示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:

输入: nums = [2,3,0,1,4]
输出: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode_go

func jump(nums []int) int {
	step, end, maxPosition := 0, 0, 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		maxPosition = Max(i+nums[i], maxPosition)
		if i == end {
			step++
			end = maxPosition
		}
	}
	return step
}

func Max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
