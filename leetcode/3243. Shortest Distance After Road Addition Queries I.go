package leetcode

import "container/list"

func ShortestDistanceAfterQueries(n int, queries [][]int) []int {
	// Create an adjacency list to represent the graph
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

	// Helper function to perform BFS and find the shortest path from 0 to n-1
	bfsShortestPath := func() int {
		visited := make([]bool, n)
		queue := list.New()
		queue.PushBack([2]int{0, 0}) // {current node, current distance}

		for queue.Len() > 0 {
			front := queue.Remove(queue.Front()).([2]int)
			node, dist := front[0], front[1]

			if node == n-1 {
				return dist
			}

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue.PushBack([2]int{neighbor, dist + 1})
				}
			}
		}

		return -1 // Shouldn't happen if the graph is properly connected
	}

	// Process each query and update the graph
	result := make([]int, len(queries))
	for i, query := range queries {
		u, v := query[0], query[1]
		graph[u] = append(graph[u], v) // Add the new road
		result[i] = bfsShortestPath()
	}

	return result
}
