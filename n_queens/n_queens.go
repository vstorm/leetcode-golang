package n_queens

var chessboards []int
var results [][]string
var w []rune

func solveNQueens(n int) [][]string {
	chessboards = make([]int, n)
	results = make([][]string, 0)
	w = make([]rune, n)
	for i := range w {
		w[i] = '.'
	}
	putQueen(0, n)
	return results
}

func putQueen(i, n int) { // 按行放置
	if i == n {
		output()
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

func output() {
	result := make([]string, 0)
	for _, c := range chessboards {
		w[c] = 'Q'
		result = append(result, string(w))
		w[c] = '.'
	}
	results = append(results, result)
}