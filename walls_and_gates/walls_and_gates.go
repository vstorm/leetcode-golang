package walls_and_gates

import "math"

func wallsAndGates(rooms [][]int) {
	gate, empty := 0, math.MaxInt32
	m := len(rooms)
	if m == 0 {
		return
	}
	n := len(rooms[0])

	q := make([][]int, 0)
	for i := 0; i < m; i += 1 {
		for j := 0; j < n; j += 1 {
			if rooms[i][j] == gate {
				q = append(q, []int{i, j})
			}
		}
	}

	directions := [][]int{
		[]int{-1, 0},
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
	}

	for ; len(q) != 0; {
		t := q[0]
		q = q[1:len(q)]
		row, col := t[0], t[1]

		for _, d := range directions {
			r, c := row+d[0], col+d[1]
			if r < 0 || c < 0 || r >= m || c >= n || rooms[r][c] != empty {
				continue
			}
			rooms[r][c] = rooms[row][col] + 1
			q = append(q, []int{r, c})
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}