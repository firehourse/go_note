
# LeetCode 第1題 Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

這算是經典題目，引用經典名言．有人相爱，有人夜里开车看海，有人leetcode第一题都做不出来。

hash map 的 kv 映射應該是 web 開發裡面最常見的應用了

其實題目挺簡單的，主要是他已經說了have exactly one solution，僅有唯一解，那要考量的點就少了很多，基本上當時解題思路也都順手記錄了下來
以前用cpp的時候如果int range不大的話用 boolen array 來做，開銷會更小一點，畢竟ｃｐｐ解出來左邊runtime都一堆大牛...但hash map的解法應該是最泛用的

```go
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
