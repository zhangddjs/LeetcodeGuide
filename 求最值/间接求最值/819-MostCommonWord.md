# [#819 Most Common Word](https://leetcode.com/problems/most-common-word/)

![Easy](/figures/Easy.svg)

## 关键词

最频繁单词、转换成求最值问题、遍历、分类、排序、哨兵、正则、HashMap、HashSet、拆分字符串

## 题目

Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.

## 简述

**输入：** 英文段落；单词黑名单

**输出：** 在英文段落中黑名单之外的最频繁出现的单词

**Notes：**

+ 1 <= 段落长度 <= 1000.
+ 0 <= 黑名单单词数 <= 100.
+ 1 <= 每个黑名单单词长度 <= 10.
+ 答案唯一，小写输出。
+ 段落只包含字母、空格或标点符号感叹号、问号、单引号、逗号、分号、句号。
+ 没有连字符或带有连字符的单词。
+ 单词仅包含字母，不能包含单引号或其他标点符号。

## 思路

这道题考察从不同特征的元素集中寻找数量最多的拥有相同特征的那类元素中的任意一个元素。

我们知道，要得到最频繁出现的元素，往往都需要至少遍历一遍整个元素集。对于这道题，看上去输入是一个字符串整体，但根据英文段落的性质，可以将这个字符串按正则表达式"\W?\s\W?|\W\$|\W"进行拆分，拆分成单词数组，随后对这个单词数组进行遍历，并用HashMap记录黑名单外且不含特殊字符的每个单词的出现次数，遍历结束后对HashMap按value进行排序，便可取得最频繁单词。时间复杂度为$O(n\log(n))$。------方法1

可以注意到，通过方法1的思路，本题已经转换成了求最值的问题，即从众多频数中求最大的频数。根据本题的要求，我们可以使用哨兵的方式来取代排序步骤，在遍历时用两个哨兵来分别记录当前的最大频数和其对应的单词。时间复杂度为$O(n)$。------方法2

## 解决方案

### 方法1-暴力解

将段落拆分成单词数组，遍历单词并记录每个符合条件的单词的出现次数，排序求出现次数最多的单词。(关键词：遍历、排序)

时间复杂度：$O(n\log(n))$   ---66%

空间复杂度：$O(n)$   ---47%

``` java
class Solution {
    public String mostCommonWord(String paragraph, String[] banned) {
        paragraph = paragraph.toLowerCase();
        String[] words = paragraph.split("\\W?\\s\\W?|\\W$|\\W");
        Map<String, Integer> map = new HashMap<>();
        Set<String> bannedSet = new HashSet<>(Arrays.asList(banned));
        for(String word : words)
            if(!bannedSet.contains(word))
                map.put(word, map.getOrDefault(word, 0) + 1);
        List<Map.Entry<String, Integer>> list = new ArrayList<>();
        list.addAll(map.entrySet());
        Collections.sort(list, (Map.Entry<String, Integer> o1, Map.Entry<String, Integer> o2) -> o2.getValue() - o1.getValue());
        return list.get(0).getKey();
    }
}
```

### 方法2-优化的暴力法(字符串处理流水线String Processing in Pipeline)

将段落拆分成单词数组，遍历单词并记录每个符合条件的单词的出现次数，实时记录局部最大频数和其对应的单词。(关键词：遍历、哨兵)

时间复杂度：$O(n)$   ---66%

空间复杂度：$O(n)$   ---68%

``` java
class Solution {
    public String mostCommonWord(String paragraph, String[] banned) {
        paragraph = paragraph.toLowerCase();
        String[] words = paragraph.split("\\W?\\s\\W?|\\W$|\\W");
        Map<String, Integer> map = new HashMap<>();
        Set<String> bannedSet = new HashSet<>(Arrays.asList(banned));
        String res = "";
        int max = 0;
        for(String word : words){
            if(!bannedSet.contains(word)){
                map.put(word, map.getOrDefault(word, 0) + 1);
                if(map.get(word) > max){
                    max = map.get(word);
                    res = word;
                }
            }
        }
        return res;
    }
}
```

## 扩展

### 扩展一些字符串正则处理方式[$^{[1]}$](#refer-anchor-1)

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
String normalizedStr = paragraph.replaceAll("[^a-zA-Z0-9 ]", " ").toLowerCase();
String[] words = normalizedStr.split("\\s+");

String[] words = p.replaceAll("\\W+" , " ").toLowerCase().split("\\s+");
```

### 扩展方法3-真正的OnePass:字符处理(Character Processing in One-Pass)[$^{[1]}$](#refer-anchor-1)

遍历每个字符，用Character.isLetter(currChar)判断是否到达单词末尾和下一个单词的开头，从而得到当前正在遍历的单词，更新频数。(关键词：字符处理、单词缓存)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
    public String mostCommonWord(String paragraph, String[] banned) {

        Set<String> bannedWords = new HashSet();
        for (String word : banned)
            bannedWords.add(word);

        String ans = "";
        int maxCount = 0;
        Map<String, Integer> wordCount = new HashMap();
        StringBuilder wordBuffer = new StringBuilder();
        char[] chars = paragraph.toCharArray();

        for (int p = 0; p < chars.length; ++p) {
            char currChar = chars[p];

            // 1). consume the characters in a word
            if (Character.isLetter(currChar)) {
                wordBuffer.append(Character.toLowerCase(currChar));
                if (p != chars.length - 1)
                    // skip the rest of the processing
                    continue;
            }

            // 2). at the end of one word or at the end of paragraph
            if (wordBuffer.length() > 0) {
                String word = wordBuffer.toString();
                // identify the maximum count while updating the wordCount table.
                if (!bannedWords.contains(word)) {
                    int newCount = wordCount.getOrDefault(word, 0) + 1;
                    wordCount.put(word, newCount);
                    if (newCount > maxCount) {
                        ans = word;
                        maxCount = newCount;
                    }
                }
                // reset the buffer for the next word
                wordBuffer = new StringBuilder();
            }
        }
        return ans;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 819-Solution](https://leetcode.com/problems/most-common-word/solution/)
