# Level 3：數學性質 🔢

> 真正的「數學味」題目（但仍然是入門）

## 進度追蹤

**完成：** 0/8 題

---

## ① 202. Happy Number（數字的數學性質）

- [ ] 完成題目
- [ ] 理解 cycle 檢測

**數學點：**
- 數位拆解
- 重複數列 → cycle 檢測（快慢指標）

**這題讓你感受：**
- 「沒有明顯公式，但用數學性質 + 思考」能解決

**LeetCode 連結：** [202. Happy Number](https://leetcode.com/problems/happy-number/)

---

## ② 204. Count Primes（質數檢測 + 埃氏篩法 Sieve of Eratosthenes）

- [ ] 完成題目
- [ ] 實作埃氏篩法

**數學點：**
- 質數
- √n
- Prime sieve（超經典）

**做完這題你就能看懂所有質數相關題目**

**LeetCode 連結：** [204. Count Primes](https://leetcode.com/problems/count-primes/)

---

## ③ 263. Ugly Number（醜數判斷）

- [ ] 完成題目
- [ ] 理解質因數分解

**數學點：**
- 只包含質因數 2, 3, 5 的數
- 不斷除以 2, 3, 5

**你會學到：**
- 質因數分解的基本思維
- 如何判斷數字的組成

**LeetCode 連結：** [263. Ugly Number](https://leetcode.com/problems/ugly-number/)

---

## ④ 258. Add Digits（數根 Digital Root）

- [ ] 完成題目（循環版本）
- [ ] 完成題目（O(1) 數學公式）

**數學點：**
- 數根 = `(n-1) % 9 + 1`（當 n > 0）
- 模 9 的數學性質

**你會學到：**
- 數位和的數學規律
- 如何從 O(log n) 優化到 O(1)

**LeetCode 連結：** [258. Add Digits](https://leetcode.com/problems/add-digits/)

---

## ⑤ 326. Power of Three（3 的冪次判斷）

- [ ] 完成題目（循環版本）
- [ ] 完成題目（數學技巧版本）

**數學點：**
- 對數判斷：`log₃(n)` 是否為整數
- 或用最大 3 的冪取模

**你會學到：**
- 對數在判斷冪次的應用
- 整數範圍內的數學技巧

**LeetCode 連結：** [326. Power of Three](https://leetcode.com/problems/power-of-three/)

---

## ⑥ 172. Factorial Trailing Zeroes（階乘尾隨零）

- [ ] 完成題目
- [ ] 理解因數 5 的計數

**數學點：**
- 尾隨零 = 因數 5 的個數
- `n/5 + n/25 + n/125 + ...`

**你會學到：**
- 階乘的質因數分解
- 如何計算特定質因數的個數

**LeetCode 連結：** [172. Factorial Trailing Zeroes](https://leetcode.com/problems/factorial-trailing-zeroes/)

---

## ⑦ 67. Add Binary（二進位加法）

- [ ] 完成題目
- [ ] 理解進位邏輯

**數學點：**
- 二進位加法
- 進位 carry 的處理

**你會學到：**
- 模擬加法運算
- 字串處理 + 數學運算結合

**LeetCode 連結：** [67. Add Binary](https://leetcode.com/problems/add-binary/)

---

## ⑧ 69. Sqrt(x)（平方根）

- [ ] 完成題目（二分搜尋）
- [ ] 完成題目（牛頓法，選修）

**數學點：**
- 二分搜尋的數學應用
- 牛頓迭代法：`x_{n+1} = (x_n + a/x_n) / 2`

**你會學到：**
- 如何用二分逼近答案
- 牛頓法的數學原理

**LeetCode 連結：** [69. Sqrt(x)](https://leetcode.com/problems/sqrtx/)

---

## 🎯 本關重點

完成這 8 題後，你應該掌握：

1. ✅ **質數** - 埃氏篩法的原理和實作
2. ✅ **數根** - 模 9 的數學性質
3. ✅ **冪次判斷** - 對數和取模的應用
4. ✅ **階乘性質** - 質因數分解的思維
5. ✅ **二進位運算** - 加法和進位的處理
6. ✅ **數值逼近** - 二分搜尋和牛頓法

準備好了就前往 [Level 4：組合與動態規劃](../level-4-combinatorics-dp/README.md) 吧！
