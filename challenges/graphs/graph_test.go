package challenge_test

import (
	"fmt"
	"reflect"
	"testing"
)

/*

Minimizing Permutations
In this problem, you are given an integer N, and a permutation,
P of the integers from 1 to N, denoted as (a_1, a_2, ..., a_N).
You want to rearrange the elements of the permutation into increasing order,
repeatedly making the following operation:
Select a sub-portion of the permutation, (a_i, ..., a_j), and reverse its order.

Your goal is to compute the minimum number of such operations
required to return the permutation to increasing order.

Signature

int minOperations(int[] arr)

Input
Array arr is a permutation of all integers from 1 to N, N is between 1 and 8


Output
An integer denoting the minimum number of operations required to arrange the permutation in increasing order

Example
If N = 3, and P = (3, 1, 2), we can do the following operations:

Select (1, 2) and reverse it: P = (3, 2, 1).

Select (3, 2, 1) and reverse it: P = (1, 2, 3).

output = 2

*/
func minOperations(arr []int) int {
	// Write your code here
	return 0
}

// https://leetcode.com/problems/permutations/discuss/346440/golang-solution-using-bfs-and-queue
func permute(nums []int) [][]int {
	queue := [][]int{{}}
	for len(queue[0]) != len(nums) {

		currentPermutation := queue[0]

		queue = queue[1:]
		nextChar := nums[len(currentPermutation)]

		for i := 0; i <= len(currentPermutation); i++ {
			// note the visted is hidden in the slice trick of low : high e.g. start : length from 0
			newPermutation := append(
				append(

					append([]int{}, currentPermutation[:i]...),

					[]int{nextChar}...,
				),
				currentPermutation[i:]...,
			)

			// above is building multiple rows of the queue matrix
			queue = append(queue, newPermutation)
		}
	}
	return queue
}

// map type we can change this to *Node if we defined it
type ID int
type Graph struct {
	Nodes map[ID]struct{}        // could keep a stat of how many children
	Edges map[ID]map[ID]struct{} // could be  map of an array of node structs that holds the weight between nodes.
}

func New() *Graph {
	g := &Graph{
		Nodes: make(map[ID]struct{}),
		Edges: make(map[ID]map[ID]struct{}),
	}

	return g
}

func reverseit(arr []int) []int {
	ret := []int{}
	for i := len(arr) - 1; i >= 0; i-- {

		ret = append(ret, arr[i])

	}
	return ret
}

func (g *Graph) BFS(start int) []int {

	frontier := []int{start}
	visited := map[int]bool{} // a graph may have a cycle
	next := []int{}
	visits := []int{}
	for 0 < len(frontier) {
		next = []int{}
		for _, node := range frontier {

			visits = append(visits, node) // recording the fact that the node was visited from the test
			visited[node] = true
			for _, n := range bfs_frontier(node, g.Edges, visited) {
				next = append(next, n)
			}
		}
		frontier = next
	}

	return append([]int{}, visits[:len(visits)-1]...) // do not record that you ended on yourself
}

func bfs_frontier(node int, nodes map[ID]map[ID]struct{}, visited map[int]bool) []int {
	next := []int{}
	iter := func(n int) bool { _, ok := visited[n]; return !ok }
	for n, _ := range nodes[ID(node)] {
		if iter(int(n)) {
			next = append(next, int(n))
		}
	}
	return next
}

// using a builder design pattern let's build nodes and edges
func (g *Graph) AddNode(id ID) *Graph {

	if _, ok := g.Nodes[id]; ok {
		return g
	}

	g.Nodes[id] = struct{}{}
	return g

}

func (g *Graph) AddEdge(a, b ID) *Graph {

	if _, ok := g.Nodes[a]; !ok {
		g.AddNode(a)
	}
	if _, ok := g.Nodes[b]; !ok {
		g.AddNode(b)
	}

	if _, ok := g.Edges[a]; !ok {
		g.Edges[a] = make(map[ID]struct{})
	}

	g.Edges[a][b] = struct{}{}

	return g
}

func (g *Graph) String() string {

	ret := ""

	for k, _ := range g.Nodes {
		ret += fmt.Sprintf("NODE: %d\n", k)
	}

	for node, edges := range g.Edges {
		ret += fmt.Sprintf("NODE: %d EDGES: %v\n", node, edges)
		for k, _ := range edges {
			ret += fmt.Sprintf("\tEDGE: %d -> %d\n", node, k)
		}
	}

	return ret
}

func TestGraph_AddEdge(t *testing.T) {
	type args struct {
		a ID
		b ID
	}

	g := New()

	tests := []struct {
		name string
		g    *Graph
		args args
		want *Graph
	}{
		{"TestAddEdges", g, args{0, 1}, g},
		{"TestAddEdges", g, args{1, 2}, g},
		{"TestAddEdges", g, args{1, 3}, g},
		{"TestAddEdges", g, args{3, 4}, g},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.AddEdge(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.AddEdge() = %v, want %v", got, tt.want)
			}
		})
	}

	fmt.Printf("Graph: %s\n", g)
}

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"debug", args{[]int{3, 1, 2}},
			[][]int{
				[]int{2, 1, 3},
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 3, 1},
				[]int{3, 2, 1},
				[]int{3, 1, 2},
			},
		},
		{"debug", args{[]int{3, 1, 2, 4}},
			[][]int{{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
