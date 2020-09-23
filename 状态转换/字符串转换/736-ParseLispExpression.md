# [#736 Parse Lisp Expression](https://leetcode.com/problems/parse-lisp-expression/)

![Hard](/figures/Hard.svg)

## 关键词

字符串分析、字符串转换、表达式、拆分字符串、遍历字符串、分情况处理、栈、HashMap、递归

## 题目

You are given a string `expression` representing a Lisp-like expression to return the integer value of.

The syntax for these expressions is given as follows.

+ An expression is either an integer, a let-expression, an add-expression, a mult-expression, or an assigned variable. Expressions always evaluate to a single integer.

+ (An integer could be positive or negative.)

+ A let-expression takes the form `(let v1 e1 v2 e2 ... vn en expr)`, where `let` is always the string `"let"`, then there are 1 or more pairs of alternating variables and expressions, meaning that the first variable `v1` is assigned the value of the expression `e1`, the second variable `v2` is assigned the value of the expression `e2`, and so on **sequentially**; and then the value of this let-expression is the value of the expression `expr`.

+ An add-expression takes the form `(add e1 e2)` where `add` is always the string `"add"`, there are always two expressions `e1, e2`, and this expression evaluates to the addition of the evaluation of `e1` and the evaluation of `e2`.

+ A mult-expression takes the form `(mult e1 e2)` where `mult` is always the string `"mult"`, there are always two expressions `e1, e2`, and this expression evaluates to the multiplication of the evaluation of `e1` and the evaluation of `e2`.

+ For the purposes of this question, we will use a smaller subset of variable names. A variable starts with a lowercase letter, then zero or more lowercase letters or digits. Additionally for your convenience, the names "add", "let", or "mult" are protected and will never be used as variable names.

+ Finally, there is the concept of scope. When an expression of a variable name is evaluated, **within the context of that evaluation**, the innermost scope (in terms of parentheses) is checked first for the value of that variable, and then outer scopes are checked sequentially. It is guaranteed that every expression is legal. Please see the examples for more details on scope.

Example:

+ (add 1 2) -> 3

+ (mult 3 (add 2 3)) -> 15

+ (let x 2 (mult x 5)) -> 10

+ (let x 2 (mult x (let x 3 y 4 (add x y)))) -> 14

+ (let x 3 x 2 x) -> 2

+ (let x 1 y 2 x (add x y) (add x y)) -> 5

+ (let x 2 (add (let x 3 (let x 4 x)) x)) -> 6

+ (let a1 3 b2 (add a1 1) b2) -> 4

## 简述

**输入：** 表达式字符串

**输出：** 表达式的解

**Notes：**

+ 表达式格合法，单个空格分隔
+ 表达式非空且长度最大是2000
+ 所有计算结果都在32位整数内

## 思路

本题考察字符串转换和数据结构设计，懂得编译原理可能会很有帮助。

首先我们观察输入表达式字符串的特征：

1. 它是一个满足题目条件的表达式。
2. 单个空格分隔每个元素。
3. 内层表达式由括号包围，最外层表达式两端也是括号包围。
4. 表达式开头必然是`(let`、`(add`、`(muti`。
5. 表达式结尾是`elm)`、`elm)))...`
6. `let`操作是线性赋值的。
7. 表达式计算从内层向外层。
8. 每个表达式都有返回值。
9. 除了`let`表达式包含返回值元素外表达式内元素是成对的。
10. 只有`let`可能有多对参数

对于这些特征，我们可以方便地想到将字符串用空格分割成字符数组(也可以用字符操作和单词缓存的思想来操作)，然后遍历数组，对每个元素进行分析并进行对应处理，根据线性赋值和嵌套计算的特性，可以想到借助队列和栈这两个数据结构来存储相关信息。

我们可以用栈来存储一个表达式的返回值和它的每个元素的赋值，赋值信息可以用Map来记录，当进行下一层运算时，上一层入栈，生成上一层map的拷贝，赋值操作不会影响上一层，当返回上一层时，则出栈。------方法1

## 解决方案

### 方法1-暴力法

拆分表达式，对每个元素进行判断与处理，利用栈和HashMap协助存储变量信息、操作符和返回结果等相关信息。(关键词：拆分字符串、遍历字符串、分情况处理、栈、HashMap)

时间复杂度：$O(N^2)$ ---48%

空间复杂度：$O(N^2)$ ---86%

