package remove_k_digits

func removeKdigits(num string, k int) string {
	cs := []rune(num)
	topI := -1
	n := k
	for _, c := range cs {
		for ;topI >= 0 && k > 0; {
			if cs[topI] > c {
				k -= 1
				topI -= 1
				continue
			}
			break
		}
		topI += 1
		cs[topI] = c
	}
	start := 0
	for ; start < len(cs)-n; start += 1 {
		if cs[start] != '0' {
			break
		}
	}
	if start == len(cs)-n {
		return "0"
	}
	return string(cs[start:len(cs)-n])
}
