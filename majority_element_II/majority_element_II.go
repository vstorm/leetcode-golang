package majority_element_II

func majorityElement(nums []int) []int {
	// n/k的众数最多只有k-1个
	// 每次都拿走3个不一样的数, 那么最后剩下的, 一定是次数出现最多的两个数，柱状图理解
	r := []int{}
	if len(nums) < 1 {
		return nums
	}
	card1, count1 := nums[0], 0
	card2, count2 := nums[0], 0
	for _, num := range nums {
		if card1 == num {
			count1 += 1
			continue
		}
		if card2 == num {
			count2 += 1
			continue
		}

		if count1 == 0 {
			card1 = num
			count1 += 1
			continue
		}
		if count2 == 0 {
			card2 = num
			count2 += 1
			continue
		}

		count1 -= 1
		count2 -= 1
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == card1 {
			count1 += 1
		} else if num == card2 {
			count2 += 1
		}

	}
	if count1 > len(nums) / 3 {
		r = append(r, card1)
	}
	if count2 > len(nums) / 3 {
		r = append(r, card2)
	}
	return r
}