# [#698 PartitiontoKEqualSumSubsets](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/)

![Medium](/figures/Medium.svg)

## 关键词

资源分配、数组、判断、等分、寻找false条件、枚举、回溯法、资源分配序列、构造子集合

## 题目

Given an array of integers `nums` and a positive integer `k`, find whether it's possible to divide this array into `k` non-empty subsets whose sums are all equal.

## 简述

**输入：** 带权重的资源数组

**输出：** 是否能平均分成权重相同的k组

**Notes：**

+ 1 <= k <= 资源数 <= 16
+ 0 < 权重 < 10000

## 思路

本题考察资源分配。

按照题意，首先我们可以尝试寻找边界false和true条件，如果要k等分，那么每一组权重之和等于权重总和/k，那么如果权重总和不能被k整除、有元素权重大于权重总和/k、剩余资源无法组成权重总和/k的权重，则不能k等分。但这些条件只能排除掉一些边界情况，而剩下的情况判断则是核心部分。

可以想到一种回溯的暴力法，遍历每一种组合的情况，换句话说也就是所有资源分配的序列组合，只要有一种满足条件，就返回true(例如输入(4,3,2,3,5,2,1)，那么找到的满足条件的组合就是(5),(1,4),(2,3),(2,3)，其资源分配序列之一为(5,1,4,2,3,2,3)，当输入数组可以排序成这个序列时，就能满足)。暴力法时间复杂度巨大，为$O(n^n)$，不做优化会导致超时。------方法1

得到暴力解法后，可以进行一个小优化，可以看出每一个子组合当前元素索引必然在前一个已选元素的索引的后面，因此在遍历时再传一个参数，每当遍历一个新的子组合时，就置为0，否则记录下一层遍历的开始索引。这样可以一定程度减少重复计算。------方法2

进行如上优化后，虽然可以通过，但还是有大量重复计算存在，举个例子就是一开始组成了子数组A，第二个是子数组B，然后递归遍历发现不行，回溯到最初，构成子数组B，然后是子数组A，显然也是不行的，但还是会递归遍历，就会出现重复计算。根据重复计算的情况，往往可以想到用缓存来优化，将不可能的组合记录到缓存中，遍历时判断当前组合是否在缓存，如果在就跳过。

这种解法的难点在于如何存储组合信息，然后是如何判断某组合是否遍历过。

