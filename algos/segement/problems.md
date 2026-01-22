# 线段树

solve problems manipulating groups of polygons on a grid.

- [x] 307. Range Sum Query - Mutable (Medium) (Fenwick Tree able)
- [x] 218. The Skyline Problem (Hard) (线段树会比较麻烦)
- [ ] 315. Count of Smaller Numbers After Self (Hard)
- [ ] 2276. Count Integers in Intervals (Advanced Hard)
- [ ] 699. Falling Squares (Advanced Hard)

好题推荐 默写用 ⭐️⭐️⭐️

- [x] 3453 Separate Squares I
- [ ] 3454 Separate Squares II

## Sweep Line

- [ ] 986 interval List Intersections
- [ ] 850 Rectangle Area II
- [ ] 1851 Minimum Interval to Include Each Query

intervals: `[1,2]`,`[3,4]`,`[2,5]`,...

### Step1 (Recording/Collect): x-Axis to event, entrance +1, leave -1

`[1,2]` -> `[1,1],[2,-1]`
`[3,4]` -> `[3,1],[4,-1]`
...

### Step2 (Sorting): sort events by +1 or -1 by situation

O(nlogn) n = 2m of intervals num

当多个事件在同一个坐标时：
- 需要先处理哪种事件，取决于你要避免什么"边界问题"

场景1：先处理进入事件（+1 优先）
适用场景：需要"贪心"尽早占据位置
sort(x0): 

场景2：先处理离开事件（-1 优先）
适用场景：区间端点相接不算重叠
sort(x1)

场景3：顺序无关紧要
适用场景：计算覆盖长度（3454 题）

### Step3 (Sweeping/Scan): 

O(n)

```golang
res := 0
cur := 0
for time, delta := range event {
  cur += delta
  res = max(cur, res)
}

```

## 模版

```golang
type SegTree struct {
    n    int
    tree []int
}

func NewSegTree(arr []int) *SegTree {
    n := len(arr)
    st := &SegTree{n: n, tree: make([]int, 2*n)}
    
    // 构建树：叶子节点
    for i := 0; i < n; i++ {
        st.tree[n+i] = arr[i]
    }
    
    // 构建树：内部节点
    for i := n - 1; i > 0; i-- {
        st.tree[i] = st.tree[2*i] + st.tree[2*i+1]
    }
    
    return st
}

// 单点更新
func (st *SegTree) Update(pos, val int) {
    pos += st.n
    st.tree[pos] = val
    for pos > 0 {
        st.tree[pos>>1] = st.tree[pos] + st.tree[pos^1]
        pos >>= 1
    }
}

// 区间查询 [l, r)
func (st *SegTree) Query(l, r int) int {
    res := 0
    l += st.n
    r += st.n   // 如果是闭区间，则 r += st.n+1
    for l < r {
        if l&1 == 1 {
            res += st.tree[l]
            l++
        }
        if r&1 == 1 {
            r-- // 如果是闭区间则可以这一行在下面 然后 r & 1 == 0，l <= r; 或者开头 r+=st.n+1
            res += st.tree[r]
        }
        l >>= 1
        r >>= 1
    }
    return res
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11}
    st := NewSegTree(arr)
    
    fmt.Println(st.Query(1, 4))  // [1,4) => 3+5+7 = 15
    st.Update(2, 10)             // arr[2] = 10
    fmt.Println(st.Query(1, 4))  // 3+10+7 = 20
}
```


支持其他操作（改merge逻辑）

```golang
// 区间最大值：把加法改为max
func (st *SegTree) Update(pos, val int) {
    pos += st.n
    st.tree[pos] = val
    for pos > 1 {
        st.tree[pos>>1] = max(st.tree[pos], st.tree[pos^1]) // 改这里
        pos >>= 1
    }
}

func (st *SegTree) Query(l, r int) int {
    res := -1e9  // 改初始值
    l += st.n
    r += st.n
    for l < r {
        if l&1 == 1 {
            res = max(res, st.tree[l])  // 改这里
            l++
        }
        if r&1 == 1 {
            r--
            res = max(res, st.tree[r])  // 改这里
        }
        l >>= 1
        r >>= 1
    }
    return res
}
```
