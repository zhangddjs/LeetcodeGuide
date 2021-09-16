// Step1 : build adjacency list
// Step2 : DFS the adjacency list, compute the longest two child lane of each node that treated as the root
// ** according to the formula that the max Diameter of one node == max of the sum of the longest two child lane of each of its child node.
// and the longest child lane of current node can be used for its parent node's longest two child lane computation
// Step3: after each dfs of a node is done, we should compare to build the max Diameter.

type node struct {
    Data int
    Next *node
}

type topTwo []int
var temp struct{}

var diameter int = 0
var visited map[int]struct{} = make(map[int]struct{})
var diag []*node = make([]*node, 0)

func treeDiameter(edges [][]int) int {
    if len(edges) <= 1 {
        return len(edges)
    }
    for i := 0; i <= len(edges); i++ {
        diag = append(diag, &node{Data:i})
    }
    for _, edge := range edges {  //build the diagram
        v1 := &node{Data:edge[0], Next:diag[edge[1]].Next}
        v2 := &node{Data:edge[1], Next:diag[edge[0]].Next}
        diag[edge[0]].Next = v2
        diag[edge[1]].Next = v1
    }
    visited[0] = temp
    theTopTwo := dfs(diag[0])
    return max(diameter, theTopTwo[0] + theTopTwo[1])
}

func dfs(cur *node) topTwo {
    myTopTwo := make(topTwo, 2, 2)
    for cur.Next != nil {
        cur = cur.Next
        if _, exist := visited[cur.Data]; exist {
            continue
        }
        visited[cur.Data] = struct{}
        childTopTwo := dfs(diag[cur.Data])
        diameter = max(diameter, childTopTwo[0] + childTopTwo[1])
        if childTopTwo[0] + 1 > myTopTwo[0] {  //swap then replace (sort)
            myTopTwo[0], myTopTwo[1] = myTopTwo[1], myTopTwo[0]
            myTopTwo[0] = childTopTwo[0] + 1
        } else if childTopTwo[0] + 1 > myTopTwo[1] {
            myTopTwo[1] = childTopTwo[0] + 1
        }
    }
    return myTopTwo
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}