package shortest_word_distance_III

import "math"

func shortestWordDistance(words []string, word1 string, word2 string) int {
	id1, id2 := -1, -1
	minD := math.MaxInt32
	for i, w := range words {
		if w == word1 {
			id1 = i
			if id2 >= 0 {
				minD = min(minD, id1 - id2)
			}
		}
		if w == word2 {
			id2 = i
			if id1 >= 0 && id1 != id2 {
				minD = min(minD, id2 - id1)
			}
		}
	}
	return minD
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
