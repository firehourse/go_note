
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps


這是leetcode第70題

這題的點主要在於理解遞迴，
那主要要從結果想回來
當我走到終點前我只會有兩個狀態
分別是走一步跟走兩步
所以答案會是 f(n) = f(n-1) + f(n-2)

那直覺當然就是這樣做
```go
func climbStairs(n int) int {
    // n是要抵達的階梯
    //在到達n之前我有兩種狀態 走一步或走兩步，走到ｎ
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }

    if n > 2 {
       return climbStairs(n-1) + climbStairs(n-2)
    }
    return 0
}
```
但這出現了一個問題就是超時
所以這時候我們就要用到cache的概念
```go
func climbStairs(n int) int {
    // 為了避免超時，我們要宣告一個array來儲存計算過的結果
    dp := make([]int, n+1)
    dp[0] = 1
    if n == 0 {
        return 1
    }
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    dp[1] = 1
    dp[2] = 2
    for i := 3; i <= n; i++{
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

```
這樣就能夠避免重複計算的問題