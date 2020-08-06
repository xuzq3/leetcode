/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (43.50%)
 * Likes:    1879
 * Dislikes: 92
 * Total Accepted:    190.4K
 * Total Submissions: 437.7K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 *
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 *
 *
 * Empty cells are indicated by the character '.'.
 *
 *
 * A sudoku puzzle...
 *
 *
 * ...and its solution numbers marked in red.
 *
 * Note:
 *
 *
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 *
 *
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	size := len(board) * len(board[0])
	p := 0
	np := -1

	var row [10][10]bool
	var col [10][10]bool
	var box [10][10]bool

	for ; p < size; p++ {
		i := p / 9
		j := p % 9
		c := board[i][j]
		if c != '.' {
			n := c - '0'
			if row[i][n] || col[j][n] || box[i/3*3+j/3][n] {
				return
			}
			row[i][n] = true
			col[j][n] = true
			box[i/3*3+j/3][n] = true
		} else {
			if np < 0 {
				np = p
			}
		}
	}
	if np >= 0 {
		r := solveNext(&board, size, np, row, col, box)
		_ = r
	}
	return
}

func solveNext(board *[][]byte, size int, p int, row [10][10]bool, col [10][10]bool, box [10][10]bool) bool {
	np := -1
	for ; p < size; p++ {
		i := p / 9
		j := p % 9
		c := (*board)[i][j]
		if c == '.' {
			np = p
			break
		}
	}
	if np < 0 {
		return true
	}

	i := np / 9
	j := np % 9
	for n := 1; n < len(row[i]); n++ {
		if !row[i][n] {
			if row[i][n] || col[j][n] || box[i/3*3+j/3][n] {
				continue
			}
			row[i][n] = true
			col[j][n] = true
			box[i/3*3+j/3][n] = true

			if !solveNext(board, size, np+1, row, col, box) {
				row[i][n] = false
				col[j][n] = false
				box[i/3*3+j/3][n] = false
				continue
			}
			(*board)[i][j] = byte(n + '0')
			return true
		}
	}
	return false
}

// @lc code=end

