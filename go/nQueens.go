package main

import "strings"

func solveNQueens(n int) [][]string {
	var board [][]byte
	var ans [][]string

	var nQueens func(int)
	isSafe := func(row, col int) bool { // O(n)
		// 1. Check horizontal
		for j := range n {
			if board[row][j] == 'Q' {
				return false
			}
		}

		// 2. Check verticall
		for i := range n {
			if board[i][col] == 'Q' {
				return false
			}
		}

		// 3. Check left diagonal
		i, j := row, col
		for i >= 0 && j >= 0 {
			if board[i][j] == 'Q' {
				return false
			}
			i--
			j--
		}

		// 4. Check right diagonal
		i, j = row, col
		for i >= 0 && j < n {
			if board[i][j] == 'Q' {
				return false
			}
			i--
			j++
		}

		return true
	}

	nQueens = func(row int) {
		if row == n {
			temp := make([]string, n)
			// Deep copy
			for i := range board {
				temp[i] = string(board[i])
			}
			ans = append(ans, temp)
			return
		}

		for j := range n {
			if isSafe(row, j) {
				board[row][j] = 'Q'
				nQueens(row + 1)
				board[row][j] = '.' // Backtrack
			}
		}
	}

	// Initialize the board
	for range n {
		board = append(board, []byte(strings.Repeat(".", n)))
	}

	nQueens(0)

	return ans
}
