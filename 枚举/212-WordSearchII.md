# [#212 Word Search II](https://leetcode.com/problems/word-search-ii)

![Hard](/figures/Hard.svg)

## 关键词

枚举、查找、遍历、回溯法、访问标记数组、记录路径(前缀)、Trie、Trie的删除、替换和修复

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

这里我们需要引入Trie(前缀树)这个数据结构，其具体实现参考[[引用]208题](/设计/208-ImplementTrie(PrefixTree).md)。如果我们将所有单词放入Trie，那么在深度遍历矩阵时，如果当前字符不属于任意单词的前缀，则回溯，否则继续查找，如果当前字符是某个单词的末尾，则添加到结果集，并将该单词从Trie中删除。

深度遍历时需要记录访问标记数组，和路径(前缀)。------方法2

## 解决方案

### 方法1-暴力回溯法

对每个单词，在二维矩阵中找到其首字母，回溯遍历每种可能的组合，记录能完全匹配的单词。(关键词：遍历、回溯法、访问标记数组)

时间复杂度：$O(m*4*3^{k-1})$ ---_TLE m个单词，长度为k_

空间复杂度：$O(kn)$ ---_TLE 单词长度k，矩阵大小n_

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

### 方法2-Trie优化的回溯法

将所有单词存入Trie，DFS遍历矩阵，当前字符如果是某单词前缀，则继续，否则回溯。记录遍历到当前字符时的路径，将遍历到的单词添加到结果中，并从Trie中删除。(关键词：遍历、回溯法、访问标记数组、记录路径(前缀)、Trie、Trie的删除)

时间复杂度：$O(m*4*3^{k-1})$ ---31%

空间复杂度：$O(kn)$ ---5%

``` java
class Solution {

    class TrieNode {
        Map<Character, TrieNode> map;
        boolean isEndOfWord;

        public TrieNode() {
            map = new HashMap<>();
            isEndOfWord = false;
        }

        public boolean isLeaf() {
            return map.isEmpty();
        }

        public TrieNode get(char c) {
            return map.get(c);
        }

        public void put(char c) {
            map.put(c, new TrieNode());
        }

        public void delete(char c) {
            map.remove(c);
        }
    }

    public List<String> findWords(char[][] board, String[] words) {
        List<String> res = new ArrayList<>();
        TrieNode trie = initTrie(words);
        int m = board.length, n = board[0].length;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                find(board, i, j, new boolean[m * n], trie, res, "");
            }
        }
        return res;
    }

    public TrieNode initTrie(String[] words) {
        TrieNode root = new TrieNode();
        for (String word : words) {
            TrieNode node = root;
            for (char c : word.toCharArray()) {
                if (node.get(c) == null) node.put(c);
                node = node.get(c);
            }
            node.isEndOfWord = true;
        }
        return root;
    }

    public boolean find(char[][] board, int i, int j, boolean[] visited, TrieNode trieNode, List<String> res, String prefix) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length ||
            trieNode.get(board[i][j]) == null || visited[i * board[0].length + j])
            return false;
        visited[i * board[0].length + j] = true;
        TrieNode cur = trieNode.get(board[i][j]);
        prefix = prefix + board[i][j];
        if (cur.isEndOfWord) {
            res.add(prefix);
            cur.isEndOfWord = false;
        }
        if (find(board, i + 1, j, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i + 1][j]);
        if (find(board, i - 1, j, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i - 1][j]);
        if (find(board, i, j + 1, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i][j + 1]);
        if (find(board, i, j - 1, (boolean[]) visited.clone(), cur, res, prefix)) cur.delete(board[i][j - 1]);
        return cur.isLeaf();
    }
}
```

## 扩展

### 扩展方法-继续优化一些空间[$^{[1]}$](#refer-anchor-1)

1. 可以在`isEndOfWord`为`true`的`TrieNode`中缓存好单词，这样就不需要再在回溯法中传递前缀了。
2. 进入下一层回溯前把当前位置置为`'#'`，等下一层回溯全结束后恢复当前位置的字符，这样就可以不用每次都单独复制一个访问标记数组来传递给下一层了。

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 212-Solution](https://leetcode.com/problems/word-search-ii/solution/)
