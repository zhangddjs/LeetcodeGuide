# Ad-hoc

> 判断两数是否同号
a ^ b < 0 异号
a ^ b > 0 同号
(dividend > 0) != (divisor > 0) 异号
dividend * divisor < 0 异号但溢出

> 清楚最右侧的1
n &= (n - 1)

> 公共前缀
```
func rangeBitwiseAnd(left int, right int) int {
    shift := 0
    for left != right {
        left >>= 1
        right >>= 1
        shift++
    }
    return left << shift
}
```

- [x] Integer to English Words - LeetCode 273
- [x] Multiply Strings - LeetCode 43
- [x] Divide Two Integers - LeetCode 29 ⭐️⭐️⭐️ 很好的思路，有点类似line sweep思想，建议学习
- [x] Bitwise AND of Numbers Range - LeetCode 201 ⭐️
- [ ] Gray Code - LeetCode 89


备选（如果需要更多）：

- [ ] Single Number (LeetCode 136)
- [ ] Number of 1 Bits (LeetCode 191)
- [ ] Reverse Bits (LeetCode 190)

- [ ] Excel Sheet Column Number - LeetCode 171 + Excel Sheet Column Title - LeetCode 168 (配套题)
- [ ] Fraction to Recurring Decimal - LeetCode 166
- [ ] UTF-8 Validation - LeetCode 393
