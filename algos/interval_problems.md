## Difference Array 差分数组

- [ ] 1109
- [ ] 1094
- [ ] 370
- [ ] 798

**差分数组的本质**：
```
前缀和 ← 差分数组 → 原数组

构建：arr[i] - arr[i-1] = diff[i]
还原：sum(diff[0..i]) = arr[i]
```

**记忆口诀**：
```
区间加法变单点
左端加，右端减
最后前缀和还原
```


差分数组是处理批量区间更新 + 最后统一查询的神器，实现简单，效率极高！

所有场景几乎都可以用线段树

## Line Sweep

218. The Skyline Problem
391. Perfect Rectangle
759. Employee Free Time
850. Rectangle Area II
1851. Minimum Interval to Include Each Query
3009. Maximum Number of Intersections on the Chart
3454. Separate Squares II
986. Interval List Intersections

## 区间贪心专题

435. Non-overlapping Intervals ⭐⭐
452. Minimum Number of Arrows to Burst Balloons ⭐⭐
56. Merge Intervals ⭐⭐


## Meeting Rooms

### Meeting Rooms I (LeetCode 252)
给定一个会议时间安排的数组 intervals，其中 `intervals[i]` = `[starti, endi]`
，判断一个人是否能够参加所有会议。
示例 1:
```
输入: intervals = [[0,30],[5,10],[15,20]]
输出: false
```

示例 2:
```
输入: intervals = [[7,10],[2,4]]
输出: true
```

#### 解法
```golang

// clarify:
// 1. constranints of the intervals range
// 2. can fit in memory?
// 3. maximum number of intervals
// 4. interval cell be integer only or can be float
// 5. solve(intervals [][]int) bool

// 评价：注意思路，meeting rooms 如果有重叠那就不能参加完所有会议，而我想成了必须所有都重叠，这是不对的。

func solve(intervals [][]int) bool {
  // sort Onlogn
  sort.Slice(intervals, func (i, j int) bool {return intervals[i][0] < intervals[j][0]})
  // scan On
  for i := 1; i < len(intervals); i++ {
    if max(intervals[i][0], intervals[i-1][0]) < min(intervals[i][1], intervals[i-1][1]) {
      return false
    }
  }
  return true
}
```

### Meeting Rooms II (LeetCode 253)
给定一个会议时间安排的数组 intervals，其中 `intervals[i]` = `[starti, endi]`，返回所需的最小会议室数量。
示例 1:
```
输入: intervals = [[0,30],[5,10],[15,20]]
输出: 2
```
示例 2:
```
输入: intervals = [[7,10],[2,4]]
输出: 1
```
约束条件:
```
1 <= intervals.length <= 10^4
0 <= starti < endi <= 10^6
```

```golang
// clarify 
// 1. length of intervals , can fit in memory?
// 2. each interval length == 2?
// 3. start and end of interval can be interger or floats?
// 4. range of start and end
// 5. solve(intervals [][]int) int

// 评价：注意一个陷阱，同一时间点，结束事件应该排在开始事件之前
func solve(intervals [][]int) int {
  events := buildEvent(intervals)
  sort.Slice(events, func(i, j int) bool {
    if events[i][0] == events[j][0] {
      return events[i][1] < events[j][1]
    }
    return events[i][0] < events[j][0]
  })
  res, cur := 0, 0
  for _, e := range events {
    cur += e[1]
    res = max(cur, res)
  }
  return res
}

func buildEvent(intervals [][]int) [][]int {
  res := make([][]int, 0, 2*len(intervals))
  for _, i := range intervals {
    res = append(res, []int{i[0], 1}, []int{i[1], -1})
  }
  return res
}
```

### Meeting Rooms III (LeetCode 2402)

双heap
