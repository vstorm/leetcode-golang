package longest_palindromic_substring

// 中心扩散法
// 枚举可能出现的回文子串的“中心位置”，从“中心位置”尝试尽可能扩散出去，得到一个回文串。
// 注意：我们要遍历每一个回文串的中心位置，奇数和偶数回文串的中心位置是不同的，可以查看下面链接，查看字符串可能的回文子串的中心都在哪些位置上
// https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zhong-xin-kuo-san-dong-tai-gui-hua-by-liweiwei1419/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := palindromeLength(s, i, i)
		left2, right2 := palindromeLength(s, i, i+1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func palindromeLength(s string, left, right int) (int, int) {
	for ; left >= 0 && right <= len(s)-1 && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
