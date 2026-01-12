# 树

## 序列化反序列化

297. 序列化反序列化Binary Tree

## Trie

前缀树，快速判断某个单词或前缀是否存在

## Interval Tree

区间树

729 My Calendar I
731 My Calendar II
732 My Calendar III
715 Range Module
56 Merge Intervals
57 Insert Interval
252 Meeting Rooms
253 Meeting Rooms II

- "Interval" = 时间/空间区间对象
- 存储实际的区间对象
- 想成"区间的搜索树"

```
// ✅ 存储和查询区间对象
insert([10, 20])  // 插入会议时间 10:00-20:00
insert([15, 25])  // 插入另一个会议
// ✅ 查找重叠区间
searchOverlap([12, 18])  
// 返回: [[10,20], [15,25]] - 找到所有冲突的会议
```

### 用途

存储和查询**多个区间对象**：
- 查找所有与给定区间重叠的区间
- 插入/删除区间
- 典型应用：日历冲突检测、会议室预订

### 结构

```
区间集合: [15,20], [10,30], [17,19], [5,20], [12,15], [30,40]

区间树（按左端点排序的BST + 每个节点存储子树最大右端点）:

              [10,30] max=40
              /              \
        [5,20] max=30      [30,40] max=40
            \                 /
         [15,20] max=20   [17,19] max=19
             /
        [12,15] max=15
```

## Segment Tree：

线段树 

307 Range Sum Query - Mutable
218 The Skyline Problem (进阶)
699 Falling Squares

- "Segment" = 数组的一段
- 处理数组下标区间
- 想成"数组的树状索引"

```
// ✅ 查询数组区间的聚合信息
sumRange(3, 7)  // 数组下标 3-7 的和
maxRange(2, 5)  // 数组下标 2-5 的最大值
// ✅ 更新数组元素
update(4, 10)   // 将下标 4 的值改为 10
rangeUpdate(2, 5, 3)  // 将下标 2-5 的值都加 3
```

### 用途

处理数组区间查询问题：

区间求和 / 最大值 / 最小值
区间更新（单点更新 / 区间批量更新）
支持懒惰传播（Lazy Propagation）

### 结构

```
数组: [1, 3, 5, 7, 9, 11]

线段树结构:
           [0,5]=36
          /          \
    [0,2]=9          [3,5]=27
    /     \          /      \
[0,1]=4  [2]=5  [3,4]=16  [5]=11
 /   \           /    \
[0]=1 [1]=3   [3]=7  [4]=9
```

## Binary Indexed Tree (BIT) / Fenwick Tree

**何时用 Fenwick 而不是 Segment Tree？**
- ✅ 只需要区间求和/单点更新
- ✅ 代码要简洁
- ❌ 需要区间更新 → 用 Segment Tree

LeetCode: 307, 308, 315, 327

## Union-Find (Disjoint Set Union, DSU)

并查集，用于查询两个元素的集合

## Monotonic Stack/Queue

虽然是栈/队列，但常与树结合

LeetCode 84: Largest Rectangle in Histogram（⭐ Google 超高频）
LeetCode 85: Maximal Rectangle
LeetCode 739: Daily Temperatures
LeetCode 907: Sum of Subarray Minimums

两种类型：
1. 单调递增栈：栈底到栈顶递增
2. 单调递减栈：栈底到栈顶递减

工作原理（以单调递减栈为例）：
- 入栈前，弹出所有小于当前元素的栈顶元素
- 保证栈内始终从底到顶递减
- 常用于找"下一个更大元素"类问题

经典应用场景：
单调栈：
  - 找每个元素左边/右边第一个比它大/小的元素
  - 求以某元素为最小值的最大区间
单调队列：
  核心思想：维护一个元素单调的双端队列，队首始终是当前窗口的最值，用于解决滑动窗口最值问题。
  - 滑动窗口最大值/最小值
  - 区间最值查询
1. 下一个更大/更小元素
2. 柱状图中最大矩形
3. 接雨水问题
4. 每日温度

## Deque

TODO

## Suffix Tree / Suffix Array

TODO

LeetCode 214: Shortest Palindrome
LeetCode 1044: Longest Duplicate Substring（⭐ Google 高频）
字符串匹配问题

## Other

### Merkle Tree

哈希树, 用于数据验证和区块链

应用场景：

- Git 版本控制
- Bitcoin/区块链
- 分布式系统数据一致性验证

### Treap (Tree + Heap)

随机化平衡二叉搜索树

### Cartesian Tree

同时满足堆性质和中序遍历性质

### Splay Tree

伸展树, 自调整的 BST，最近访问的节点移到根

