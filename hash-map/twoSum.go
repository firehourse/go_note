package hashmap

/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func TwoSum(nums []int, target int) []int {
	//題目輸入nums 與 target都會是int
	// 只會有一個加總結果,且不可使用重複值
	// nums看起來是一個slice
	// target - 輸入 且確認該key是否為存在，即可知道答案，value儲存位置訊息，return遍歷pointer + 位置訊息即可
	// 注意return 要求 slice
	m := make(map[int]int)
	i := 0

	for i < len(nums) {
		check := target - nums[i]

		if value, ok := m[check]; ok {
			return []int{value, i}
		} else {
			m[nums[i]] = i
		}
		i++
	}
	return nil

}
