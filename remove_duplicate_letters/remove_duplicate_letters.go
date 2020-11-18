package remove_duplicate_letters

func removeDuplicateLetters(s string) string {
	stack := make([]rune, 0, len(s))
	m := make(map[rune]int)
	for _, c := range s {
		m[c] += 1
	}

	set := make(map[rune]bool)
	for _, c := range s {
		m[c] -= 1
		if !set[c] {
			for ; len(stack) > 0; {
				top := stack[len(stack)-1]
				// 前一个字符，如果比当前字符大，而且后续还会出现
				if top > c && m[top] >= 1 {
					stack = stack[0 : len(stack)-1]
					delete(set, top)
					continue
				}
				break
			}
			stack = append(stack, c)
			set[c] = true
		}
	}
	return string(stack)
}
