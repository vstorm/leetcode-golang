package nested_list_weight_sum

func depthSum(nestedList []*NestedInteger) int {
	sum := 0
	for _, n := range nestedList {
		sum += dfs(1, n)
	}
	return sum
}

func dfs(i int, n *NestedInteger) int {
	sum := 0
	if n.IsInteger() {
		sum += i * n.GetInteger()
	} else {
		for _, n := range n.GetList() {
			sum += dfs(i+1, n)
		}
	}
	return sum
}