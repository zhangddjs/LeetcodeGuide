# [#49 Group Anagrams](https://leetcode.com/problems/group-anagrams/)

![Medium](/figures/Medium.svg)

## 关键词

分类、字符串、遍历、字符串排序、HashMap、字符串表示计数数组/Map

## 题目

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## 简述

**输入：** 字符串数组

**输出：** 分类后的集合

**Notes：**

+ 1 <= 字符串数组长度 <= 10$^4$
+ 0 <= 字符串长度 <= 100
+ 字符串由小写字母构成

## 思路

本题考察分类，分类题往往会用到HashSet、HashMap等数据结构，我们需要分析这道题是否也可以用到。

本题需要将同构的字符串归到同一类，可以得知，同构的字符串排序后字母顺序是相同的，因此将每个字符串按字典序对字符进行排序，插入到HashMap中，所有字符串遍历完后便可得到分类结果。------方法1

仔细观察可以发现，对于两个字符串，如果是同构，那么它们每个字符的数量是相等的。因此只需要对每个字符串进行扫描和字符数量统计，插入到HashMap对应的位置，就可以得到结果。

此方法的难点在于如何选定数据结构作为key来存储字符数量信息。有一个很好的方法是通过字符串，来记录，我们知道小写字母一共有26个，因此字符串中可以用26个通过`'#'`拼接的数字来分别记录每个字母的出现次数。------方法2[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-暴力法(排序分类)

遍历字符串数组并排序每个字符串，插入到HashMap，返回HashMap值的集合。(关键词：遍历、字符串排序、HashMap)

时间复杂度：$O(mn\log(m))$ ---50%

空间复杂度：$O(mn)$ ---88%

``` java
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String str : strs) {
            char[] arr = str.toCharArray();
            Arrays.sort(arr);
            String tmp = new String(arr);
            if (map.get(tmp) == null) map.put(tmp, new ArrayList<String>());
            map.get(tmp).add(str);
        }
        return map.values().stream().collect(Collectors.toList());      //return new ArrayList(map.values());
    }
}
```

### 方法2-计数分类[$^{[1]}$](#refer-anchor-1)

统计每个字符串每个字符出现的次数，插入到HashMap，返回HashMap值的集合。(关键词：计数、字符串表示计数数组/Map、HashMap)

时间复杂度：$O(mn)$

空间复杂度：$O(mn)$

``` java
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        if (strs.length == 0) return new ArrayList();
        Map<String, List> ans = new HashMap<String, List>();
        int[] count = new int[26];
        for (String s : strs) {
            Arrays.fill(count, 0);
            for (char c : s.toCharArray()) count[c - 'a']++;

            StringBuilder sb = new StringBuilder("");
            for (int i = 0; i < 26; i++) {
                sb.append('#');
                sb.append(count[i]);
            }
            String key = sb.toString();
            if (!ans.containsKey(key)) ans.put(key, new ArrayList());
            ans.get(key).add(s);
        }
        return new ArrayList(ans.values());
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 49-Solution](https://leetcode.com/problems/group-anagrams/solution/)
