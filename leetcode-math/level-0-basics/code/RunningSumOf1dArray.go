package code

/*
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.

Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/

func RunningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	numlen := len(nums)
	// 先宣告另一個slice來裝
	result := make([]int, numlen)
	// 先直接給第0位
	result[0] = nums[0]
	// 從第一位開始 每次就是把前面的加起來而已
	for i := 1; i < numlen; i++ {
		result[i] = result[i-1] + nums[i]
	}
	return result
}
