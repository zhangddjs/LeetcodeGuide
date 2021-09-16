
/**
 * step1: build a 2-D array arr, use i, j to identify the index of left and right side.
 * initial i = 0, j = n - 1, arr[i][j] = -2
 * 
 * step2: then use dp to compute the min Operations
 *
 * will OOM
 */

var dp [][]int

func minOperations(nums []int, x int) int {
    dp = make([][]int, len(nums))
    for i := range dp {
        dp[i] = make([]int, len(nums))
    }
    return compute(nums, x, 0, len(nums) - 1);
}

func compute(nums []int, x, i, j int) int {
    if i > j && x != 0 {
        return -1
    }
    if x < 0 {
        dp[i][j] = -1
        return -1
    }
    if x == 0 {
        return 0
    }
    if dp[i][j] != 0 {
        return dp[i][j]
    }

    left := compute(nums, x - nums[i], i + 1, j)
    right := compute(nums, x - nums[j], i, j - 1)
    if left == -1 && right == -1 {
        dp[i][j] = -1
        return -1
    }
    dp[i][j] = getMinOperations(left, right)
    return dp[i][j]
}

func getMinOperations(left, right int) int {
    if left == -1 {
        return right + 1
    }
    if right == -1 {
        return left + 1
    }
    min := left
    if right < left {
        min = right
    }
    return min + 1
}


//----------------------------------------------------
func minOperations(nums []int, x int) int {
    i, j, count, minOpt, flag := 0, len(nums) - 1, 0, len(nums) - 1, false
    for i <= j && nums[i] <= x {    // find the left bound
        x -= nums[i]
        i++
        count++
    }
    if i > j {    // means the sum of all elems is <= x
        if x == 0 {
            return count
        } else {
            return -1
        }
    }
    if x == 0 {    //means x can be minused to 0, set temp min Opt nums
        flag = true
        minOpt = min(minOpt, count)
    }
    for i >= 0 {    // find the right bound
        if nums[j] > x && i > 0 {    // x can not be minused, slide the left bound
            i--
            x += nums[i]
            count--
        }
        if nums[j] > x {    // left bound can not be slided, because is left most
            if i == 0 {
                break
            } else {
                continue
            }
        }
        x -= nums[j]    // slide the right bound
        j--
        count++
        if x == 0 {    // record the opt num
            flag = true
            minOpt = min(minOpt, count)
        }
    }
    if flag {
        return minOpt
    }
    return -1
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}


//Brute Force : just traverse.