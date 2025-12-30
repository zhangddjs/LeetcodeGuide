package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &Heap{make([]*ListNode, 0, len(lists))}
	for _, l := range lists {
		h.Push(l)
	}
	res := &ListNode{0, nil}
	tail := res
	node := h.Pop()
	for node != nil {
		tail.Next = node
		tail = tail.Next
		h.Push(node.Next)
		node = h.Pop()
	}
	return res.Next
}

type Heap struct {
	Arr []*ListNode
}

func (h *Heap) swiftUp() {
	cur := len(h.Arr) - 1
	parent := (cur - 1) / 2
	for cur > 0 {
		if h.Arr[parent].Val > h.Arr[cur].Val {
			h.Arr[parent], h.Arr[cur] = h.Arr[cur], h.Arr[parent]
		}
		cur, parent = parent, (parent-1)/2
	}
}

func (h *Heap) swiftDown() {
	cur := 0
	l, r := 2*cur+1, 2*cur+2
	for cur < len(h.Arr) {
		minnode := h.Arr[cur].Val
		minIdx := cur
		if l < len(h.Arr) && h.Arr[l].Val < minnode {
			minnode = h.Arr[l].Val
			minIdx = l
		}
		if r < len(h.Arr) && h.Arr[r].Val < minnode {
			minnode = h.Arr[r].Val
			minIdx = r
		}
		if minIdx == cur {
			break
		}
		h.Arr[cur], h.Arr[minIdx] = h.Arr[minIdx], h.Arr[cur]
		cur = minIdx
		l, r = 2*cur+1, 2*cur+2
	}
}

func (h *Heap) Push(node *ListNode) {
	if node == nil {
		return
	}
	h.Arr = append(h.Arr, node)
	h.swiftUp()
}

func (h *Heap) Pop() *ListNode {
	if len(h.Arr) == 0 {
		return nil
	}
	n := h.Arr[0]
	h.Arr[0] = h.Arr[len(h.Arr)-1]
	h.Arr = h.Arr[:len(h.Arr)-1]
	h.swiftDown()
	return n
}
