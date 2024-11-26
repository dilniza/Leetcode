package leetcode

import (
	"container/heap"
	"strings"
)

type State struct {
	Board [][]int
	Cost  int
	Steps int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func SlidingPuzzle(board [][]int) int {
	goalBoard := [][]int{{1, 2, 3}, {4, 5, 0}}
	initialState := State{Board: board, Cost: calculateManhattanDistance(board, goalBoard)}
	pq := &PriorityQueue{initialState}
	heap.Init(pq)
	visited := make(map[string]bool)
	visited[boardString(board)] = true

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		if boardString(current.Board) == boardString(goalBoard) {
			return current.Steps
		}

		for _, nextState := range getPossibleMoves(current.Board) {
			boardKey := boardString(nextState.Board)
			if !visited[boardKey] {
				visited[boardKey] = true
				nextState.Cost = current.Steps + 1 + calculateManhattanDistance(nextState.Board, goalBoard)
				nextState.Steps = current.Steps + 1
				heap.Push(pq, nextState)
			}
		}
	}

	return -1
}

func calculateManhattanDistance(board, goal [][]int) int {
	distance := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			val := board[i][j]
			if val != 0 {
				goalRow, goalCol := findGoalPosition(val, goal)
				distance += absSP(i-goalRow) + absSP(j-goalCol)
			}
		}
	}
	return distance
}

func findGoalPosition(val int, goal [][]int) (int, int) {
	for i := 0; i < len(goal); i++ {
		for j := 0; j < len(goal[0]); j++ {
			if goal[i][j] == val {
				return i, j
			}
		}
	}
	return -1, -1
}

func boardString(board [][]int) string {
	var builder strings.Builder
	for _, row := range board {
		for _, val := range row {
			builder.WriteByte(byte(val) + '0')
		}
	}
	return builder.String()
}

func getPossibleMoves(board [][]int) []State {
	moves := []State{}
	row, col := findEmpty(board)
	possibleMoves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, move := range possibleMoves {
		newRow, newCol := row+move[0], col+move[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) {
			newBoard := copyBoard(board)
			newBoard[row][col], newBoard[newRow][newCol] = newBoard[newRow][newCol], newBoard[row][col]
			moves = append(moves, State{Board: newBoard})
		}
	}
	return moves
}

func copyBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for i := range board {
		newBoard[i] = make([]int, len(board[0]))
		copy(newBoard[i], board[i])
	}
	return newBoard
}

func findEmpty(board [][]int) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func absSP(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
