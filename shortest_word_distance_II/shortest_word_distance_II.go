package shortest_word_distance_II

import "math"

type WordDistance struct {
	m map[string][]int
}


func Constructor(words []string) WordDistance {
	wd := WordDistance{m: map[string][]int{}}
	for i, w := range words {
		wd.m[w] = append(wd.m[w], i)
	}
	return wd
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
	minD := math.MaxInt32
	l1, l2 := this.m[word1], this.m[word2]
	for i, j := 0, 0; i < len(l1) && j < len(l2); {
		minD = min(minD, int(math.Abs(float64(l1[i]-l2[j]))))
		if minD == 1 {
			return minD
		}
		if l1[i] < l2[j] {
			i += 1
		} else {
			j += 1
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


/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */
