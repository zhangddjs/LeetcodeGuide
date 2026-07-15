# Stirling

## 相关题

https://leetcode.com/problems/find-the-number-of-possible-ways-for-an-event/description/

## 概念

s(n,k)

- 将n个不同元素刚好分成非空k组的方案数。
- 将 n 个有区别的元素,划分成 k 个无区别的非空子集的方案数。

如果组有标签（顺序有影响）：方案数 = perm(n,k)*s(n,k) == comb(n,k)*f(n,k)

f(n,k) 为满射，即将n个不同元素刚好放进k个有标签的非空组的方案数。

| 数            | 递推                      | 系数含义                                      |
| ---           | ---                       | ---                                           |
| 第二类 S(n,k) | S(n−1,k−1) + k·S(n−1,k)   | 加入已有组:选 k 组之一;新开组:组无标签,系数 1 |
| 满射 f(n,k)   | k·f(n−1,k−1) + k·f(n−1,k) | 两支都带 k:新开的组也要选是哪个舞台           |
| 第一类 `[n,k]`  | `[n−1,k−1]` + (n−1)·`[n−1,k]` | 加入已有轮换:插到 n−1 个元素中某一个的后面    |

## 实现

```py
# Stirling impl
@cache
def s(n, k):
    if n < k: return 0
    if k == 1: return 1
    return (k * s(n - 1, k) + s(n - 1, k - 1))

# Related Function
math.perm(x, k) # 全排列，可以在递推时维护一个perm变量使得查询为O1
math.comb(x, k) # 组合 O(k) 查询。也可以帕斯卡三角形递推C(i,j) = C(i−1,j−1) + C(i−1,j),打表 O(x·k),查询 O(1)
```

