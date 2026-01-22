# DP Problems

## 经典题

- [x] 72. Edit Distance ⭐⭐⭐
- [x] 115. Distinct Subsequences ⭐⭐
- [x] 139. Word Break ⭐⭐ 确实花了点时间走了弯路，需要知道这个题的pattern
- [x] 140. Word Break II ⭐⭐⭐ 可以考虑用Trie来优化字典
- [ ] 312. Burst Balloons ⭐⭐ 逆向思维

### 逆向思维

记忆锚点
区间 DP 的万能公式：
在区间 `[i, j]` 中枚举一个"分割点/决策点"k
- 要么是"第一个被处理的"
- 要么是"最后一个被处理的"
选择让子问题更简单的那个！

关键洞察 1：消除的顺序 vs 添加的顺序
类似的思维转换：
- 汉诺塔：不想"怎么移走"，而想"最大的盘子最后移到哪"
- 构建 BST：不想"怎么删节点"，而想"根节点是哪个"
- 矩阵链乘法：不想"先乘哪两个"，而想"最后一次乘法拆分在哪"

如何训练这种思维？
1. 识别模式："删除"问题常可逆向思考
当题目涉及：
删除、消除、戳破
每次操作改变问题结构
正向思考状态复杂

→ 尝试逆向：最后删除哪个？
2. 问自己关键问题
❌ "第一步做什么？"
✅ "什么是不变的？最后一步是什么？"
对于 Burst Balloons：

不变的：左右边界
最后一步：最后戳破的那个气球，它的邻居是确定的（边界）

3. 题目特征提示
当你看到这些特征时，考虑逆向：
操作顺序影响结果
每次操作改变问题结构
正向递归的子问题定义很复杂
题目有"区间"的感觉

思维触发点：

"先做什么"想不通 → 试试"最后做什么"
子问题边界不清 → 固定某个点（最后操作的点）
相邻关系乱套 → 让"最后的邻居"成为确定的边界

## 建议刷

- [x] 1. LeetCode 322 - Coin Change
- [ ] 2. LeetCode 416 - Partition Equal Subset Sum ⭐⭐⭐⭐⭐ (knapsack)
- [ ] 3. LeetCode 494 - Target Sum (knapsack)
- [ ] 4. LeetCode 139 - Word Break

Hard题精选：
- [ ] LeetCode 124 - Binary Tree Maximum Path Sum
- [ ] LeetCode 787 - Cheapest Flights Within K Stops
- [ ] LeetCode 337 - House Robber III
- [ ] LeetCode 968 - Binary Tree Cameras
- [x] LeetCode 207 - Course Schedule
- [x] LeetCode 210 - Course Schedule II
- [ ] LeetCode 630 - Course Schedule III
- [ ] LeetCode 1462 - Course Schedule IV
- [ ] LeetCode 857 - Minimum Cost to Hire K Workers
- [x] LeetCode 688 - Knight Probability in Chessboard** ⭐⭐⭐⭐⭐
- [ ] LeetCode 837 - New 21 Game

## 反例

1091. Shortest Path in Binary Matrix (因为字问题相互依赖)
