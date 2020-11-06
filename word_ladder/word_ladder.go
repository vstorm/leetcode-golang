package word_ladder

import "math"

var wordNum int
var wordIds map[string]int
var edges map[int][]int

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordIds = make(map[string]int)
	edges = make(map[int][]int)
	addWord(beginWord)
	for _, word := range wordList {
		addWord(word)
	}
	if _, ok := wordIds[endWord]; !ok {
		return 0
	}

	dif := make([]int, wordNum)
	for i := 0; i < wordNum; i += 1 {
		dif[i] = math.MaxInt32
	}
	beginId, endId := wordIds[beginWord], wordIds[endWord]
	dif[beginId] = 0

	q := []int{wordIds[beginWord]}
	for ;len(q) != 0; {
		t := q[0]
		q = q[1:len(q)]
		if t == endId {
			return dif[endId] / 2 + 1
		} else {
			for _, id := range edges[t] {
				if dif[id] == math.MaxInt32 {
					dif[id] = dif[t] + 1
					q = append(q, id)
				}
			}
		}
	}

	return 0
}

func addWord(word string) {
	wordIds[word] = wordNum
	wordNum += 1
	var temp rune
	w := []rune(word)
	for i := range w {
		temp, w[i] = w[i], '*'
		nw := string(w)
		if _, ok := wordIds[nw]; !ok {
			wordIds[nw] = wordNum
			wordNum += 1
		}
		w[i] = temp
		addEdge(wordIds[word], wordIds[nw])
	}
}

func addEdge(id1, id2 int) {
	if edges[id1] == nil {
		edges[id1] = make([]int, 0)
	}
	if edges[id2] == nil {
		edges[id2] = make([]int, 0)
	}
	edges[id1] = append(edges[id1], id2)
	edges[id2] = append(edges[id2], id1)
}
