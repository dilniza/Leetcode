package leetcode

import (
	"container/list"
)

func MinimumObstacles(grid [][]int) int {
	// Get grid dimensions
	m, n := len(grid), len(grid[0])

	// Directions for moving up, down, left, and right
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Deque for 0-1 BFS
	deque := list.New()

	// Distance matrix to keep track of the minimum obstacles removed
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1e9 // Initialize with a large number
		}
	}

	// Start from the top-left corner
	dist[0][0] = 0
	deque.PushFront([]int{0, 0})

	for deque.Len() > 0 {
		// Pop from the deque
		cur := deque.Remove(deque.Front()).([]int)
		x, y := cur[0], cur[1]

		// Explore neighbors
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			// Check if the neighbor is within bounds
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				// Calculate new cost
				newCost := dist[x][y] + grid[nx][ny]

				// If a shorter path to (nx, ny) is found
				if newCost < dist[nx][ny] {
					dist[nx][ny] = newCost

					// Add to deque based on cost
					if grid[nx][ny] == 0 {
						deque.PushFront([]int{nx, ny}) // No obstacle, add to front
					} else {
						deque.PushBack([]int{nx, ny}) // Obstacle, add to back
					}
				}
			}
		}
	}

	// Return the minimum obstacles removed to reach the bottom-right corner
	return dist[m-1][n-1]
}
