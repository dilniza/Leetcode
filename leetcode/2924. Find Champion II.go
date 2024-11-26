package leetcode

func FindChampion(n int, edges [][]int) int {
    // Step 1: Handle edge cases for invalid input
    if n == 0 {
        return -1
    }

    // Step 2: Initialize in-degree array
    inDegree := make([]int, n)

    // Step 3: Compute in-degrees from edges
    for _, edge := range edges {
        inDegree[edge[1]-1]++
    }

    // Step 4: Find nodes with in-degree 0
    champion := -1
    for i := 0; i < n; i++ {
        if inDegree[i] == 0 {
            if champion != -1 {
                // More than one team with in-degree 0
                return -1
            }
            champion = i
        }
    }

    // Step 5: Return the unique champion or -1 if none
    return champion
}



// func FindChampion(n int, edges [][]int) int {
// 	if n == 0 {
// 		return -1
// 	}

// 	inDegree := make([]int, n+1)
// 	outDegree := make([]int, n+1)

// 	for _, edge := range edges {
// 		outDegree[edge[0]]++
// 		inDegree[edge[1]]++
// 	}

// 	for i := 1; i <= n; i++ {
// 		if inDegree[i] == n-1 && outDegree[i] == 0 {
// 			return i
// 		}
// 	}

// 	return -1
// }