# Level 2：位元運算 💻

> 位元（Bitwise）＋ 二進位數學

## 進度追蹤

**完成：** 0/6 題

---

## ① 191. Number of 1 Bits（Brian Kernighan Algorithm）

- [ ] 完成題目
- [ ] 理解 `n & (n-1)` 的性質

**數學點：**
- `n & (n-1)` 把最低位的 1 消掉
- 這不是技巧，是位元的「數學性質」

**你會學到：**
- 為什麼每次可減少一個 bit
- 演算法複雜度跟 1 的數量有關，而不是 n bit 長度

**LeetCode 連結：** [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)

---

## ② 190. Reverse Bits（位移 + bit-mask）

- [ ] 完成題目
- [ ] 理解 bit reversal

**數學點：**
- 二進位 bit reversal
- 逐位 shift + mask 的本質

**從這題開始，你會感受到 bit 操作的 fun**

**LeetCode 連結：** [190. Reverse Bits](https://leetcode.com/problems/reverse-bits/)

---

## ③ 231. Power of Two（2 的冪次判斷）

- [ ] 完成題目
- [ ] 理解 bit 技巧

**數學點：**
- `n & (n-1) == 0` 判斷是否為 2 的冪
- 2 的冪在二進位只有一個 1

**你會學到：**
- 位元運算的數學性質
- 如何用一行判斷

**LeetCode 連結：** [231. Power of Two](https://leetcode.com/problems/power-of-two/)

---

## ④ 342. Power of Four（4 的冪次判斷）

- [ ] 完成題目
- [ ] 理解 4 的冪的位元特徵

**數學點：**
- 4 的冪 = 2 的偶數冪
- 位元位置的數學規律

**你會學到：**
- 如何在 Power of Two 基礎上加條件
- 位元位置的奇偶性

**LeetCode 連結：** [342. Power of Four](https://leetcode.com/problems/power-of-four/)

---

## ⑤ 137. Single Number II（出現三次，找出現一次）

- [ ] 完成題目
- [ ] 理解位元計數

**數學點：**
- 每個 bit 位計數 % 3
- 位元的模運算

**你會學到：**
- XOR 的延伸：模 3 的位元運算
- 如何用位元解決「出現 k 次」問題

**LeetCode 連結：** [137. Single Number II](https://leetcode.com/problems/single-number-ii/)

---

## ⑥ 260. Single Number III（兩個只出現一次的數）

- [ ] 完成題目
- [ ] 理解分組 XOR

**數學點：**
- 先 XOR 全部得到 `a ^ b`
- 用最低位的 1 分組

**你會學到：**
- 如何用 XOR 分離兩個數
- 位元分組的技巧

**LeetCode 連結：** [260. Single Number III](https://leetcode.com/problems/single-number-iii/)

---

## 🎯 本關重點

完成這 6 題後，你應該掌握：

1. ✅ **Brian Kernighan 算法** - `n & (n-1)` 的神奇性質
2. ✅ **位元反轉** - shift 和 mask 的組合
3. ✅ **冪次判斷** - 用位元特徵判斷 2/4 的冪
4. ✅ **XOR 進階** - 模運算和分組技巧
5. ✅ **位元計數** - 如何用位元解決「出現 k 次」問題

準備好了就前往 [Level 3：數學性質](../level-3-math-properties/README.md) 吧！
