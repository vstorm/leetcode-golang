package dungeon_game

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	for i := m-1; i >= 0; i -= 1 {
		for j := n-1; j >= 0; j -= 1 {
			if i == m - 1 && j == n - 1 {
				dungeon[i][j] = max(1 - dungeon[i][j], 1)
			} else if i == m - 1 {
				dungeon[i][j] = max(dungeon[i][j+1] - dungeon[i][j], 1)
			} else if j == n - 1 {
				dungeon[i][j] = max(dungeon[i+1][j] - dungeon[i][j], 1)
			} else {
				dungeon[i][j] = max(min(dungeon[i][j+1] - dungeon[i][j], dungeon[i+1][j] - dungeon[i][j]), 1)
			}
		}
	}
	return dungeon[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}