# 930 Binary Subarrays With Sum

Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

A subarray is a contiguous part of the array.

Example 1:

Input: nums = [1,0,1,0,1], goal = 2
Output: 4
Explanation: The 4 subarrays are bolded and underlined below:
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
Example 2:

Input: nums = [0,0,0,0,0], goal = 0
Output: 15
 

題目給定一個二元陣列跟一個整數的目標，要求找出陣列中和為目標值得子陣列數量
陣列有提到　subarray 是連續非空的序列


這邊記錄一下我的思路，並不是一次想到最佳解但心路歷程大概是這樣
```
我剛想到一個做法 但我知道標準做法應該不是這樣 就是說我創建一個嵌套陣列 那key會是 sum 而 value則會是一個array裡面存放了所有sum是這個值的位置
然後我會從左加到右，當我 sum - goal = 某個值的時候
那因為沒有負數的關係
理論上來說，那個k裡面的值數量就會是相應的次數
那再優化
我value裡面存的是出現的次數
這樣應該就可以了
```

```go
// 因為是bianary array 所以不會有負數，那就可以使用雙指針
// 用哈希或者說dp感覺是一定能解
func numSubarraysWithSum(nums []int, goal int) int {
   hmap := make(map[int]int)
   hmap[0] = 1
   sum := 0
   result := 0
   for _, v := range nums {
      sum += v
      
      value, ok := hmap[sum-goal]
      if ok {
        result += value
      }
      hmap[sum] += 1
   }
   return result
}