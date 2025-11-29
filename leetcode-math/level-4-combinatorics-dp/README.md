# Level 4：組合與動態規劃 🎲

> 組合＋高級 DP（完成前面後再來）

## 進度追蹤

**完成：** 0/9 題

---

## ① 62. Unique Paths

- [ ] 完成題目（DP 解法）
- [ ] 完成題目（組合公式解法）

**數學點：**
- 典型 combinatorial counting
- `C(m+n-2, m-1)`
- 完全用組合公式就能 O(1) 計算

**LeetCode 連結：** [62. Unique Paths](https://leetcode.com/problems/unique-paths/)

---

## ② 343. Integer Break（最大乘積 → 數學最大化）

- [ ] 完成題目
- [ ] 理解為什麼拆成 3 最大

**數學點：**
- 拆成 3 會最大
- 微積分思想（不用算微積分）
- 數學優化題

**LeetCode 連結：** [343. Integer Break](https://leetcode.com/problems/integer-break/)

---

## ③ 279. Perfect Squares（完全平方數）

- [ ] 完成題目（DP 解法）
- [ ] 理解四平方和定理（選修）

**數學點：**
- 拉格朗日四平方和定理
- 動態規劃 + 數學優化

**你會學到：**
- 如何用 DP 解決數學問題
- 數學定理如何優化演算法

**LeetCode 連結：** [279. Perfect Squares](https://leetcode.com/problems/perfect-squares/)

---

## ④ 50. Pow(x, n)（快速冪）

- [ ] 完成題目
- [ ] 理解快速冪原理

**數學點：**
- 快速冪：`x^n = (x^2)^(n/2)`
- 分治思想

**你會學到：**
- 如何從 O(n) 優化到 O(log n)
- 指數運算的數學技巧

**LeetCode 連結：** [50. Pow(x, n)](https://leetcode.com/problems/powx-n/)

---

## ⑤ 372. Super Pow（超級次方）

- [ ] 完成題目
- [ ] 理解模運算性質

**數學點：**
- `(a * b) % m = ((a % m) * (b % m)) % m`
- 快速冪 + 模運算

**你會學到：**
- 模運算的分配律
- 如何處理超大指數

**LeetCode 連結：** [372. Super Pow](https://leetcode.com/problems/super-pow/)

---

## ⑥ 264. Ugly Number II（第 n 個醜數）

- [ ] 完成題目
- [ ] 理解三指標法

**數學點：**
- 動態生成醜數序列
- 三指標合併

**你會學到：**
- 如何動態生成特定數列
- 多指標合併的技巧

**LeetCode 連結：** [264. Ugly Number II](https://leetcode.com/problems/ugly-number-ii/)

---

## ⑦ 313. Super Ugly Number（k 個質因數的醜數）

- [ ] 完成題目
- [ ] 理解多指標擴展

**數學點：**
- Ugly Number II 的一般化
- k 指標合併

**你會學到：**
- 如何將演算法一般化
- 動態規劃的擴展思維

**LeetCode 連結：** [313. Super Ugly Number](https://leetcode.com/problems/super-ugly-number/)

---

## ⑧ 168. Excel Sheet Column Title（十進位 → 26 進位）

- [ ] 完成題目
- [ ] 理解特殊進位制

**數學點：**
- 反向的 171 題
- 沒有 0 的 26 進位（1-26 而非 0-25）

**你會學到：**
- 特殊進位制的轉換
- 為什麼要 `(n-1) / 26`

**LeetCode 連結：** [168. Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/)

---

## ⑨ 29. Divide Two Integers（不用除法實作除法）

- [ ] 完成題目
- [ ] 理解位移優化

**數學點：**
- 用減法 + 位移實作除法
- 倍增思想

**你會學到：**
- 如何用基本運算實作複雜運算
- 位移的數學意義（乘以 2 的冪）

**LeetCode 連結：** [29. Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)

---

## 🎯 本關重點

完成這 9 題後，你應該掌握：

1. ✅ **組合計數** - 路徑問題的組合公式
2. ✅ **數學優化** - 如何用數學找最大值
3. ✅ **快速冪** - 分治思想的經典應用
4. ✅ **模運算** - 處理超大數字的技巧
5. ✅ **動態生成** - 多指標合併的思維
6. ✅ **特殊進位制** - 沒有 0 的進位系統
7. ✅ **位移技巧** - 用位移實作乘除法

準備好了就前往 [Level 5：進階數學挑戰](../level-5-advanced/README.md) 吧！
