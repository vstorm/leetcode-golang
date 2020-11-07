package n_queens_II

var chessboards []int
var num int

func totalNQueens(n int) int {
	chessboards = make([]int, n)
	num = 0
	putQueen(0, n)
	return num
}

func putQueen(i, n int) { // 按行放置
	if i == n {
		num += 1
	} else {
		for j := 0; j < n; j += 1 {
			chessboards[i] = j
			if checkValid(i) {
				putQueen(i+1, n)
			}
		}
	}
}

func checkValid(i int) bool {
	for j := 0; j < i; j += 1 {
		// 同列，左对角线，右对角线
		if (chessboards[i] == chessboards[j]) || (chessboards[i]-chessboards[j] == i-j) || (chessboards[j]-chessboards[i] == i-j) {
			return false
		}
	}
	return true
}
