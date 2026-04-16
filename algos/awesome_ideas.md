# Awesome Ideas

## 矩阵转复数坐标系

[longestIncreasingPath](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/solutions/78381/short-python-by-stefanpochmann-4cow)

这是一个用复数表示二维坐标的巧妙实现，解题思路是拓扑排序（从小到大处理）。


```py
def longestIncreasingPath(self, matrix):
    matrix = {i + j*1j: val
              for i, row in enumerate(matrix)
              for j, val in enumerate(row)}
    length = {}
    for z in sorted(matrix, key=matrix.get):
        length[z] = 1 + max([length[Z]
                             for Z in z+1, z-1, z+1j, z-1j
                             if Z in matrix and matrix[z] > matrix[Z]]
                            or [0])
    return max(length.values() or [0])
```

**第一步：把矩阵转成复数字典**

```py
matrix = {i + j*1j: val ...}
```

用复数 i + j*1j 表示坐标 (row, col)，实部是行，虚部是列。比如 (2, 3) → 2 + 3j。结果是 {复数坐标: 值} 的字典。

**第二步：按值从小到大遍历**

```py
for z in sorted(matrix, key=matrix.get):
```

从最小值开始处理，保证处理当前格子时，所有比它小的邻居已经算好了。

**第三步：计算每个格子的最长路径**

```py
length[z] = 1 + max([length[Z]
                     for Z in z+1, z-1, z+1j, z-1j
                     if Z in matrix and matrix[z] > matrix[Z]]
                    or [0])
```
- z+1, z-1, z+1j, z-1j 是上下左右四个邻居（实部±1是上下，虚部±1是左右）
- 只看比当前格子值小的邻居（`matrix[z] > matrix[Z]`），取其中最长路径 +1
- `or [0]` 处理没有合法邻居的情况，此时路径长度为 1

> 技巧本质是在DAG上做DP, 不是所有 DFS 都能这样做，关键条件是DAG（有向无环图）。

时间复杂度会比普通缓存DP稍微高一些，O(mnlog(mn))


## 元组切片赋值技巧

[summaryRanges](https://leetcode.com/problems/summary-ranges/solutions/63193/6-lines-in-python-by-stefanpochmann-nwhk)

```py
def summaryRanges(self, nums):
    ranges = []
    for n in nums:
        if not ranges or n > ranges[-1][-1] + 1:
            ranges += [],
        ranges[-1][1:] = n,
    return ['->'.join(map(str, r)) for r in ranges]
```

这个解法很巧妙，逐行解释一下：

```py
ranges += [],
```

- 等价于 ranges.append([])
- 注意末尾的逗号，使得 [] 被当作元组 ([],) 添加，这是一个 Python 小技巧

```py
ranges[-1][1:] = n,
```

- n, 是单元素元组 (n,)
- 对最后一个 range 做切片赋值
- 如果该 range 是 []（刚创建），`[1:]` = (n,) → `[n]`，即只有起点
- 如果该 range 是 `[start]`，`[1:]` = (n,) → `[start, n]`，更新终点
- 每次遇到连续数字都只更新终点，非常简洁

```py
'->'.join(map(str, r))
```

- range 只有一个元素 `[start]` → "start"
- range 有两个元素 `[start, end]` → "start->end"

整体思路就是：维护一个区间列表，每个区间用 [start] 或 [start, end] 表示，遇到不连续的就开新区间，否则更新终点。切片赋值这个写法把"初始化"和"更新"合并成了一行，是这段代码最核心的技巧。


## 矩阵旋转

[spiralOrder](https://leetcode.com/problems/spiral-matrix/solutions/20571/1-liner-in-python-ruby-by-stefanpochmann-rqep)

```py
def spiralOrder(self, matrix):
    return matrix and [*matrix.pop(0)] + self.spiralOrder([*zip(*matrix)][::-1])
```

`zip(*matrix)[::-1]` 实现了矩阵的逆时针旋转90°：

`zip(*matrix)` 是转置（行变列）
`[::-1]` 是上下翻转
两步合起来 = 逆时针旋转

所以整个递归的逻辑是：

取出第一行（最上层）
把剩余矩阵逆时针旋转90°
递归处理旋转后的矩阵，继续取第一行

每次旋转后原来的左列变成了新的第一行，正好对应螺旋顺序：上→右→下→左。

**复杂度分析：**
你的递归版本（spiralOrderSingle）：

时间：O(m×n)，每个元素访问一次
空间：O(min(m,n))，递归深度

旋转版本：

时间：O(m×n×min(m,n))，每次 zip(*matrix) 需要遍历整个剩余矩阵，一共旋转 min(m,n) 次
空间：O(m×n)，每次 zip 都创建新的矩阵
