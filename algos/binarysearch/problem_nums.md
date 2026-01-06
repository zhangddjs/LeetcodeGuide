# Binary Search

## 好题列表

- 找第一个
LeetCode 35 - 搜索插入位置
LeetCode 162 - 寻找峰值
LeetCode 69 - x 的平方根

- 找精确值⚠️
LeetCode 704 - 二分查找

- 最后一个
LeetCode 34 - 在排序数组中查找元素的第一个和最后一个位置
LeetCode 275 - H指数 II（找最后一个满足条件的位置）
LeetCode 278 - 第一个错误的版本（找第一个 bad）

## 左闭右闭 vs 左闭右开的选择

> 这纯粹是编码风格选择，不存在某题更适合某种写法。建议选定一种风格并坚持使用，避免混乱。
> 真正需要区分的情况：目标是"找到精确值" vs "找边界/最优解"

LeetCode 35 - 搜索插入位置

用两种解法解出来: left, right = 0, len(nums) - 1 or left, right = 0, len(nums)

```py
# 左闭右闭 [left, right]
def searchInsert(nums, target):
    left, right = 0, len(nums) - 1  # right = len - 1
    while left <= right:  # 可以相等
        mid = (left + right) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return left

# 左闭右开 [left, right)
def searchInsert(nums, target):
    left, right = 0, len(nums)  # right = len
    while left < right:  # 不能相等
        mid = (left + right) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid  # 不是 mid - 1
    return left
```

## 循环条件 left <= right vs left < right

LeetCode 69 - x 的平方根

```py
# left <= right：区间完全消失才停止
def mySqrt(x):
    left, right = 0, x
    while left <= right:
        mid = (left + right) // 2
        if mid * mid <= x:
            left = mid + 1
        else:
            right = mid - 1
    return right  # 返回 right

# left < right：区间剩一个元素时停止
def mySqrt(x):
    left, right = 0, x + 1
    while left < right:
        mid = (left + right) // 2
        if mid * mid <= x:
            left = mid + 1
        else:
            right = mid
    return left - 1  # 返回 left - 1
```

> 理论上，所有二分查找题都可以用 while left <= right 解决。但在某些场景下，while left < right 会更简洁、更直观。

什么时候 while left < right 更优？ 当满足以下条件时：
- 答案一定存在（不会返回 -1）
- mid 本身可能是答案（不能直接排除 mid）
- 最终返回的是指针位置（不是具体值的查找）

### 场景1：找边界值（返回的是指针位置）

LeetCode 162 - 寻找峰值

```py
# 写法1：while left < right（简洁）
def findPeakElement(nums):
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] < nums[mid + 1]:
            left = mid + 1
        else:
            right = mid
    return left  # 循环结束时 left == right，直接返回

# 写法2：while left <= right（需要额外逻辑）
def findPeakElement(nums):
    left, right = 0, len(nums) - 1
    while left <= right:
        if left == right:  # 需要特判
            return left
        mid = (left + right) // 2
        if nums[mid] < nums[mid + 1]:
            left = mid + 1
        else:
            right = mid - 1  # 问题：可能错过答案！
    return left  # 还需要边界处理
```

### 场景2：二分答案（答案不在数组中，而是某个满足条件的值）

```py
# 写法1：while left < right（简洁）
def minEatingSpeed(piles, h):
    left, right = 1, max(piles)
    while left < right:
        mid = (left + right) // 2
        if canFinish(piles, mid, h):
            right = mid  # mid 可能是答案
        else:
            left = mid + 1
    return left  # 循环结束时 left 就是最小速度

def canFinish(piles, speed, h):
    return sum((p + speed - 1) // speed for p in piles) <= h

# 写法2：while left <= right（需要额外变量）
def minEatingSpeed(piles, h):
    left, right = 1, max(piles)
    result = right  # 需要额外变量记录答案
    while left <= right:
        mid = (left + right) // 2
        if canFinish(piles, mid, h):
            result = mid  # 记录可能的答案
            right = mid - 1  # 继续找更小的
        else:
            left = mid + 1
    return result  # 返回记录的答案
```


## 边界更新 left = mid + 1 vs left = mid

**关键**：当使用 left = mid 或 right = mid 时：
- 必须配合 while left < right（不能用 <=）
- 计算 mid 时可能需要上取整：mid = (left + right + 1) // 2
  配对规则是：
  left = mid ↔ 上取整
  right = mid ↔ 下取整

**实践建议：**
- 精确搜索建议使用 while left <= right 配 mid ± 1
- 如果用 while left < right，记住配对规则：
  - left = mid 必配上取整
  - right = mid 必配下取整

### mid 不可能是答案时 → 用 mid ± 1

LeetCode 69 - x 的平方根

```py
def mySqrt(x):
    left, right = 0, x
    while left <= right:
        mid = (left + right) // 2
        if mid * mid <= x:
            left = mid + 1  # mid² 太小，答案在右边
        else:
            right = mid - 1  # mid² 太大，答案在左边
    return right  # right 是最后一个满足条件的
```

### mid 可能是答案时 → 用 mid（但要防死循环）

LeetCode 162 - 寻找峰值

这道题需要仔细处理边界避免死循环：

