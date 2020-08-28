# [#836 Rectangle Overlap]([leetcodelink](https://leetcode.com/problems/rectangle-overlap/))

![Easy](/figures/Easy.svg)

## 关键词

求交集、寻找相交条件、遍历相交条件、逆转思维、几何数学、不相交条件和相交条件互推、条件优化、相交条件判别法不唯一(Check Position、Check Area)

## 题目

A rectangle is represented as a list `[x1, y1, x2, y2]`, where `(x1, y1)` are the coordinates of its bottom-left corner, and `(x2, y2)` are the coordinates of its top-right corner.

Two rectangles overlap if the area of their intersection is positive.  To be clear, two rectangles that only touch at the corner or edges do not overlap.

Given two (axis-aligned) rectangles, return whether they overlap.

## 简述

**输入：** 两个矩阵

**输出：** 是否重叠

**Notes：**

+ 两个矩阵信息由左下角和由上角的点坐标表示
+ 点坐标为整形，在$-10^9$~$10^9$之间
+ 边界相交不算

## 思路

本题考察如何判断两组元素是否有交集，需要根据题目信息确定相交或不相交的判断条件，只要有任意元素满足相交条件，则相交。

这道题的输入实际上是两组由无穷多个的点元素构成的两个矩形面，如果一个矩形范围内的某点位置在另一个矩形范围内，则两矩形相交。本题中虽然有多种相交情况，但其核心相交条件始终是：矩形1左上角坐标在矩形2右下角的左上，不高于矩形2左上角，矩形1右上角在矩形2左下角的右上。做完判断后，如果不符合则对调矩形1和矩形2再进行判断，可得到是否相交的结果。------方法1

相交情况较多时，遍历所有条件可能实现起来比较复杂，如果将思维逆转，遍历判断所有不相交的情况，如果都不满足，那么就是相交，这样或许能够减少很多工作量。------方法2

## 解决方案

### 方法1-暴力法

遍历判断所有相交条件，返回结果。(关键词：寻找相交条件，遍历相交条件，几何数学)

时间复杂度：$O(1)$ ---100%

空间复杂度：$O(1)$ ---75%

``` java
class Solution {
    class Point{
        int x;
        int y;
    }
    class Rect{
        Point bottomLeft = new Point();
        Point bottomRight = new Point();
        Point topLeft = new Point();
        Point topRight = new Point();
        Rect(int[] rec){
            bottomLeft.x = rec[0];
            bottomLeft.y = rec[1];
            bottomRight.x = rec[2];
            bottomRight.y = rec[1];
            topLeft.x = rec[0];
            topLeft.y = rec[3];
            topRight.x = rec[2];
            topRight.y = rec[3];
        }
    }
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        Rect rect1 = new Rect(rec1);
        Rect rect2 = new Rect(rec2);
        if((rect1.topLeft.x < rect2.bottomRight.x &&
            rect1.topLeft.y > rect2.bottomRight.y) &&
           rect1.topLeft.y <= rect2.topLeft.y &&
           (rect1.topRight.x > rect2.bottomLeft.x &&
            rect1.topRight.y > rect2.bottomLeft.y)) return true;
        if((rect2.topLeft.x < rect1.bottomRight.x &&
            rect2.topLeft.y > rect1.bottomRight.y) &&
           rect2.topLeft.y <= rect1.topLeft.y &&
           (rect2.topRight.x > rect1.bottomLeft.x &&
            rect2.topRight.y > rect1.bottomLeft.y)) return true;
        return false;
    }
}
```

### 方法2-逆转思维法

遍历判断所有不相交条件，返回结果。(关键词：逆转思维，遍历不相交条件，几何数学)

时间复杂度：$O(1)$ ---100%

空间复杂度：$O(1)$ ---42%

``` java
class Solution {
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        if ((rec1[0] >= rec2[2]) ||
            (rec1[1] >= rec2[3]) ||
            (rec2[0] >= rec1[2]) ||
            (rec2[1] >= rec1[3])) return false;
        return true;
    }
}
```

## 扩展

### 扩展方法1-方法1和方法2的OneLine优化(Check Position)[$^{[1]}$](#refer-anchor-1)

上面讲述的两种方法其实还有继续化简的可能性。当得到方法2时，通过不相交条件逆推相交条件的方法，便可以得到方法1的优化版本，同时也是方法2的优化版，一行搞定。(关键词：不相交条件和相交条件互推、条件优化)

``` java
class Solution {
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        return rec1[0] < rec2[2] && rec2[0] < rec1[2] && rec1[1] < rec2[3] && rec2[1] < rec1[3];
    }
}
```

### 扩展方法2-另一种相交条件判别法(Check Area)[$^{[2,3]}$](#refer-anchor-2)

将输入的点转化成矩形的四条边`l,b,r,t`(for left, bottom, right, and top)，从平行边中分别取两个矩形中最左边的r和最右边的l以及最下边的t和最上边的b，得到r与l的距离以及t和b的距离，这个距离代表了水平方向或者垂直方向发生重叠时重叠的长度大小，当两个距离为正时，表示重叠，如果其中一个为负，则表示水平或垂直方向没有重叠，两个矩形也就不会重叠。(关键词：相交条件判别法不唯一、几何数学)

核心算法如下：

``` java
l1,b1,r1,t1 = rec1
l2,b2,r2,t2 = rec2
width = min(r1,r2) - max(l1,l2)
height = min(t1,t2) - max(b1,b2)
return width > 0 and height > 0
```

简化后为：

``` java
class Solution {
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        return (Math.min(rec1[2], rec2[2]) > Math.max(rec1[0], rec2[0]) && // width > 0
                Math.min(rec1[3], rec2[3]) > Math.max(rec1[1], rec2[1]));  // height > 0
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 836-Discuss](https://leetcode.com/problems/rectangle-overlap/discuss/132340/C++JavaPython-1-line-Solution-1D-to-2D)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 836-Comment](https://leetcode.com/problems/rectangle-overlap/discuss/132340/C++JavaPython-1-line-Solution-1D-to-2D/140104)

<div id="refer-anchor-3"></div>

+ [3] [Leetcode. 836-Solution](https://leetcode.com/problems/rectangle-overlap/solution/)