``` java
class Solution {
    public int evaluate(String expression) {
        String [] elements = expression.split(" ");
        Stack<Map<String, String>> stack = new Stack<>();
        Stack<String> keys = new Stack<>();
        Stack<String> optkeys = new Stack<>();
        Map<String, String> curMap = new HashMap<>();
        int res = 0;
        for (int i = 0; i < elements.length; ++i) {
            String element = elements[i];
            if (element.startsWith("(")) {
                if (element.equals("(let")) {
                    stack.push(curMap);
                    curMap = new HashMap<String, String>(curMap);
                }
                optkeys.push(element);
                keys.push(element);
            } else if (element.endsWith(")")) {
                StringBuilder valuebuf = new StringBuilder();
                int j;
                for (j = 0; element.charAt(j) != ')'; ++j) valuebuf.append(element.charAt(j));
                String value = curMap.getOrDefault(valuebuf.toString(), valuebuf.toString());
                for (j = j; j < element.length(); ++j) {
                    String opt = optkeys.pop();
                    String subres = keys.pop();
                    subres = curMap.getOrDefault(subres, subres);
                    if (opt.equals("(add")){
                        subres = Integer.valueOf(subres) + Integer.valueOf(value) + "";
                        keys.pop();
                    }
                    else if (opt.equals("(mult")){
                        subres = Integer.valueOf(subres) * Integer.valueOf(value) + "";
                        keys.pop();
                    }
                    else {
                        subres = value.toString();
                        curMap = stack.pop();
                    }
                    value = subres;
                }
                if(keys.isEmpty()) {
                    keys.push(value);
                    break;
                }
                String key = keys.peek();
                if (key.equals(optkeys.peek())) keys.push(value);
                else {
                    key = keys.pop();
                    curMap.put(key, curMap.getOrDefault(value, value));
                }
            } else {
                String key = keys.peek();
                if (key.equals(optkeys.peek())) keys.push(element);
                else {
                    key = keys.pop();
                    curMap.put(key, curMap.getOrDefault(element, element));
                }
            }
        }
        return Integer.valueOf(keys.pop());
    }
}
```

## 扩展

### 扩展方法-递归分析[$^{[1]}$](#refer-anchor-1)

将迭代法进行优化，拆分子问题，转化成更简洁的递归形式。

时间复杂度：$O(N^2)$

空间复杂度：$O(N^2)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
    ArrayList<Map<String, Integer>> scope;
    public Solution() {
        scope = new ArrayList();
        scope.add(new HashMap());
    }

    public int evaluate(String expression) {
        scope.add(new HashMap());
        int ans = evaluate_inner(expression);
        scope.remove(scope.size() - 1);
        return ans;
    }

    public int evaluate_inner(String expression) {
        if (expression.charAt(0) != '(') {
            if (Character.isDigit(expression.charAt(0)) || expression.charAt(0) == '-')
                return Integer.parseInt(expression);
            for (int i = scope.size() - 1; i >= 0; --i) {
                if (scope.get(i).containsKey(expression))
                    return scope.get(i).get(expression);
            }
        }

        List<String> tokens = parse(expression.substring(
                expression.charAt(1) == 'm' ? 6 : 5, expression.length() - 1));
        if (expression.startsWith("add", 1)) {
            return evaluate(tokens.get(0)) + evaluate(tokens.get(1));
        } else if (expression.startsWith("mult", 1)) {
            return evaluate(tokens.get(0)) * evaluate(tokens.get(1));
        } else {
            for (int j = 1; j < tokens.size(); j += 2) {
                scope.get(scope.size() - 1).put(tokens.get(j-1), evaluate(tokens.get(j)));
            }
            return evaluate(tokens.get(tokens.size() - 1));
        }
    }

    public List<String> parse(String expression) {
        List<String> ans = new ArrayList();
        int bal = 0;
        StringBuilder buf = new StringBuilder();
        for (String token: expression.split(" ")) {
            for (char c: token.toCharArray()) {
                if (c == '(') bal++;
                if (c == ')') bal--;
            }
            if (buf.length() > 0) buf.append(" ");
            buf.append(token);
            if (bal == 0) {
                ans.add(new String(buf));
                buf = new StringBuilder();
            }
        }
        if (buf.length() > 0)
            ans.add(new String(buf));

        return ans;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 736-Solution](https://leetcode.com/problems/parse-lisp-expression/solution/)
