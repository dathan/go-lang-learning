package algo

func dfs(node int, nodes map[int][]int, fn func(int)) {
	dfs_recur(nodes, node, map[int]bool{}, fn)
}

func dfs_recur(nodes map[int][]int, node int, v map[int]bool, fn func(int)) {
	v[node] = true
	fn(node)
	for _, n := range nodes[node] {
		if _, ok := v[n]; !ok {
			dfs_recur(nodes, n, v, fn)
		}
	}
}
