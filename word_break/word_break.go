package word_break

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordMap := make(map[string]bool, n)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	result := make([]bool, n + 1)
	result[0] = true
	for i := 0; i < n; i += 1 {
		for j := -1; j < i; j += 1 {
			if result[j+1] == true && wordMap[s[j+1:i+1]] {
				result[i + 1] = true
				break
			}
		}
	}
	return result[n]
}
