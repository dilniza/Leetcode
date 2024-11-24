package leetcode

func TwoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

// func twoSum(nums []int, target int) []int {
//     hashMap := make(map[int]int) // Maps numbers to their indices
//     for i, num := range nums {
//         complement := target - num
//         if index, found := hashMap[complement]; found {
//             return []int{index, i} // Return indices of the two numbers
//         }
//         hashMap[num] = i // Store the current number with its index
//     }
//     return nil // No solution found
// }