官方解法给出了一个非常巧妙的方式，将访问数组转化成了一个整形used，倒数第i位记录输入数组第i位是否遍历过，再用一个大小为$2^N$的数组存储所有可能的访问情况，这样就可以解决相当一部分重复的运算量。------方法3[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-暴力法

寻找false条件，回溯法遍历，枚举所有可能的分配序列情况，只要有一种符合，则返回true。(关键词：寻找false条件、回溯法、枚举、资源分配序列)

时间复杂度：$O(n^n)$ ---TLE

空间复杂度：$O(n)$ ---TLE

``` java
class Solution {
    public boolean canPartitionKSubsets(int[] nums, int k) {
        boolean [] seen = new boolean[nums.length];
        int max = nums[0], sum = 0, weight = 0;
        for (int num : nums) {
            max = Math.max(num, max);
            sum += num;
        }
        weight = sum / k;
        if (sum % k != 0 || max > weight) return false;
        return dfs(nums, seen, k, weight, weight, 0);
    }

    public boolean dfs(int[] nums, boolean[] seen, int k, int weight, int originWeight, int visited){
        if (weight == 0 && k != 0) weight = --k == 0 ? 0 : originWeight;
        if (visited == nums.length) return !(weight != 0 || k != 0);
        for(int i = 0; i < nums.length; ++i)
            if (!seen[i] && nums[i] <= weight) {
                seen[i] = true;
                if (dfs(nums, seen, k, weight - nums[i], originWeight, visited + 1)) return true;
                seen[i] = false;
            }
        return false;
    }
}
```

### 方法2-优化的暴力法

寻找false条件，回溯法遍历，记录当前所在组合的最右元素索引减少重复计算。(关键词：寻找false条件、回溯法、枚举、资源分配序列)

时间复杂度：$O(n^n)$ ---68%

空间复杂度：$O(n)$ ---21%

``` java
class Solution {
    public boolean canPartitionKSubsets(int[] nums, int k) {
        boolean [] seen = new boolean[nums.length];
        int max = nums[0], sum = 0, weight = 0;
        for (int num : nums) {
            max = Math.max(num, max);
            sum += num;
        }
        weight = sum / k;
        if (sum % k != 0 || max > weight) return false;
        return dfs(nums, seen, k, weight, weight, 0, 0);
    }

    public boolean dfs(int[] nums, boolean[] seen, int k, int weight, int originWeight, int visited, int index){
        if (weight == 0 && k != 0) {
            weight = --k == 0 ? 0 : originWeight;
            index = 0;
        }
        if (visited == nums.length) return !(weight != 0 || k != 0);
        for(int i = index; i < nums.length; ++i)
            if (!seen[i] && nums[i] <= weight) {
                seen[i] = true;
                if (dfs(nums, seen, k, weight - nums[i], originWeight, visited + 1, i + 1)) return true;
                seen[i] = false;
            }
        return false;
    }
}
```

### 方法3-动态规划[$^{[1]}$](#refer-anchor-1)

回溯法遍历，用整形记录输入数组每个位置是否访问过，用备忘录记录所有访问信息及其对应的结果。(关键词：回溯法、动态规划、枚举、资源分配序列)

时间复杂度：$O(n*2^n)$ ---43%

空间复杂度：$O(2^n)$ ---8%

自顶向下：

``` java
enum Result { TRUE, FALSE }

class Solution {
    boolean search(int used, int todo, Result[] memo, int[] nums, int target) {
        if (memo[used] == null) {
            memo[used] = Result.FALSE;
            int targ = (todo - 1) % target + 1;
            for (int i = 0; i < nums.length; i++) {
                if ((((used >> i) & 1) == 0) && nums[i] <= targ) {
                    if (search(used | (1<<i), todo - nums[i], memo, nums, target)) {
                        memo[used] = Result.TRUE;
                        break;
                    }
                }
            }
        }
        return memo[used] == Result.TRUE;
    }
    public boolean canPartitionKSubsets(int[] nums, int k) {
        int sum = Arrays.stream(nums).sum();
        if (sum % k > 0) return false;

        Result[] memo = new Result[1 << nums.length];
        int test = 1 << nums.length;
        memo[(1 << nums.length) - 1] = Result.TRUE;
        return search(0, sum, memo, nums, sum / k);
    }
}
```

自底向上：

``` java
class Solution {
    public boolean canPartitionKSubsets(int[] nums, int k) {
        int N = nums.length;
        Arrays.sort(nums);
        int sum = Arrays.stream(nums).sum();
        int target = sum / k;
        if (sum % k > 0 || nums[N - 1] > target) return false;

        boolean[] dp = new boolean[1 << N];
        dp[0] = true;
        int[] total = new int[1 << N];

        for (int state = 0; state < (1 << N); state++) {
            if (!dp[state]) continue;
            for (int i = 0; i < N; i++) {
                int future = state | (1 << i);
                if (state != future && !dp[future]) {
                    if (nums[i] <= target - (total[state] % target)) {
                        dp[future] = true;
                        total[future] = total[state] + nums[i];
                    } else {
                        break;
                    }
                }
            }
        }
        return dp[(1 << N) - 1];
    }
}

```

## 扩展

### 扩展方法-构造子集合法(Search by Constructing Subset Sums)[$^{[1]}$](#refer-anchor-1)

构造长度为k的数组group，每个位置的值初始为0，且不应大于sum/k。用回溯法遍历整个输入数组，如果能将数组中每个元素刚好全部分配到group中，则返回true。(关键词：构造子集合)

时间复杂度：$O(k^{n-k}k!)$

空间复杂度：$O(n)$

``` java
class Solution {
    public boolean search(int[] groups, int row, int[] nums, int target) {
        if (row < 0) return true;
        int v = nums[row--];
        for (int i = 0; i < groups.length; i++) {
            if (groups[i] + v <= target) {
                groups[i] += v;
                if (search(groups, row, nums, target)) return true;
                groups[i] -= v;
            }
            if (groups[i] == 0) break;
        }
        return false;
    }

    public boolean canPartitionKSubsets(int[] nums, int k) {
        int sum = Arrays.stream(nums).sum();
        if (sum % k > 0) return false;
        int target = sum / k;

        Arrays.sort(nums);
        int row = nums.length - 1;
        if (nums[row] > target) return false;
        while (row >= 0 && nums[row] == target) {
            row--;
            k--;
        }
        return search(new int[k], row, nums, target);
    }
}
```

### 用整形作为访问数组

用32位整形作为访问数组，判断数组第i位是否访问过

``` java
if (!((used >> i) & 1)) used |= (1<<i)    //如果没访问，则访问
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 698-Solution](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/solution/)
