//Depth-first search is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking. So the basic idea is to start from the root or any arbitrary node and mark the node and move to the adjacent unmarked node and continue this loop until there is no unmarked adjacent node. Then backtrack and check for other unmarked nodes and traverse them. Finally print the nodes in the path.
// Algorithm:
// Create a recursive function that takes the index of node and a visited array.
// Mark the current node as visited and print the node.
// Traverse all the adjacent and unmarked nodes and call the recursive function with index of adjacent node.

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
