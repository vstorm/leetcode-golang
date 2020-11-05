package surrounded_regions

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				dfs(i, j, board)
			}
		}
	}

	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	return
}

func dfs(i, j int, board [][]byte) {
	board[i][j] = 'A'

	m := len(board)
	n := len(board[0])
	if i > 1 && board[i-1][j] == 'O' {
		dfs(i-1, j, board)
	}
	if j > 1 && board[i][j-1] == 'O' {
		dfs(i, j-1, board)
	}
	if i < m-1 && board[i+1][j] == 'O' {
		dfs(i+1, j, board)
	}
	if j < n-1 && board[i][j+1] == 'O' {
		dfs(i, j+1, board)
	}
}
