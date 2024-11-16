package leetcode

func minimumSubarrayLength(nums []int, k int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    
    minLen := n + 1
    orMap := make(map[int]int) // OR value to earliest index
    orMap[0] = -1 // Base case
    
    currOrs := make(map[int]int)
    for i := 0; i < n; i++ {
        nextOrs := make(map[int]int)
        nextOrs[nums[i]] = i
        for orVal, idx := range currOrs {
            newOr := orVal | nums[i]
            if _, exists := nextOrs[newOr]; !exists || idx < nextOrs[newOr] {
                nextOrs[newOr] = idx
            }
        }
        currOrs = nextOrs

        for orVal, startIdx := range currOrs {
            if orVal >= k {
                minLen = min(minLen, i - startIdx + 1)
            }
        }
    }
    
    if minLen == n+1 {
        return -1
    }
    return minLen
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}