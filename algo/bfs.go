package algo

/**
*
Breadth first traversal is accomplished by enqueueing each level of a tree
sequentially as the root of any subtree is encountered. There are 2 cases in
the iterative algorithm.

Root case : The traversal queue is initially empty so the root node must be
added before the general case.
General case: Process any items in the queue, while also expanding their
children, stop if the queue was empty. The general case will halt after
processing the bottom level as leaf nodes have no children.

Input: A search problem. A search-problem abstracts out the problem
specific requirements from the actual search algorithm.

Output: An ordered list of actions to be followed to reach from start state
to the goal state.
*/
func bfs(start int, nodes map[int][]int, fn func(int)) {
	frontier := []int{start}
	visited := map[int]bool{}
	next := []int{}

	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {
			visited[node] = true
			fn(node) // recording the fact that the node was visited from the test
			for _, n := range bfs_frontier(node, nodes, visited) {
				next = append(next, n)
			}
		}
		frontier = next
	}
}

func bfs_frontier(node int, nodes map[int][]int, visited map[int]bool) []int {
	next := []int{}
	iter := func(n int) bool { _, ok := visited[n]; return !ok }
	for _, n := range nodes[node] {
		if iter(n) {
			next = append(next, n)
		}
	}
	return next
}
