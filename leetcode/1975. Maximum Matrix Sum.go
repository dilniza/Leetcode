package leetcode

func MaxMatrixSum(matrix [][]int) int64 {
    n := len(matrix)
    totalSum := int64(0)
    smallestAbs := int(1e5 + 1) // Any value larger than the constraints
    negativeCount := 0

    // Step 1: Process the matrix
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            val := matrix[i][j]
            totalSum += int64(abs(val)) // Sum absolute values
            if val < 0 {
                negativeCount++ // Count negatives
            }
            smallestAbs = minMatrixSum(smallestAbs, abs(val)) // Track smallest absolute value
        }
    }

    // Step 2: Calculate maximum sum
    if negativeCount%2 == 0 {
        return totalSum
    } else {
        return totalSum - 2*int64(smallestAbs)
    }
}

// Helper functions
func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func minMatrixSum(a, b int) int {
    if a < b {
        return a
    }
    return b
}

