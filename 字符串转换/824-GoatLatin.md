# 824题-GoatLatin

![简单题](https://img.shields.io/github/release/ustctug/ustcthesis/all.svg)

## 关键词

拆分字符串、拼接字符串、遍历、字符处理、字符串截取、单词缓存、HashSet

## 描述

**输入：** 字符串

**输出：** 根据给定case转化后的目标字符串

**Notes：**

- 1 <= 字符串长度 <= 150.
- 输入字符串包含大小写字母和空格，每个单词之间1个空格。

## 思路

本题考察字符串操作，按照题目所提供的三个case进行合适的处理即可。

从case中可以看出，本题要对每个单词进行判断和处理，因此第一步是拆分字符串，拆成数组后进行遍历，根据case删除单词首字母或在词尾添加字符，拼接到输出字符串末尾，这一步也可以在遍历完成后将所有单词拼接并通过空格符分隔。------方法1

参考[[引用]819题](..\求最值\间接求最值\819题-最频繁单词.md)可知，通过字符处理+单词缓存的方法可以以OnePass的方式完成整个操作流程。此处不再赘述。

## 解决方案

### 方法1-暴力解

拆分字符串，遍历单词数组，按case处理，拼接到输出。(关键词：拆分，遍历，字符处理，字符串截取)

时间复杂度：$O(n)$ ---91%

空间复杂度：$O(k)$ ---80%

``` java
//遍历时拼接
class Solution {
    public String toGoatLatin(String S) {
        String [] words = S.split(" ");
        Set<Character> vowels = new HashSet<>();
        Collections.addAll(vowels, new Character[]{'a','e','i','o','u', 'A', 'E', 'I', 'O', 'U'});
        StringBuilder suffix = new StringBuilder("ma");
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < words.length; ++i) {
            suffix.append("a");
            StringBuilder wordBuilder = new StringBuilder(words[i]);
            if (!vowels.contains(words[i].charAt(0))) {
                wordBuilder.deleteCharAt(0);
                wordBuilder.append(words[i].charAt(0));
            }
            wordBuilder.append(suffix);
            res.append(wordBuilder);
            if(i != words.length - 1) res.append(" ");
        }
        return res.toString();
    }
}
```

``` java
//遍历后拼接
class Solution {
    public String toGoatLatin(String S) {
        String [] words = S.split(" ");
        Set<Character> vowels = new HashSet<>();
        Collections.addAll(vowels, new Character[]{'a','e','i','o','u', 'A', 'E', 'I', 'O', 'U'});
        String suffix = "ma";
        String res = "";
        for (int i = 0; i < words.length; ++i) {
            suffix += "a";
            if (!vowels.contains(words[i].charAt(0)))
                words[i] = words[i].substring(1, words[i].length()) + words[i].charAt(0);
            words[i] += suffix;
        }
        for (String word : words)
            res += word + " ";
        return res.substring(0, res.length() - 1);
    }
}
```

### 方法2-字符处理

遍历字符串，用缓存记录当前遍历的单词的每个字符，遍历到空格时说明一个单词中的字符已全部缓存，进行操作拼接到输出。(关键词：遍历，字符处理，单词缓存)

时间复杂度：$O(n)$   ---91%

空间复杂度：$O(1)$   ---98%

``` java
class Solution {
    public String toGoatLatin(String S) {
        Set<Character> vowels = new HashSet<>();
        Collections.addAll(vowels, new Character[]{'a','e','i','o','u', 'A', 'E', 'I', 'O', 'U'});
        StringBuilder suffix = new StringBuilder("ma");
        StringBuilder res = new StringBuilder();
        StringBuilder word = new StringBuilder();
        S += " ";
        for (char c : S.toCharArray()) {
            if (Character.isLetter(c)){
                word.append(c);
                continue;
            }
            suffix.append('a');
            if (!vowels.contains(word.charAt(0)))
                word.append(word.charAt(0)).deleteCharAt(0);
            res.append(word).append(suffix).append(" ");
            word = new StringBuilder();
        }
        return res.deleteCharAt(res.length() - 1).toString();
    }
}
```
