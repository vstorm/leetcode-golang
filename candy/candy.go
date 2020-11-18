package candy

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := 0; i < n; i += 1 {
		candies[i] = 1
	}

	for i := 1; i < n; i += 1 {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i -= 1 {
		if ratings[i] > ratings[i+1] {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	sum := 0
	for _, c := range candies {
		sum += c
	}
	return sum
}
