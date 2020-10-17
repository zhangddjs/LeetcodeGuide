# [#133 Clone Graph](https://leetcode.com/problems/clone-graph)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、图、深拷贝、映射、DFS

## 题目

Given a reference of a node in a **connected** undirected graph.

Return a **deep copy** (clone) of the graph.

Each node in the graph contains a val (`int`) and a list (`List[Node]`) of its neighbors.

``` java
class Node {
    public int val;
    public List<Node> neighbors;
}
```

## 简述

**输入：** 无向连通图(某节点)

**输出：** 图的深拷贝(输入节点的拷贝)

**Notes：**

+ 1 <= 节点值 <= 100
+ 节点值唯一
+ 节点数量不超过100
+ 没有重复的边或自环
+ 节点的值代表自己在图中的顺序，比如1代表图中第1个节点，2代表第2个...

## 思路

本题考察数据结构的深拷贝操作。

本题和[[引用]138题](138-CopyListwithRandomPointer.md)类似，在138题中，使用了通过建立映射关系来完成拷贝的方法，该方法同样可以应用到本题的求解上。然后由于本题输入是无向图，因此可以从图中一节点通过DFS或BFS的方法遍历所有节点并进行深拷贝。------方法1、方法2

## 解决方案

### 方法1-DFS法

建立映射关系，DFS遍历拷贝每个节点。(关键词：映射、DFS)

时间复杂度：$O(n + e)$ ---96%

空间复杂度：$O(n)$ ---12%

``` java
/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;

    public Node() {
        val = 0;
        neighbors = new ArrayList<Node>();
    }

    public Node(int _val) {
        val = _val;
        neighbors = new ArrayList<Node>();
    }

    public Node(int _val, ArrayList<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
}
*/

class Solution {
    public Node cloneGraph(Node node) {
        if (node == null) return null;
        Node res = new Node(node.val);
        Set<Node> visited = new HashSet<>();
        Map<Node, Node> map = new HashMap<>();
        Node cur = res;
        map.put(node, cur);
        dfs(cur, visited, map, node.neighbors);
        return res;
    }

    public void dfs(Node cur, Set<Node> visited, Map<Node, Node> map, List<Node> neighbors) {
        if (visited.contains(cur)) return;
        visited.add(cur);
        for (Node neighbor : neighbors) {
            if (map.get(neighbor) == null) map.put(neighbor, new Node(neighbor.val));
            cur.neighbors.add(map.get(neighbor));
            dfs(map.get(neighbor), visited, map, neighbor.neighbors);
        }
    }
}
```

### 方法2-BFS法

建立映射关系，BFS遍历拷贝每个节点。(关键词：映射、BFS)

时间复杂度：$O(n + e)$ ---96%

空间复杂度：$O(n)$ ---12%

``` java
class Solution {
    public Node cloneGraph(Node node) {
        if (node == null) return null;
        Node res = new Node(node.val);
        Set<Node> visited = new HashSet<>();
        Map<Node, Node> map = new HashMap<>();
        Queue<Node> queue = new LinkedList<>();
        queue.offer(node);
        map.put(node, res);
        visited.add(node);
        while (!queue.isEmpty()) {
            node = queue.poll();
            Node cur = map.get(node);
            for (Node neighbor : node.neighbors) {
                if (!visited.contains(neighbor)) {
                    queue.offer(neighbor);
                    visited.add(neighbor);
                }
                if (map.get(neighbor) == null) map.put(neighbor, new Node(neighbor.val));
                cur.neighbors.add(map.get(neighbor));
            }
        }

        return res;
    }
}
```

## 扩展

### 扩展方法-不用访问标记的DFS法[$^{[1]}$](#refer-anchor-1)

利用Map进行访问标记，每构造一个映射关系，就完成了其所有邻居节点的深拷贝，这样存进map里的节点就是已经深拷贝完成的节点了。

时间复杂度：$O(n + e)$

空间复杂度：$O(n)$

``` java
/**
 * Copyright: LeetCode(https://leetcode.com)
 * Author: LeetCode
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
    private HashMap <Node, Node> visited = new HashMap <> ();
    public Node cloneGraph(Node node) {
        if (node == null) {
            return node;
        }

        if (visited.containsKey(node)) {
            return visited.get(node);
        }

        Node cloneNode = new Node(node.val, new ArrayList());
        visited.put(node, cloneNode);

        for (Node neighbor: node.neighbors) {
            cloneNode.neighbors.add(cloneGraph(neighbor));
        }
        return cloneNode;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 133-Solution](https://leetcode.com/problems/clone-graph/solution/)
