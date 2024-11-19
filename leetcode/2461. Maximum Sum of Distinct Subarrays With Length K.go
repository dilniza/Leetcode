package leetcode

func MaximumSubarraySum(nums []int, k int) int64 {
    maxSum := int64(0)
    currentSum := int64(0)
    freq := make(map[int]int)
    n := len(nums)

    for i := 0; i < n; i++ {
        // Add the current number to the window
        currentSum += int64(nums[i])
        freq[nums[i]]++

        // If the window size exceeds k, remove the leftmost element
        if i >= k {
            left := nums[i-k]
            currentSum -= int64(left)
            freq[left]--
            if freq[left] == 0 {
                delete(freq, left)
            }
        }

        // Check if the current window contains k distinct elements
        if len(freq) == k {
            if currentSum > maxSum {
                maxSum = currentSum
            }
        }
    }

    return maxSum
}


// func MaximumSubarraySum(nums []int, k int) int64 {
// 	maxSum := int64(0)
// 	length := len(nums)

// 	for i := 0; i <= length-k; i++ {
// 		subarr := nums[i : i+k]
// 		if isDistinct(subarr) {
// 			sum := int64(0)
// 			for _, v := range subarr {
// 				sum += int64(v)
// 			}
// 			if sum > maxSum {
// 				maxSum = sum
// 			}
// 		}
// 	}

// 	return maxSum
// }

// func isDistinct(arr []int) bool {
// 	seen := make(map[int]bool)
// 	for _, v := range arr {
// 		if seen[v] {
// 			return false
// 		}
// 		seen[v] = true
// 	}
// 	return true
// }
