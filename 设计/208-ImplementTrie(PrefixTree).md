# [#208 Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/)

![Medium](/figures/Medium.svg)

## 关键词

设计、Trie、前缀树

## 题目

Implement a trie with `insert`, `search`, and `startsWith` methods.

## 简述

**输入：** ------

**输出：** ------

**Notes：**

+ 假设所有输入字符串只包含小写英文单词
+ 所有输入字符串非空

## 思路

本题考察前缀树数据结构的设计，需要知道前缀树的概念和特性才能解决本问题。前缀树是字符串处理中十分重要的数据结构。

Trie是TrieNode组成的集合，首先我们需要知道TrieNode如何定义：

``` java
class TrieNode {
    Map<Character, TrieNode> children;  //字符-节点映射
    boolean endOfWord;  //是否是词尾
}
```

Trie头结点代表空字符`''`，同时也代表空字符为所有插入的字符串的公共前缀。Trie常用的方法有插入、查找、前缀判断和删除，而本题中只需要实现插入、查找和前缀判断。

对于每个插入操作，从左到右遍历单词，对单词中每个字符，从当前节点的孩子节点中寻找该字符对应的节点，从而使得当前路径所组成的字符串为该单词的前缀。如果存在，则说明前缀存在，令当前节点为该节点，否则，创建一个新节点，插入到孩子节点中，并令当前节点为该节点。遍历完单词后令当前节点的`endOfWord`属性为`true`，在查找时便可依此来判断所查单词是否存在。

查找、前缀判断都可以基于插入逻辑来实现。------方法1

## 解决方案

### 方法1-暴力法

根据Trie的特性设计节点，并设计插入、查找、前缀判断的方法。(关键词：Trie)

时间复杂度：$O(n)$ ---32%   _n代表key长度_

空间复杂度：插入$O(n)$ 查找$O(1)$ ---5%

``` java
class Trie {

    public Map<Character, Trie> children;
    public boolean endOfWord;

    /** Initialize your data structure here. */
    public Trie() {
        children = new HashMap<>();
        endOfWord = false;
    }

    /** Inserts a word into the trie. */
    public void insert(String word) {
        Trie cur = this;
        for (char c : word.toCharArray()) {
            cur.children.put(c, cur.children.getOrDefault(c, new Trie()));
            cur = cur.children.get(c);
        }
        cur.endOfWord = true;
    }

    /** Returns if the word is in the trie. */
    public boolean search(String word) {
        Trie cur = this;
        for (char c : word.toCharArray()) {
            cur = cur.children.get(c);
            if (cur == null) return false;
        }
        return cur.endOfWord;
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    public boolean startsWith(String prefix) {
        Trie cur = this;
        for (char c : prefix.toCharArray()) {
            cur = cur.children.get(c);
            if (cur == null) return false;
        }
        return true;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */
```

## 扩展

### 前缀树Trie[$^{[1,2]}$](#refer-anchor-1)

Trie又称前缀树或字典树，是面试中十分重要的数据结构。

Trie树实际上是一个确定有限状态自动机(DFA)。

Trie主要应用场景：

1. Autocomplete(搜索提示)
2. Spell checker
3. IP routing (Longest prefix matching)
4. T9 predictive text(手机9宫格)
5. Solving word games

[Trie教学视频](https://youtu.be/AXjmTQ8LEoI)

[Trie动画演示](https://www.cs.usfca.edu/~galles/visualization/Trie.html)

### 前缀树的删除

删除时，找到最后一个字符对应的节点，如果该节点还有孩子，则直接将endOfWord置为false，否则将其从父节点的孩子中移除。如果移除该节点后父节点变为叶子节点，则从爷爷节点中删除父节点，并进行迭代删除，直到删除当前节点后父节点不再是叶子节点。

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Wikipedia. Trie](https://zh.wikipedia.org/wiki/Trie)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 208-Solution](https://leetcode.com/problems/implement-trie-prefix-tree/solution/)
