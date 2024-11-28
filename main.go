package main

import (
	"fmt"
	"leet/leetcode"
)

func main() {
	grid := [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}
	res := leetcode.MinimumObstacles(grid)
	fmt.Println(res)

	grid = [][]int{{0, 1, 1}, {1, 1, 0}}
	res = leetcode.MinimumObstacles(grid)
	fmt.Println(res)
}

// func main() {
// 	queries := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
// 	n := 5
// 	res := leetcode.ShortestDistanceAfterQueries(n, queries)
// 	fmt.Println(res)
// }

// func main() {
// 	n := 3
// 	edges := [][]int{{1, 2}, {2, 3}}
// 	res := leetcode.FindChampion(n, edges)
// 	fmt.Println(res)

// 	n = 4
// 	edges = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}
// 	res = leetcode.FindChampion(n, edges)
// 	fmt.Println(res)
// }

// func main() {
// 	board := [][]int{{1,2,3}, {4,0,5}}
// 	move := leetcode.SlidingPuzzle(board)
// 	fmt.Println(board, "moves: ", move)

// 	board = [][]int{{1,2,3}, {5,4,0}}
// 	move = leetcode.SlidingPuzzle(board)
// 	fmt.Println(board, "moves: ", move)

// 	board = [][]int{{4,1,2}, {5,0,3}}
// 	move = leetcode.SlidingPuzzle(board)
// 	fmt.Println(board, "moves: ", move)
// }

// func main() {
// 	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	res := leetcode.MaxMatrixSum(matrix)
// 	fmt.Println(res)

// 	matrix = [][]int{{1, -1}, {-1, 1}}
// 	res = leetcode.MaxMatrixSum(matrix)
// 	fmt.Println(res)

// 	box := [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}
// 	res1 := leetcode.RotateTheBox(box)
// 	fmt.Println(box)
// 	fmt.Println(res1)

// 	box = [][]byte{{'#', '.', '#'}}
// 	res1 = leetcode.RotateTheBox(box)
// 	fmt.Println(res1)

// 	target := 5
// 	nums := []int{1, 2, 3, 4}
// 	res2 := leetcode.TwoSum(nums, target)
// 	fmt.Println(res2)
// }

// func main() {
// 	s := "abc"
// 	k := 2
// 	res := leetcode.TakeCharacters(s, k)
// 	fmt.Println(res)

// 	s = "abc"
// 	k = 3
// 	res = leetcode.TakeCharacters(s, k)
// 	fmt.Println(res)
// }

// func main() {
// 	code := []int{5, 7, 1, 4}
// 	k := 2
// 	res := leetcode.MaximumSubarraySum(code, k)
// 	fmt.Println(res)

// 	code = []int{4, 4, 4}
// 	k = 3
// 	res = leetcode.MaximumSubarraySum(code, k)
// 	fmt.Println(res)
// }

// func main() {
// 	code := []int{5, 7, 1, 4}
// 	k := 3
// 	res := leetcode.Decrypt(code, k)
// 	fmt.Println(res)

// 	code = []int{2, 4, 9, 3}
// 	k = 2
// 	res = leetcode.Decrypt(code, k)
// 	fmt.Println(res)

// 	code = []int{2, 4, 9, 3}
// 	k = -2
// 	res = leetcode.Decrypt(code, k)
// 	fmt.Println(res)
// }

// func main() {
// 	var nums []int
// 	nums = []int{1,11,10}
// 	k := 3
// 	res := leetcode.ResultsArray(nums, k)
// 	fmt.Println(res)

// 	nums = []int{1,2,9,8}
// 	k = 2
// 	res = leetcode.ResultsArray(nums, k)
// 	fmt.Println(res)
// }
