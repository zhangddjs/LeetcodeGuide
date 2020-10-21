# [#212 Word Search II](https://leetcode.com/problems/word-search-ii)

![Hard](/figures/Hard.svg)

## 关键词

查找、回溯法、访问标记数组、

## 题目

Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

## 简述

**输入：** 二维字符矩阵; 目标单词

**输出：** 按题目条件能在矩阵中匹配的目标单词

**Notes：**

+ 矩阵中所有字符为小写英文字符
+ 每个目标单词唯一

## 思路

本题考察查找。

容易想到的暴力法是对于二维矩阵中每个字符可能组成的所有单词情况添加到HashSet中，然后对于每个单词，直接判断Set中是否存在即可，但这样做的时间和空间开销十分巨大，实现困难。

于是进行初步的优化操作，即对于每个目标单词，找到其首字母在二维矩阵中所有的出现位置，然后从每个位置出发，通过回溯法遍历上下左右4个方向，从能够匹配下一个字符的方向继续往下遍历，直到遍历路径和目标单词完全匹配，则添加到结果队列。如果所有情况都不能完全匹配当前单词，则当前目标单词不存在，继续下一个单词。在遍历过程中，通过访问标记数组表示每个位置是否被访问过。------方法1

但该方法有明显的弊端，就是如果单词量比较多，同时大量单词拥有共同前缀和大量重复字符时，重复的运算量将会非常可观。因此我们需要继续进行优化。

减少重复运算的常用方法是对运算过的子问题的结果（前缀）进行缓存，那么本题什么时候进行缓存，用什么样的数据结构来存储子问题结果呢？



## 解决方案

### 方法1-回溯法

对每个单词，在二维矩阵中找到其首字母，回溯遍历每种可能的组合，记录能完全匹配的单词。(关键词：回溯法、访问标记数组)

时间复杂度：$O()$ ---_TLE_

空间复杂度：$O()$ ---_TLE_

``` java
class Solution {
    public List<String> findWords(char[][] board, String[] words) {
        List<String> res = new ArrayList<>();
        for (String word : words) {
            boolean flag = false;
            for (int i = 0; i < board.length && !flag; ++i) {
                for (int j = 0; j < board[0].length && !flag; ++j) {
                    if (board[i][j] == word.charAt(0) &&
                        findWord(board, word, i, j, new boolean[board.length * board[0].length])) {
                        res.add(word);
                        flag = true;
                    }
                }
            }
        }
        return res;
    }

    public boolean findWord(char[][] board, String word, int i, int j, boolean[] visited) {
        if (word == null || word.equals("")) return true;
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length ||
            visited[i * board[0].length + j] || board[i][j] != word.charAt(0))
            return false;
        visited[i * board[0].length + j] = true;
        return findWord(board, word.substring(1), i + 1, j, (boolean[]) visited.clone()) ||
               findWord(board, word.substring(1), i - 1, j, (boolean[]) visited.clone()) ||
               findWord(board, word.substring(1), i, j + 1, (boolean[]) visited.clone()) ||
               findWord(board, word.substring(1), i, j - 1, (boolean[]) visited.clone());
    }
}
```

## 扩展

### TODO扩展一些知识和方法[$^{[1]}$](#refer-anchor-1)

内容

``` java
/**
 * Copyright: LeetCode(https://leetcode.com)
 * Author: covfefe
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
//Extension Solution
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 0-Solution]()
