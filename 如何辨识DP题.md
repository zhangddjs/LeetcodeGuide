# 如何辨识DP题

DP是比较难且比较重要的算法，不仅需要掌握其原理，还需要掌握如何分析哪些题目可以用DP解。

对于一个典型的DP题，必然需要枚举所有可能，因此有了回溯法。

如果有大量重复计算，就改进成备忘录递归。

进一步把递归改成迭代形式，就是DP。

**核心就是枚举，回溯，备忘录，dp。**

状态转移方程取决于题目要求的目标输出。求最大值就取Max，求最小值就取Min，判断就取true or false，统计就+1或+0。

题目[712-MinimumASCIIDeleteSumforTwoStrings]有很详细的过程详解。

[712-MinimumASCIIDeleteSumforTwoStrings]:求最值/间接求最值/712-MinimumASCIIDeleteSumforTwoStrings.md

## 案例

### 寻找最长公共子串

``` java
class Solution {
    public String longestSubStr(String s) {
        if (s == null || s.length() < 2) return s;
        String s2 = new StringBuffer(s).reverse().toString();
        int [][] dp = new int [s.length()][s.length()];
        int maxLen = 1, p = 0;
        //init
        for (int i = 0; i < s.length(); ++i) {
            dp[0][i] = s.charAt(0) == s2.charAt(i) ? 1 : 0;
            dp[i][0] = s2.charAt(0) == s.charAt(i) ? 1 : 0;
        }
        //dp
        for (int i = 1; i < s.length(); ++i) {
            for (int j = 1; j < s2.length(); ++j) {
                dp[i][j] = s.charAt(i) == s2.charAt(j) ? dp[i - 1][j - 1] + 1 : 0;
                if (dp[i][j] > maxLen $$ s.charAt(i) == s.charAt()) {
                    maxLen = dp[i][j];
                    p = i;
                }
            }
        }
        //build res
        return s.substring(p - maxLen + 1, p + 1);
    }
}
```
