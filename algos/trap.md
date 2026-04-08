# 陷阱

## 78. Subsets
```golang
func compute(nums []int, res *[][]int) {
    if len(nums) == 0 {
        return
    }
    for _, e := range *res {
        *res = append(*res, append(e, nums[0])) // 会改底层数组
    }
    compute(nums[1:], res)
}

func compute(nums []int, res *[][]int) {
    if len(nums) == 0 {
        return
    }
    for _, e := range *res {
        newSubset := append([]int{}, e...)
        newSubset = append(newSubset, nums[0])
        *res = append(*res, newSubset)
    }
    compute(nums[1:], res)
}
```
