# [#151 Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、字符串、逆置、拆分字符串、反向遍历、拼接字符串、正则、字符处理、双端队列Deque、两次逆置

## 题目

Given an input string, reverse the string word by word.

## 简述

**输入：** 字符串

**输出：** 将每个单词逆置后的字符串

**Notes：**

+ 单词为一个无空格字符序列
+ 输入字符串可能在开头或结尾包含多个空格，逆置后这些空格不复存在
+ 单词间的连续空格转换成单个空格

## 思路

本题考察字符串转换。

根据单词间空格分隔，单词内无空格等特性，可以想到以空格作为分隔符将字符串拆分成单词数组，然后再用从后往前遍历单词数组实现逆置，拼接到结果字符串即可获得结果。需要注意多个空格连续的情况。------方法1

## 解决方案

### 方法1-暴力法

将字符串按照空格拆分成单词数组，从后往前遍历并进行拼接。(关键词：拆分字符串、反向遍历、拼接字符串)

时间复杂度：$O(n)$ ---55%

空间复杂度：$O(n)$ ---88%

``` java
class Solution {
    public String reverseWords(String s) {
        String[] words = s.trim().split("\\s+");
        StringBuilder res = new StringBuilder();
        for (int i = words.length - 1; i >= 0; --i) res.append(words[i] + " ");
        return res.toString().trim();
    }
}
```

3-line写法:

``` java
class Solution {
    public String reverseWords(String s) {
        StringBuilder res = new StringBuilder();
        for (String word : s.trim().split("\\s+")) res.insert(0, " " + word);
        return res.toString().trim();
    }
}
```

## 扩展

### 扩展方法1-逆置整个字符串后逆置每个单词[$^{[1]}$](#refer-anchor-1)

去除两端空格，逆置整个字符串，再对每个单词进行逆置。(关键词：字符处理、两次逆置)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
  public StringBuilder trimSpaces(String s) {
    int left = 0, right = s.length() - 1;
    // remove leading spaces
    while (left <= right && s.charAt(left) == ' ') ++left;

    // remove trailing spaces
    while (left <= right && s.charAt(right) == ' ') --right;

    // reduce multiple spaces to single one
    StringBuilder sb = new StringBuilder();
    while (left <= right) {
      char c = s.charAt(left);

      if (c != ' ') sb.append(c);
      else if (sb.charAt(sb.length() - 1) != ' ') sb.append(c);

      ++left;
    }
    return sb;
  }

  public void reverse(StringBuilder sb, int left, int right) {
    while (left < right) {
      char tmp = sb.charAt(left);
      sb.setCharAt(left++, sb.charAt(right));
      sb.setCharAt(right--, tmp);
    }
  }

  public void reverseEachWord(StringBuilder sb) {
    int n = sb.length();
    int start = 0, end = 0;

    while (start < n) {
      // go to the end of the word
      while (end < n && sb.charAt(end) != ' ') ++end;
      // reverse the word
      reverse(sb, start, end - 1);
      // move to the next word
      start = end + 1;
      ++end;
    }
  }

  public String reverseWords(String s) {
    // converst string to string builder 
    // and trim spaces at the same time
    StringBuilder sb = trimSpaces(s);

    // reverse the whole string
    reverse(sb, 0, sb.length() - 1);

    // reverse each word
    reverseEachWord(sb);

    return sb.toString();
  }
}
```

### 扩展方法2-双端队列Deque[$^{[1]}$](#refer-anchor-1)

去除首尾空格，将每个单词压入队列，用join方法拼接队列。(关键词：双端队列、拼接字符串)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
  public String reverseWords(String s) {
    int left = 0, right = s.length() - 1;
    // remove leading spaces
    while (left <= right && s.charAt(left) == ' ') ++left;

    // remove trailing spaces
    while (left <= right && s.charAt(right) == ' ') --right;

    Deque<String> d = new ArrayDeque();
    StringBuilder word = new StringBuilder();
    // push word by word in front of deque
    while (left <= right) {
      char c = s.charAt(left);

      if ((word.length() != 0) && (c == ' ')) {
        d.offerFirst(word.toString());
        word.setLength(0);
      } else if (c != ' ') {
        word.append(c);
      }
      ++left;
    }
    d.offerFirst(word.toString());

    return String.join(" ", d);
  }
}
```

### 扩展方法3-调用java内置API实现[$^{[1,2]}$](#refer-anchor-1)

拆分字符串为单词数组后，将其转成集合，调用集合的逆置接口进行逆置，再调用String的join接口进行拼接。

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
public String reverseWords(String s) {
    String[] words = s.trim().split(" +");
    Collections.reverse(Arrays.asList(words));
    return String.join(" ", words);
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 151-Solution](https://leetcode.com/problems/reverse-words-in-a-string/solution/)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 151-Discuss](https://leetcode.com/problems/reverse-words-in-a-string/discuss/47781/Java-3-line-builtin-solution)
