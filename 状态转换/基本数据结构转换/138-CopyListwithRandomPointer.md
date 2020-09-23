# [#138 Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、链表、深拷贝、双指针、两次遍历、三次遍历、HashMap、合并链表、链表还原

## 题目

A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a **deep copy** of the list.

The Linked List is represented in the input/output as a list of n nodes. Each node is represented as a pair of `[val, random_index]` where:

`val`: an integer representing `Node.val`
`random_index`: the index of the node (range from `0` to `n-1`) where random pointer points to, or `null` if it does not point to any node.

## 简述

**输入：** 每个节点包含随机指针的单链表

**输出：** 单链表的深拷贝

**Notes：**

+ -10000 <= 节点值 <= 10000
+ 随机指针可以指向任何一个节点和空节点
+ 节点数不大于1000

## 思路

本题考察单链表的深拷贝。

我们尝试直接遍历原链表，对每个节点用三个new的节点对象存放其数据信息、next指针和random指针。此时会发现一个问题，就是在给random创建节点时，该创建节点并不是链表中的节点。因此我们不能直接new，而是需要在构造好的新链表寻找目标节点的位置，然后令random指向它即可。

根据以上思路我们可以得到一个暴力法，即通过一次遍历初始化好深拷贝的新链表的每个节点和每个节点的next指针，然后再通过双指针第二次遍历，对每个节点(原链表`n1`，深拷贝链表`n2`}，记录原链表的节点`n1`的random指针，然后再用双指针在两个链表中遍历，当原链表指针指向的是random指向的对象时，则深拷贝链表指针也到了对应的位置，此时令当前深拷贝链表的节点`n2`的random指针指向该位置对应的节点即可。------方法1

可以发现在暴力法中第二次遍历时的时间复杂度比较高，主要的时间花费在寻找原链表`n1`节点random指针指向节点在深拷贝链表中的对应位置，因此如果能够用缓存存储该映射关系，寻找时间也就可以降下来了。

既然是存储映射关系，那么我们可以首先想到用HashMap数据结构作为缓存。因此我们可以在构造深拷贝链表前，先遍历一遍原链表，对于每一个节点的random指针，以其指向节点`p`作为key，用`p`初始化的节点`q`作为value存入HashMap，此时映射关系便构建完成。

在后续双指针遍历构造深拷贝链表时，先在HashMap中寻找是否有原链表当前节点`n`的key，如果有，代表存在random指针指向当前节点，此时取value链接到深拷贝链表尾部，如果没有，说明不存在random指针指向该节点，则通过`n`的数据初始化一个节点`cur`链接到深拷贝链表尾部。对于`cur`节点的random指针，可以直接用原链表当前节点`n`的random指针指向的节点`p`从map中得到`q`，令`cur`的random指向`q`即可。(`cur.random = map.get(n.random)`}------方法2

不过本题映射关系还可以通过另一种方式来建立，由于题目没有要求不能改原链表，所以我们可以先遍历一遍构建一个简单深拷贝链表(random信息还没有拷贝}，然后在遍历时将原链表的每个节点插入到其在深拷贝链表中对应的拷贝节点前面，此时便建立了映射关系。(原链表节点下一个节点即该节点的深拷贝节点}

再用双指针遍历一遍深拷贝链表，每次指针走两步，一个指针`p`遍历原链表节点，另一个指针`q`遍历深链表节点。则`q.random=p.random。next`。

最后再遍历一遍链表将原链表节点删除即可得到目标链表。不过此时可以在删除的同时重构原链表，使得原链表在方法结束时未被改变。------方法3

## 解决方案

### 方法1-暴力法

遍历两次，第一次构造简单深拷贝链表，第二次构造random指针，全程双指针法。(关键词：双指针、两次遍历)

时间复杂度：$O(n^2)$ ---20%

空间复杂度：$O(1)$ ---89%

