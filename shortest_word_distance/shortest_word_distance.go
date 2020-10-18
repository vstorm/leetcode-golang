package shortest_word_distance

import "math"

func shortestDistance(words []string, word1 string, word2 string) int {
	i1, i2 := -1, -1
	minDistance := len(words)
	for i := range words {
		if words[i] == word1 {
			i1 = i
		}
		if words[i] == word2 {
			i2 = i
		}
		if i1 != -1 && i2 != -1 {
			minDistance = min(minDistance, int(math.Abs(float64(i2 - i1))))
		}
	}
	return minDistance
}

func min(x, y int) int{
	if x < y {
		return x
	}
	return y
}