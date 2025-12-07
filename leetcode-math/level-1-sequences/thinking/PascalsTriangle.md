# 帕斯卡三角形 leetcode 118

題目:
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
```
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
```
Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]

這題以前解過但從沒有真的自己從頭到尾想出來過，所以這次決定不依靠AI自己好好思考了一下
首先其實題目已經告訴了你每個數字都是上面的加總，這時候首先會想到的其實是DP跟遞迴
畢竟帕斯卡三角形本身其實就是一個遞迴的思想
但實際在求解的時候他主要是必須要回傳一個會自然生長的二維 array 
感覺有那麼點樹狀結構的感覺
但我這邊使用了for迴圈來解決
只要將他們看做x 跟 y 軸
x軸代表第幾row
y軸代表該row的第幾個數字
然後每個row的長度其實就是row的index + 1
所以我們可以先宣告一個二維陣列
然後針對每個row進行初始化
接著針對每個row的每個數字進行計算
頭尾的數字都是1
中間的數字則是上面兩個數字的加總
這樣就可以完成整個帕斯卡三角形的生成
```go
func generate(numRows int) [][]int {
    // 宣告一個二維陣列 視作x跟y軸
    result := make([][]int, numRows)
    result[0] = []int{1}
    if numRows == 1 {
        return result
    }
    result[1] = []int{1,1}
    if numRows == 2 {
        return result
    }
    for i := 2; i<=numRows-1;i++ {
        // 直接初始化陣列
        result[i] = make([]int,i+1)
        // 內部循環
        for j := 0; j<= i; j++ {
            // 先處理頭尾
            if j == 0 {
                result[i][j] = 1
                continue
            }
            if j == i {
                result[i][j] = 1
                continue
            }
            // 3,1 = 2,0 + 2,1
            result[i][j] = result[i-1][j] + result[i-1][j-1]
        }
    }
    return result
}