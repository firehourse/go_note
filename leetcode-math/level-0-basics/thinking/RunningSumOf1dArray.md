這題是 leetcode 1480 題

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.

 
 題目會給定一個陣列，這陣列並沒有說會是一個sort過的，而結果則是每個前項的和
 也就是說他是具有順序性的，只需要使用迴圈去累加就可以了


```go
func runningSum(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    numlen := len(nums)
    // 先宣告另一個slice來裝
    result := make([]int,numlen)
    // 先直接給第0位
    result[0] = nums[0]
    // 從第一位開始 每次就是把前面的加起來而已
    for i:= 1;i<numlen;i++ {
        result[i] = result[i-1] + nums [i]
    }
    return result
}
```
那但是這邊其實有個更直接的方式，主要是針對記憶體的優化，也就是你可以直接修改原陣列

```go
func runningSum(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }
    return nums
}
```