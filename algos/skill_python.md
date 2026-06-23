# Python 技巧

## 统计list里每个元素的数量

Counter

```py
nums: List[int]
for v, c in Counter(nums).items()

#可以相减
def canConstruct(self, ransomNote, magazine):
    return not collections.Counter(ransomNote) - collections.Counter(magazine)

# 取value最大的
c.most_common()
c.most_common(2)# 只取前2
```

Sum

```py
def findDegrees(self, matrix: list[list[int]]) -> list[int]:
  return [sum(row) for row in matrix]
  # return [Counter(row)[1] for row in matrix]  # 能用，但没必要
```

## 任一

any

```py
any(c in ps or all(c % p for p in ps) for v, c in Counter(nums).items() if c > 1)
```

## 相邻/间隔/双指针/矩阵

```py
# ================================
# 用zip_longest同时取两个字符串
def mergeAlternately(self, w1, w2):
    return ''.join(a + b for a, b in zip_longest(w1, w2, fillvalue=''))

# zip_longest → ('a','p'), ('b','q'), ('','r'), ('','s')
# a+b         →  "ap"      "bq"       "r"       "s"
# join        → "apbqrs"

# ================================
# zip 同时取相邻的元素
def romanToInt(self, s: str) -> int:
  res = 0
  roman = {
      'I': 1, 'V': 5, 'X': 10, 'L': 50,
      'C': 100, 'D': 500, 'M': 1000
  }

  for a, b in zip(s, s[1:]):
      if roman[a] < roman[b]: res -= roman[a]
      else: res += roman[a]
  return res + roman[s[-1]]

# pairwise 同时取相邻元素
# A = [1, 3, 2, 5]
# pairwise(A) → (1,3), (3,2), (2,5)
def minOperations(self, A: list[int]) -> int:
    return sum(max(0, a - b) for a, b in pairwise(A))


# ================================
# 用迭代器遍历
def isSubsequence(self, s: str, t: str) -> bool:
  it = iter(t)
  return all(c in it for c in s)

# ================================
# 多指针 经典应用 转置矩阵
matrix = [[1,2,3],[4,5,6],[7,8,9]]
transposed = list(zip(*matrix))
# [(1,4,7), (2,5,8), (3,6,9)]
[*zip(*matrix)][::-1] # 旋转矩阵 逆时针90度, 先转置再上下翻转。
[*zip(*matrix[::-1])] # 顺时针90°, 先上下翻转再转置
[row[::-1] for row in matrix[::-1]]  # 旋转180度 上下翻转 + 每行反转
matrix[:] = zip(*matrix[::-1]) # "就地"
# zip(*matrix) — O(m×n)，遍历所有元素
# [::-1] — O(m) 或 O(n)，只是反转行/列的顺序，可以忽略
# [*...] — O(m×n)，解包
```

```py
# ================================
# 将两个数组搞成字典，分数 -> 名次 场景
dict(zip(sort, rank))
# 用原始分数每个元素去查表
list(map(dict(zip(sort, rank)).get, nums))
# 等价于
res = []
for score in nums:
    res.append(d.get(score))

```

## 切片操作

```py
[start:stop:step]
a = [1, 2, 3, 4, 5]
a[::-1]   # [5, 4, 3, 2, 1] 反转
a[::2]    # [1, 3, 5]，每隔一个取一个
a[1::2]   # [2, 4]，从索引1开始每隔一个
a[::-2]   # [5, 3, 1]，倒序每隔一个
"hello"[::-1]  # "olleh"
```

## 取反妙用

```py
class Solution:
    def rotate(self, A):
        n = len(A)
        for i in range(n/2):
            for j in range(n-n/2):
                A[i][j], A[~j][i], A[~i][~j], A[j][~i] = \
                         A[~j][i], A[~i][~j], A[j][~i], A[i][j]

# ~i 是按位取反，等于 -(i+1)
# 在这里利用的是 Python 负数索引的特性：
# A[-1]   # 最后一个元素
# A[-2]   # 倒数第二个
# 所以 ~i 就是 A[-(i+1)]，即从末尾数第 i+1 个，等价于 A[n-1-i]：
# A[~i]  ==  A[-(i+1)]  ==  A[n-1-i]
# A[~j]  ==  A[-(j+1)]  ==  A[n-1-j]

class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = [c.lower() for c in s if c.isalnum()]
        return all (s[i] == s[~i] for i in range(len(s)//2))
```

## 重叠元素

```py
# map(func, iterable) 把 func 依次应用到 iterable 每个元素上。


J = "aA"
S = "aAAbbbb"

sum(s in J for s in S)
# → True, True, True, False, False, False, False
sum(map(J.count, S))
# map(lambda s: J.count(s), S)  # 完全等价的写法
# 对S每个字符统计在J中的次数, O(mn)
# → [1, 1, 1, 0, 0, 0, 0]

sum(map(S.count, J))
# 对J每个字符统计在S中的次数, O(mn)
# → [1, 2]
```

## 邻接表 defaultdict 妙用

```python
# adj = {} wrong
adf = defaultdict(list)
for u,v in edges:
    adj[u].append(v)
    adj[v].append(u)

# 使用prev 节省掉visited set
def dfs(node: int, prev: int) -> int:
    d = 0
    for c in graph[node]:
        if c != prev:
            d = max(d, dfs(c, node) + 1)
    return d
```

