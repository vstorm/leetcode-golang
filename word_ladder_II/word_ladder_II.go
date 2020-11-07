package word_ladder_II

import (
	"math"
)

var idWords map[int]string
var wordIds map[string]int
var edges map[int][]int

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	idWords = make(map[int]string)
	wordIds = make(map[string]int)
	edges = make(map[int][]int)
	results := make([][]string, 0)
	// 建立id到单词和单词到id的映射
	for i, word := range wordList {
		idWords[i] = word
		wordIds[word] = i
	}
	if _, ok := wordIds[beginWord]; !ok {
		idWords[len(wordList)] = beginWord
		wordIds[beginWord] = len(wordList)
		wordList = append(wordList, beginWord)
	}

	// 如果endWord不存在
	if _, ok := wordIds[endWord]; !ok {
		return results
	}

	// 建立转换关系图
	for i := 0; i < len(wordList)-1; i += 1 {
		for j := i + 1; j < len(wordList); j += 1 {
			w1, w2 := wordList[i], wordList[j]
			if canTransform(w1, w2) {
				addEdge(wordIds[w1], wordIds[w2])
			}
		}
	}
	// 从beginWord开始，到达该单词的最小步数
	costs := make([]int, len(wordIds))
	for i := 0; i < len(costs); i += 1 {
		costs[i] = math.MaxInt32
	}
	beginId, endId := wordIds[beginWord], wordIds[endWord]
	costs[beginId] = 0

	// 注意：队列保存的是到达该单词的整条路径
	q := make([][]int, 0)
	q = append(q, []int{beginId})

	// 开始广度遍历
	for ; len(q) != 0; {
		now := q[0]
		q = q[1:len(q)]
		last := now[len(now)-1]
		if last == endId {
			tmp := make([]string, 0)
			for _, id := range now {
				tmp = append(tmp, idWords[id])
			}
			results = append(results, tmp)
		} else {
			for _, to := range edges[last] {
				if costs[last] + 1 <= costs[to] {	// 注意这个判断条件
					costs[to] = costs[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					q = append(q, tmp)
				}
			}
		}
	}
	return results
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

func canTransform(w1, w2 string) bool {
	diff := 0
	for i:=0; i < len(w1); i += 1 {
		if w1[i] != w2[i] {
			diff += 1
		}
	}
	return diff == 1
}