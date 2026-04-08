# Perm

| 枚举对象      | 数量级   | 复杂度类型  | 示例              |
| ---------     | -------- | ----------- | ------            |
| 数组元素      | n        | O(n)        | 线性              |
| 所有对 (i,j)  | n²       | O(n²)       | 多项式            |
| 所有三元组    | n³       | O(n³)       | 多项式            |
| k个元素的组合 | C(n,k)   | O(n^k)      | 多项式（k固定时） |
| 所有子集      | 2^n      | O(2^n)      | 指数级            |
| 所有排列      | n!       | O(n!)       | 阶乘级            |
| 所有分区      | Bell(n)  | O(n^n)      | 超指数            |
| 所有树        | n^(n-2)  | O(n^(n-2))  | 指数级            |

## Template

```golang
func backtrack(路径, 选择列表) {
  if 满足结束条件 {
    result = append(result, 路径的拷贝)
    return
  }
  
  for 选择 in 选择列表 {
    做选择 (将该选择从选择列表移除, 加入路径)
    backtrack(路径, 新的选择列表)
    撤销选择 (将该选择从路径移除, 恢复选择列表)
  }
}
```


- 78.Subset
- 90.Subset II

```golang
func subsetsWithDup(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)   // easy for deduplicate
    compute(nums, 0, &[]int{}, &res)
    return res
}

func compute(nums []int, start int, path *[]int, res *[][]int) {
    tmp := append([]int{}, (*path)...)
    *res = append(*res, tmp)
    for i := start; i < len(nums); i++ {
        if i > start && nums[i] == nums[i-1] {
            continue    // deduplicate
        }
        *path = append(*path, nums[i])
        compute(nums, i+1, path, res)
        *path = (*path)[:len(*path)-1]
    }
}
```