```py
# 错误写法 - 会死循环
def findPeakElement(nums):
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] < nums[mid + 1]:
            left = mid  # ❌ 当 left=0, right=1 时会死循环
        else:
            right = mid
    return left

# 正确写法
def findPeakElement(nums):
    left, right = 0, len(nums) - 1
    while left < right:
        mid = (left + right) // 2
        if nums[mid] < nums[mid + 1]:
            left = mid + 1  # ✓ 正确
        else:
            right = mid
    return left
```

### 需要上取整的场景

LeetCode 275 - H指数 II（找最后一个满足条件的位置）

```py
# ❌ 错误写法 - 会死循环
def hIndex(citations):
    n = len(citations)
    left, right = 0, n - 1
    
    while left < right:
        mid = (left + right) // 2  # 下取整
        
        if citations[mid] >= n - mid:
            right = mid  # mid 可能是答案，保留
        else:
            left = mid  # ⚠️ 这里是关键！
    
    return n - left

# 测试：citations = [0, 1]
# left=0, right=1
# mid = (0+1)//2 = 0
# 如果走到 left = mid，则 left=0, right=1
# 下一轮：mid = (0+1)//2 = 0 ← 又是0！
# 又走到 left = mid，则 left=0, right=1
# 死循环！

# ✅ 正确写法 - 上取整
def hIndex(citations):
    n = len(citations)
    left, right = 0, n - 1
    
    while left < right:
        mid = (left + right + 1) // 2  # 上取整！
        
        if citations[mid] >= n - mid:
            right = mid - 1
        else:
            left = mid  # 现在安全了
    
    return n - left if citations[left] >= n - left else 0

# 测试：citations = [0, 1]
# left=0, right=1
# mid = (0+1+1)//2 = 1 ← 上取整得到1
# 假设走到 left = mid，则 left=1, right=1
# 循环结束 ✓
```

LeetCode 278 - 第一个错误的版本（找第一个 bad）

```py
# 找第一个 True
def firstBadVersion(n):
    left, right = 1, n
    while left < right:
        mid = (left + right) // 2  # 下取整
        if isBadVersion(mid):
            right = mid  # mid 可能是第一个，保留
        else:
            left = mid + 1  # mid 不是，排除
    return left

# 这里用 right = mid，所以配下取整，没问题。
# 但如果改成找最后一个 Good：

# 找最后一个 False
def lastGoodVersion(n):
    left, right = 1, n
    while left < right:
        mid = (left + right + 1) // 2  # 必须上取整！
        if isBadVersion(mid):
            right = mid - 1
        else:
            left = mid  # mid 可能是最后一个 good
    return left
```


## 找第一个 vs 最后一个满足条件

LeetCode 34 - 在排序数组中查找元素的第一个和最后一个位置

```py
def searchRange(nums, target):
    # 找第一个 >= target 的位置（左边界）
    def findLeft():
        left, right = 0, len(nums)
        while left < right:
            mid = (left + right) // 2
            if nums[mid] < target:  # 严格小于
                left = mid + 1
            else:  # >= target 时收缩右边界
                right = mid
        return left
    
    # 找最后一个 <= target 的位置（右边界）
    def findRight():
        left, right = 0, len(nums)
        while left < right:
            mid = (left + right) // 2
            if nums[mid] <= target:  # 小于等于
                left = mid + 1  # 继续向右找
            else:
                right = mid
        return left - 1  # 减 1 得到最后一个 <= target
    
    left_idx = findLeft()
    if left_idx >= len(nums) or nums[left_idx] != target:
        return [-1, -1]
    return [left_idx, findRight()]
```

## "找到精确值" vs "找边界/最优解"

### 找精确值（找到就返回）

LeetCode 704 - 二分查找

```py
def search(nums, target):
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid  # 找到就返回
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1  # 没找到
```


### 找边界/最优解（需要遍历完整个可能区间）

LeetCode 34 - 查找第一个和最后一个位置


## 总结

### 核心选择策略

| 场景       | 循环条件      | 边界更新    | mid计算   | 何时使用                 |
| ------     | ---------     | ---------   | --------- | ---------                |
| 找精确值⚠️ | left <= right | mid ± 1     | 下取整    | 标准查找，找到即返回     |
| 找第一个   | left < right  | right = mid | 下取整    | 边界查找，答案确保存在   |
| 找最后一个 | left < right  | left = mid  | 上取整⚠️  | 边界查找，答案确保存在   |
| 二分答案   | left < right  | right = mid | 下取整    | 最优化问题，答案在范围内 |

### 死循环避免

| 危险组合                          | 死循环条件               | 解决方案          |
| ----------                        | ------------             | ----------        |
| left = mid & 下取整               | left=0, right=1 时 mid=0 | 改用上取整        |
| while left <= right & right = mid | mid 是答案时被跳过       | 改用 left < right |

### 配对规则（记忆口诀）

- **left = mid** ↔ **上取整**
- **right = mid** ↔ **下取整** 
- **mid ± 1** ↔ **下取整**

### 快速决策树

```
是否精确查找？
├─ 是 → while left <= right + mid ± 1
└─ 否（找边界/最优）
   ├─ 找第一个 → while left < right + right = mid + 下取整
   └─ 找最后一个 → while left < right + left = mid + 上取整
```


