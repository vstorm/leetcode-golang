package string_to_Integer

import (
	"math"
)

func myAtoi(s string) int {
	m := map[string][]string {
		"start": []string{"start", "symbol", "number", "end"},
		"symbol": []string{"end", "end", "number", "end"},
		"number": []string{"end", "end", "number", "end"},
		"end": []string{"end", "end", "end", "end"},
	}

	status := "start"
	symbol := 1
	result := 0
	for _, c := range s {
		status = m[status][getCol(c)]
		if status == "symbol" {
			if c == '-' {
				symbol = -1
			}
		} else if status == "number" {
			num := int(c - '0')
			complete := math.MaxInt32 / 10
			if complete == result {
				if symbol == 1 {
					remainder := math.MaxInt32 % 10
					if remainder < num {
						return math.MaxInt32
					}
				} else {
					remainder := -(math.MinInt32 % 10)
					if remainder < num {
						return math.MinInt32
					}
				}
			} else if complete < result {
				if symbol == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
			result = result*10 + num
		} else if status == "end" {
			break
		}
	}
	return symbol * result
}

func getCol(c int32) int {
	if c == ' ' {
		return 0
	} else if c == '-' || c == '+' {
		return 1
	} else if c >= '0' && c <= '9' {
		return 2
	} else {
		return 3
	}
}
