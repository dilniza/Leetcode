package leetcode

func MaximumSubarraySum(nums []int, k int) int64 {
	maxSum := int64(0)
	length := len(nums)

	for i := 0; i <= length-k; i++ {
		subarr := nums[i : i+k]
		if isDistinct(subarr) {
			sum := int64(0)
			for _, v := range subarr {
				sum += int64(v)
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func isDistinct(arr []int) bool {
	seen := make(map[int]bool)
	for _, v := range arr {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}
