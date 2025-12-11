# Ways to Split Array Into Good Subarrays
You are given a binary array nums.

A subarray of an array is good if it contains exactly one element with the value 1.

Return an integer denoting the number of ways to split the array nums into good subarrays. As the number may be too large, return it modulo 109 + 7.

A subarray is a contiguous non-empty sequence of elements within an array.

 

Example 1:

Input: nums = [0,1,0,0,1]
Output: 3
Explanation: There are 3 ways to split nums into good subarrays:
- [0,1] [0,0,1]
- [0,1,0] [0,1]
- [0,1,0,0] [1]


// 題目問的是，把array切分，每個片段剛好都只有１個１，這樣有多少個可行性
// 思考了一下大概想到的是我需要兩個1之間有多少個0 那他就會有幾種解法
// 那看起來其實有點像高中數學的骰骰子，兩個6面的骰子他會有幾種結果這樣
func numberOfGoodSubarraySplits(nums []int) int {
    const MOD = 1_000_000_007

    // 紀錄所有 1 的位置
    var mem []int
    for i, v := range nums {
        if v == 1 {
            mem = append(mem, i)
        }
    }

    // 沒有 1 無法切
    if len(mem) == 0 {
        return 0
    }

    // 只有 1 個 1 只能整段
    if len(mem) == 1 {
        return 1
    }

    // 開始利用兩個 1 之間的距離
    result := 1
    for i := 1; i < len(mem); i++ {
        diff := mem[i] - mem[i-1] 
        result = (result * diff) % MOD
    }

    return result
}
