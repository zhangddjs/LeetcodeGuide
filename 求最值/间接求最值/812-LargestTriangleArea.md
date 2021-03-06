# [#812 Largest Triangle Area](https://leetcode.com/problems/largest-triangle-area/)

![Easy](/figures/Easy.svg)

## 关键词

最大面积、暴力法、遍历、排序、哨兵、数学、凸包

## 题目

You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

## 简述

**输入：** 点的数组

**输出：** 任意三个点组成的三角形中最大三角形的面积

**Notes：**

+ 3 < 点的个数 < 50;
+ 点集无重复
+ 点的x、y坐标范围在[-50,50];
+ 精确到10^-6;

## 思路

这是一道间接求最大的问题，通过给定的数组，求出其某个组合的某个特征的最大值。考察了几何数学知识点-如何在平面点集中选择三个点以组成最大的三角形，如果能够推倒出来便可解决本题。

那么如何选择这三个点？首先我们需要知道的是，三角形面积可以怎么算。目前主要可以通过以下三种方式进行计算：

1.给定三个顶点：

`return |(p1.x - p0.x) * (p2.y - p0.y) - (p2.x - p0.x) * (p1.y - p0.y)| / 2` ------①

2.给定三条边长：

``` c
s = (a + b + c) / 2;
return sqrt(s * (s - a) * (s - b) * (s - c));
```

3.一个角两个邻边：

`return 0.5 * a * b * sin(C)`

当知道给定三个点就能得出组成的三角形面积后，接下来就需要寻找可以比较的特征。

我们发现最大三角形取的三个点满足①式最大，如果能够**遍历**所有三个点的组合情况，记录由公式①算出的结果，遍历结束后将结果进行降序**排序**，第一个便是最大三角形的面积，这也就是暴力解法，三重循环实现遍历，时间复杂度为$O(\binom{n}{3}\log(\binom{n}{3}))$≈$O(n^{3}\log(n^{3}))$。------方法1

可以注意到，当所有面积存到数组或集合时，由于只需要获取最大值，所以并不一定需要借助排序来获取，我们还可以通过哨兵的方法，在遍历所有情况时用一个哨兵来储存当前的最大值，遍历完后全局最大值就是哨兵存储的值。这样可一定程度上提高效率，降低复杂度。其时间复杂度为$O(\binom{n}{3})$≈$O(n^{3})$。------方法2

## 解决方案

### 方法1-暴力法

遍历所有三个点的组合情形对应的三角形面积并排序求其中最大值返回。(关键词：遍历、排序)

时间复杂度：$O(\binom{n}{3}\log(\binom{n}{3}))$≈$O(n^{3}\log(n^{3}))$   ---5%

空间复杂度：$O(\binom{n}{3})$≈$O(n^{3})$   ---5%

``` java
class Solution {
    public double largestTriangleArea(int[][] points) {
        List<Double> list = new ArrayList<>();
        for(int i = 0; i < points.length - 2; i++)
            for(int j = i; j < points.length - 1; j++)
                for(int k = j; k < points.length; k++)
                    list.add(triangleArea(points[i], points[j], points[k]));
        Collections.sort(list);
        return list.get(list.size() - 1);
    }

    public double triangleArea(int[] point0, int[] point1, int[] point2){
        return Math.abs((point1[1] - point0[1]) * (point2[0] - point0[0]) -
                        (point2[1] - point0[1]) * (point1[0] - point0[0])) / 2.0;
    }
}
```

### 方法2-优化的暴力法

遍历所有三个点的组合情形对应的三角形面积，并实时记录局部最大值。(关键词：遍历、哨兵)

时间复杂度：$O(\binom{n}{3})$≈$O(n^{3})$   ---40%

空间复杂度：$O(1)$   ---12%

``` java
class Solution {
    public double largestTriangleArea(int[][] points) {
        double max = Double.MIN_VALUE;
        for(int i = 0; i < points.length - 2; i++)
            for(int j = i; j < points.length - 1; j++)
                for(int k = j; k < points.length; k++)
                    max = Math.max(triangleArea(points[i], points[j], points[k]), max);
        return max;
    }

    public double triangleArea(int[] point0, int[] point1, int[] point2){
        return Math.abs((point1[1] - point0[1]) * (point2[0] - point0[0]) -
                        (point2[1] - point0[1]) * (point1[0] - point0[0])) / 2.0;
    }
}
```

