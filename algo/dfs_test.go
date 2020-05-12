package algo

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {

	var nodes = map[int][]int{
		1:  []int{2, 7, 8},
		2:  []int{1, 3, 6},
		3:  []int{2, 4, 5},
		4:  []int{3},
		5:  []int{3},
		6:  []int{2},
		7:  []int{1},
		8:  []int{1, 9, 12},
		9:  []int{8, 10, 11},
		10: []int{9},
		11: []int{9},
		12: []int{8},
	}

	visited := []int{}
	dfs(1, nodes, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)

}
