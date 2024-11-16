package leetcode

func ResultsArray(nums []int, k int) []int {
    var res []int

    for i := 0; i <= len(nums)-k; i++ {
        subarray := nums[i : i+k] // Extract the subarray of size k
        if isSortedAndConsecutive(subarray) {
            res = append(res, subarray[k-1]) // Append the maximum (last element)
        } else {
            res = append(res, -1) // Append -1 if not valid
        }
    }

    return res
}

// Check if a subarray is sorted in ascending order and has consecutive elements
func isSortedAndConsecutive(subarray []int) bool {
    for i := 0; i < len(subarray)-1; i++ {
        if subarray[i]+1 != subarray[i+1] { // Check for consecutive elements
            return false
        }
    }
    return true
}
