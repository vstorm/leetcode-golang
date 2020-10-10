package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	ml := 0
	m := make(map[byte]bool)
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			for ;;{
				if s[i] != s[j] {
					delete(m, s[i])
					i++
				} else {
					i++
					break
				}
			}
		} else {
			l := j - i + 1
			if l > ml {
				ml = l
			}
			m[s[j]] = true
		}
	}
	return ml
}