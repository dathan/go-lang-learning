package algo

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	var nodes = map[int][]int{
		1:  []int{2, 3, 4},
		2:  []int{1, 5, 6},
		3:  []int{1},
		4:  []int{1, 7, 8},
		5:  []int{2, 9, 10},
		6:  []int{2},
		7:  []int{4, 11, 12},
		8:  []int{4},
		9:  []int{5},
		10: []int{5},
		11: []int{7},
		12: []int{7},
	}

	visited := []int{}
	bfs(12, nodes, func(node int) {
		visited = append(visited, node)
	})
	fmt.Println(visited)
}