## 扩展

有大佬用凸包的方法减少了计算量，[答案链接](https://leetcode.com/problems/largest-triangle-area/discuss/179046/C)

### 凸包算法[$^{[1]}$](#refer-anchor-1)

#### 1. 增量式算法

逐次将点加入，然后检查之前的点是否在新的凸包上。由于每次都要检查所有之前的点，时间复杂度为$O(n^{2})$。

#### 2. 包裹法（Jarvis步进法）

首先由一点必定在凸包的点开始，例如最左的一点$A_{1}$。然后选择$A_{2}$点使得所有点都在$A_{1}A_{2}$的右方，这步骤的时间复杂度是$O(n)$，要比较所有点以$A_{1}$为原点的极坐标角度。以$A_{2}$为原点，重复这个步骤，依次找到$A_{3},A_{4},...,A_{k},A_{1}$。这总共有$k$步。因此，时间复杂度为$O(kn)$。

#### 3. 葛立恒（Graham）扫描法

由最底的一点$A_{1}$开始（如果有多个这样的点，那么选择最左边的），计算它跟其他各点的连线和x轴正向的角度，按小至大将这些点排序，称它们的对应点为$A_{2},A_{3},...,A_{n}$。这里的时间复杂度可达$O(n\log {n})$。

考虑最小的角度对应的点$A_{3}$。若由$A_{2}$到$A_{3}$的路径相对$A_{1}$到$A_{2}$的路径是向右转的（可以想象一个人沿$A_{1}$走到$A_{2}$，他站在$A_{2}$时，是向哪边改变方向），表示$A_{3}$不可能是凸包上的一点，考虑下一点由$A_{2}$到$A_{4}$的路径；否则就考虑$A_{3}$到$A_{4}$的路径是否向右转……直到回到$A_{1}$。

这个算法的整体时间复杂度是$O(n\log {n})$，注意每点只会被考虑一次，而不像Jarvis步进法中会考虑多次。

这个算法由葛立恒在1972年发明。它的缺点是不能推广到二维以上的情况。

伪代码：

``` alg
# 当ccw函数的值为正的时候，三个点为“左转”（counter-clockwise turn），如果是负的，则是“右转”的，而如果
# 为0，则三点共线，因为ccw函数计算了由p1,p2,p3三个点围成的三角形的有向面积
function ccw(p1, p2, p3):
    return (p2.x - p1.x)*(p3.y - p1.y) - (p2.y - p1.y)*(p3.x - p1.x)

let N           = number of points
let points[N+1] = the array of points
swap points[1] with the point with the lowest y-coordinate
sort points by polar angle with points[1]
# points[0] 是结束循环的标记点
let points[0] = points[N]
# M 是围成凸包的点的个数
let M = 1
for i = 2 to N:
    # Find next valid point on convex hull.
    while ccw(points[M-1], points[M], points[i]) <= 0:
        if M > 1:
            M -= 1
        # All points are collinear
        else if i == N:
            break
        else
            i += 1
        # 更新M，并把points[i]放到正确的位置
    M += 1
    swap points[M] with points[i]
```

#### 4. 单调链

将点按x坐标的值排列，再按y坐标的值排列。

选择x坐标为最小值的点，在这些点中找出y坐标的值最大和y坐标的值最小的点。对于x坐标为最大值也是这样处理。将两组点中y坐标值较小的点连起。在这条线段下的点，找出它们之中y坐标值最大的点，又在它们之间找x坐标值再最小和最大的点……如此类推。

时间复杂度是$O(n\log {n})$。

#### 5. 分治法

将点集X分成两个不相交子集。求得两者的凸包后，计算这两个凸包的凸包，该凸包就是X的凸包。时间复杂度是$O(n\log {n})$。

#### 6. 快包法（Akl-Toussaint启发式）

选择最左、最右、最上、最下的点，它们必组成一个凸四边形（或三角形）。这个四边形内的点必定不在凸包上。然后将其余的点按最接近的边分成四部分，再进行快包法（QuickHull）。

## 参考

<div id="refer-anchor-1"></div>

- [1] [Wikipedia. 凸包](https://zh.wikipedia.org/wiki/%E5%87%B8%E5%8C%85#%E8%91%9B%E7%AB%8B%E6%81%92%EF%BC%88Graham%EF%BC%89%E6%89%AB%E6%8F%8F%E6%B3%95)
