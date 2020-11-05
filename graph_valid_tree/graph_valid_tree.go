package graph_valid_tree


var fa []int

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	fa = make([]int, n)
	for i := 0; i < n; i += 1 {
		fa[i] = i
	}

	for _, edge := range edges {
		i, j := edge[0], edge[1]
		fa[find(j)] = find(i)

	}

	for i := 0; i < n; i += 1 {
		find(i)
	}

	f := fa[0]
	for i := 1; i < n; i += 1 {
		if f != fa[i] {
			return false
		}
	}

	return true
}

func find(i int) int {
	if fa[i] == i {
		return i
	}
	fa[i] = find(fa[i])
	return fa[i]
}