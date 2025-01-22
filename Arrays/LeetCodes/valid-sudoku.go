package main

// https://leetcode.com/problems/valid-sudoku
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([][]map[byte]bool, 3, 3)
	for i := 0; i < 3; i++ {
		boxes[i] = make([]map[byte]bool, 3)

	}

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i/3][i%3] = make(map[byte]bool)

	}
	for i := range board {
		bi := i / 3

		for j := range board[i] {
			bj := j / 3
			value := board[i][j]

			if value != '.' {
				if boxes[bi][bj][value] {
					return false
				}
				if rows[i][value] {
					return false
				}
				if cols[j][value] {
					return false
				}
				rows[i][value] = true
				cols[j][value] = true
				boxes[bi][bj][value] = true
			}

		}
	}
	return true
}
