# Level 1：數列與組合 📈

> 開始碰「組合」與「數列」  
> 示範：費波那契、帕斯卡三角

## 進度追蹤

**完成：** 0/5 題

---

## ① 70. Climbing Stairs（費波那契 Fibonacci）

- [ ] 完成題目
- [ ] 理解遞推關係式

**數學點：**
- `f(n) = f(n-1) + f(n-2)`
- 用遞推關係式取代暴力

**你會學到：**
- 什麼是遞推式 recurrence
- 為什麼 Fibonacci 其實描述「有幾種方式」

**LeetCode 連結：** [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

---

## ② 118. Pascal's Triangle（揚輝三角 = 二項式係數）

- [ ] 完成題目
- [ ] 理解組合數性質

**數學點：**
- `C(n,k) = C(n-1, k-1) + C(n-1, k)`

**你會碰到：**
- 組合學（combinatorics）的核心概念
- 之後能接：二項式展開、組合數公式、路徑計數

**LeetCode 連結：** [118. Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/)

---

## ③ 119. Pascal's Triangle II（用數學公式求）

- [ ] 完成題目
- [ ] 優化：用前一個算下一個

**數學點：**
- `C(n,k) = n! / (k! (n-k)!)`
- 看到階乘 factorial

**演算法優化：**
- 用前一個算下一個，不用計階乘

**LeetCode 連結：** [119. Pascal's Triangle II](https://leetcode.com/problems/pascals-triangle-ii/)

---

## ④ 509. Fibonacci Number（標準費波那契）

- [ ] 完成題目（遞迴版本）
- [ ] 完成題目（迭代版本）
- [ ] 完成題目（矩陣快速冪，選修）

**數學點：**
- 經典 Fibonacci 數列
- 遞推 vs 遞迴的差異

**你會學到：**
- 為什麼遞迴會超時
- 如何用迭代優化到 O(n)
- （進階）矩陣快速冪可以到 O(log n)

**LeetCode 連結：** [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

---

## ⑤ 1137. N-th Tribonacci Number（三項遞推）

- [ ] 完成題目
- [ ] 對比 Fibonacci 的差異

**數學點：**
- `T(n) = T(n-1) + T(n-2) + T(n-3)`
- 三項遞推關係

**你會學到：**
- Fibonacci 的延伸
- 多項遞推的通用思維

**LeetCode 連結：** [1137. N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/)

---

## 🎯 本關重點

完成這 5 題後，你應該掌握：

1. ✅ **遞推關係** - Fibonacci 和 Tribonacci 的本質
2. ✅ **組合數** - Pascal's Triangle 的數學意義
3. ✅ **階乘優化** - 如何避免重複計算階乘
4. ✅ **遞迴 vs 迭代** - 理解為什麼遞迴會超時
5. ✅ **動態規劃基礎** - 用前面的結果算後面

準備好了就前往 [Level 2：位元運算](../level-2-bitwise/README.md) 吧！
