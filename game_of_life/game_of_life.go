package game_of_life

func gameOfLife(board [][]int) {
	// -1-活变死的细胞，0-死细胞, 1-活细胞，2-死变活的细胞
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	neighbors := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			live := 0
			for _, neighbor := range neighbors {
				row, col := i+neighbor[0], j+neighbor[1]
				if row >= 0 && col >= 0 && row < m && col < n {
					if board[row][col] == 1 || board[row][col] == -1 {
						live += 1
					}
				}
			}
			// 规则1和规则3
			if board[i][j] == 1 && (live < 2 || live > 3) {
				board[i][j] = -1
			}
			if board[i][j] == 0 && live == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if board[i][j] == -1 {
				board[i][j] = 0
			}
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}