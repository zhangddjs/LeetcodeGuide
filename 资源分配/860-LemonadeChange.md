# [#860 Lemonade Change](https://leetcode.com/problems/lemonade-change/)

![Easy](/figures/Easy.svg)

## 关键词

数组、资源分配、判断、遍历、拆分情况、贪心算法、枚举false条件、HashMap、HashMap优化

## 题目

At a lemonade stand, each lemonade costs `$5`.

Customers are standing in a queue to buy from you, and order one at a time (in the order specified by `bills`).

Each customer will only buy one lemonade and pay with either a `$5`, `$10`, or `$20` bill.  You must provide the correct change to each customer, so that the net transaction is that the customer pays `$5`.

Note that you don't have any change in hand at first.

Return `true` if and only if you can provide every customer with correct change.

## 简述

**输入：** 用户付款数组

**输出：** 能否全部成功找零

**Notes：**

+ 输入数组长度 <= 10000
+ 付款数额只能是5、10、20
+ 一杯汽水价格是5
+ 每个消费者消费1杯汽水

## 思路

本题考察资源分配，判断拥有的资源能否合理分配给所有消费者。

首先我们需要拆分一下情况，当一个消费者付款5元时，不需要找零，付款10元时，手上必须有一张5元找零，付款20元时，手上必须有15元，也就是一张5元和一张10元或3张5元(没有10元就用2张5元)。如果消费者付款10元，手上有1张10元，则不足以找零。可见面值越大，应对情况越少。因此我们可以用一个HashMap保存当前拥有的各数额的零钱数量，当需要找零时，按照贪心策略从map中取零钱即可。------方法1

由于零钱面额固定，并且不是很多，所以可以将HashMap优化成数组，加快读写效率。------方法2

## 解决方案

### 方法1-暴力法

拆分情况，分析不同情况并选择找零策略，遍历数组，贪心算法。(关键词：遍历、拆分情况、贪心算法、枚举false条件、HashMap)

时间复杂度：$O(n)$ ---13%

空间复杂度：$O(1)$ ---98%

``` java
class Solution {
    public boolean lemonadeChange(int[] bills) {
        if(bills.length != 0 && bills[0] != 5) return false;
        Map<Integer, Integer> map = new HashMap<>();
        for (int bill : bills) {
            map.put(bill, map.getOrDefault(bill, 0) + 1);
            int change = bill - 5;
            while (bill - 5 != 0) {
                while (change != 0 && (change >= bill || map.getOrDefault(change, 0) == 0)) change -= 5;
                if (change == 0) return false;
                bill -= change;
                map.put(change, map.get(change) - 1);
            }
        }
        return true;
    }
}
```

### 方法2-优化的暴力法

优化HashMap为数组，加快读写并进一步节省空间(关键词：枚举false条件、HashMap优化)

时间复杂度：$O(n)$ ---78%

空间复杂度：$O(1)$ ---98%

``` java
class Solution {
    public boolean lemonadeChange(int[] bills) {
        if(bills.length != 0 && bills[0] != 5) return false;
        int [] changes = new int[2];
        for (int bill : bills) {
            if(bill != 20) changes[bill / 5 - 1]++;
            switch(bill){
            case 10:
                if (changes[0]-- == 0) return false;
                break;
            case 20:
                if (changes[1] != 0) {
                    if(changes[0]-- == 0) return false;
                    changes[1]--;
                } else if (changes[0] < 3) return false;
                else changes[0] -= 3;
                break;
            }
        }
        return true;
    }
}
```

## 扩展

### 扩展方法1-更少的false条件枚举[$^{[1]}$](#refer-anchor-1)

在遍历数组时先进行计算，在计算结束后判断false条件。

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
public boolean lemonadeChange(int[] bills) {
    int five = 0, ten = 0;
    for (int i : bills) {
        if (i == 5) five++;
        else if (i == 10) {five--; ten++;}
        else if (ten > 0) {ten--; five--;}
        else five -= 3;
        if (five < 0) return false;
    }
    return true;
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 860-Discuss](https://leetcode.com/problems/lemonade-change/discuss/143719/C++JavaPython-Straight-Forward)