``` java
/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/

class Solution {
    public Node copyRandomList(Node head) {
        if (head == null) return null;
        Node head2 = new Node(head.val);
        Node  n1 = head.next, n2 = head2;
        while (n1 != null) {
            n2.next = new Node(n1.val);
            n1 = n1.next;
            n2 = n2.next;
        }
        n1 = head;
        n2 = head2;
        Node p1 = head, p2 = head2;
        while (n1 != null) {
            Node random = n1.random;
            while (p1 != random) {
                p1 = p1.next;
                p2 = p2.next;
            }
            n2.random = p2;
            n1 = n1.next;
            n2 = n2.next;
            p1 = head;
            p2 = head2;
        }
        return head2;
    }
}
```

### 方法2-映射法

遍历两次，第一次构造映射关系，第二次构造深拷贝链表。(关键词：双指针、两次遍历、HashMap)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---93%

``` java
class Solution {
    public Node copyRandomList(Node head) {
        Node dummy = new Node(0);
        Node  n1 = head, n2 = dummy;
        Map<Node, Node> map = new HashMap<>();
        while (n1 != null) {
            if (n1.random != null) map.put(n1.random, new Node(n1.random.val));
            n1 = n1.next;
        }
        n1 = head;
        while (n1 != null) {
            n2.next = map.getOrDefault(n1, new Node(n1.val));
            n2 = n2.next;
            n2.random = map.get(n1.random);
            n1 = n1.next;
        }
        return dummy.next;
    }
}
```

### 方法3-优化的映射法

遍历三次，第一次构造简单深拷贝链表，合并原链表构建映射关系，第二次构造random指针，第三次拆分原链表和深拷贝链表，还原原来链表。(关键词：双指针、三次遍历、链表合并、链表还原)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---57%   _不计结果空间_

``` java
class Solution {
    public Node copyRandomList(Node head) {
        if (head == null) return null;
        Node dummy = new Node(0);
        Node  n1 = head, n2 = dummy;
        while (n1 != null) {
            Node tmp = n1.next;
            n2.next = n1;
            n2.next.next = new Node(n1.val);
            n1 = tmp;
            n2 = n2.next.next;
        }
        n1 = head;
        n2 = n1.next;
        while (n1 != null) {
            n2.random = n1.random == null ? null : n1.random.next;
            n2 = n2.next == null ? dummy : n2.next.next;
            n1 = n1.next.next;
        }
        n1 = head;
        while (n2.next != null) {
            n2.next = n2.next.next;
            n2 = n2.next;
            n1.next = n2.next;
            n1 = n1.next;
        }
        return dummy.next;
    }
}
```

## 扩展

### 扩展方法1-递归法[$^{[1]}$](#refer-anchor-1)

将链表想象成图，然后用深度优先遍历来构造next指针和random指针，并通过HashMap保存遍历过的节点及其深拷贝节点的映射关系。(关键词：链表转化成图、深度优先、HashMap)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */

public class Solution {
  // HashMap which holds old nodes as keys and new nodes as its values.
  HashMap<Node, Node> visitedHash = new HashMap<Node, Node>();

  public Node copyRandomList(Node head) {

    if (head == null) {
      return null;
    }

    // If we have already processed the current node, then we simply return the cloned version of
    // it.
    if (this.visitedHash.containsKey(head)) {
      return this.visitedHash.get(head);
    }

    // Create a new node with the value same as old node. (i.e. copy the node)
    Node node = new Node(head.val, null, null);

    // Save this value in the hash map. This is needed since there might be
    // loops during traversal due to randomness of random pointers and this would help us avoid
    // them.
    this.visitedHash.put(head, node);

    // Recursively copy the remaining linked list starting once from the next pointer and then from
    // the random pointer.
    // Thus we have two independent recursive calls.
    // Finally we update the next and random pointers for the new node created.
    node.next = this.copyRandomList(head.next);
    node.random = this.copyRandomList(head.random);

    return node;
  }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 138-Solution](https://leetcode.com/problems/copy-list-with-random-pointer/solution/)
