# [#17 Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number)

![Medium](/figures/Medium.svg)

## 关键词

枚举、全排列、回溯法、DFS、记录路径、BFS

## 题目

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in **any order**.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

![image](https://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png)

## 简述

**输入：** 手机键盘按键序列

**输出：** 所有可能的字母组合(无视顺序)

**Notes：**

+ 0 <= 序列长度 <= 4
+ 按键范围`['2', '9']`

## 思路

本题考察枚举，通常需要运用回溯法思想来解决此类问题。

对于本题，可以联想成图或者n叉树，然后使用深度遍历方法，对于序列中某个位置键位，将其对应的所有可能的字母和之前的遍历路径进行组合，然后对于每个组合继续遍历序列中下一个键位。------方法1

## 解决方案

### 方法1-回溯法

深度优先遍历，记录遍历路径，到最深层时加入结果列表。(关键词：DFS、记录路径)

时间复杂度：$O(3^n*4^m)$ ---38%

空间复杂度：$O(3^n*4^m)$ ---40%

``` java
class Solution {
    Map<Character, String> map = new HashMap<>();
    public List<String> letterCombinations(String digits) {
        init();
        List<String> res = new ArrayList<>();
        helper(digits, "", res);
        return res;
    }

    public void init() {
        map.put('2', "abc");
        map.put('3', "def");
        map.put('4', "ghi");
        map.put('5', "jkl");
        map.put('6', "mno");
        map.put('7', "pqrs");
        map.put('8', "tuv");
        map.put('9', "wxyz");
    }

    public void helper(String digits, String path, List<String> res) {
        if (digits == null || digits.equals("")) {
            if (!path.equals("")) res.add(path);
            return;
        }
        String nextDigits = digits.substring(1);
        char cur = digits.charAt(0);
        helper(nextDigits, path + map.get(cur).charAt(0), res);
        helper(nextDigits, path + map.get(cur).charAt(1), res);
        helper(nextDigits, path + map.get(cur).charAt(2), res);
        if (cur == '7' || cur == '9')
            helper(nextDigits, path + map.get(cur).charAt(3), res);
    }
}
```

## 扩展

### 扩展方法-BFS法[$^{[1]}$](#refer-anchor-1)

本题在使用DFS的同时也可以用BFS的方法来实现。需要借助一个队列，时间和空间复杂度和DFS法相同。

``` java
/**
 * Copyright: LeetCode(https://leetcode.com)
 * Author: lirensun
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
public List<String> letterCombinations(String digits) {
    LinkedList<String> ans = new LinkedList<String>();
    if(digits.isEmpty()) return ans;
    String[] mapping = new String[] {"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
    ans.add("");
    for(int i =0; i<digits.length();i++){
        int x = Character.getNumericValue(digits.charAt(i));
        while(ans.peek().length()==i){
            String t = ans.remove();
            for(char s : mapping[x].toCharArray())
                ans.add(t+s);
        }
    }
    return ans;
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 17-Discuss](https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/8064/My-java-solution-with-FIFO-queue)
