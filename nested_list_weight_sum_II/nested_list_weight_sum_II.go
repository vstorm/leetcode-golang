package nested_list_weight_sum_II

func depthSumInverse(nestedList []*NestedInteger) int {
	maxD := calDepth(1, nestedList)
	return dfs(maxD, nestedList)
}

func calDepth(i int, nestedList []*NestedInteger) int {
	maxD := 0
	for _, n := range nestedList {
		if n.IsInteger() {
			maxD = max(maxD, i)
		} else {
			d := calDepth(i+1, n.GetList())
			maxD = max(maxD, d)
		}
	}
	return maxD
}

func dfs(i int, nestedList []*NestedInteger) int {
	sum := 0

	for _, n := range nestedList {
		if n.IsInteger() {
			sum += i * n.GetInteger()
		} else {
			sum += dfs(i-1, n.GetList())
		}
	}
	return sum
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}