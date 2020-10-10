package valid_parentheses

func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) > 0 && stack[len(stack)-1] == m[c] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
