package challenges_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

/*
A builder is looking to build a row of N houses that can be of K different colors.
He has a goal of minimizing cost while ensuring that no two neighboring houses are of the same color.

Given an N by K matrix where the nth row and kth column represents the cost to build the nth house with kth color,
return the minimum cost which achieves this goal.
*/

type Point struct {
	i     int
	j     int
	color int
}

var rows int = 4
var cols int = 5

var HouseColorCost [][]int = make([][]int, cols)
var N int    // ROW
var K int    // k = kth color
var cost int // the n,k value

func setup() {

	for i := 0; i < cols; i++ {
		HouseColorCost[i] = make([]int, cols)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			x := i
			y := j
			x++
			y++
			HouseColorCost[i][j] = x * y
		}
	}

	var duplicates []Point
	// find the duplicates (build a list) for each spot look at the neighbors
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dupes := checkNeighbors(i, j)
			if len(dupes) > 0 {
				duplicates = append(duplicates, dupes...)
			}
		}
	}

	var done map[int]bool = make(map[int]bool)
	var used map[int]bool = make(map[int]bool)

	for _, point := range duplicates {
		if done[point.color] == true {
			fmt.Printf("Skip POINT: %+v\n", point)
			continue
		}
		for j := len(duplicates) - 1; j >= 0; j-- {
			replaceColor := duplicates[j]
			if replaceColor.color != point.color && !used[replaceColor.color] {
				fmt.Printf("POINT: +%v REPLACE: +%v\n", point, replaceColor)
				HouseColorCost[point.i][point.j] = replaceColor.color
				done[point.color] = true
				used[replaceColor.color] = true
				break
			}
		}
	}

	duplicates = []Point{}
	// find the duplicates (build a list) for each spot look at the neighbors
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			dupes := checkNeighbors(i, j)
			if len(dupes) > 0 {
				duplicates = append(duplicates, dupes...)
			}
		}
	}

	spew.Dump(duplicates)

}

// this assumes a grid NxN - we are wasting space to avoid maxing this contextual
func checkNeighbors(i, j int) []Point {
	fmt.Println("Checking neighbor")
	colorCheck := HouseColorCost[i][j]
	colors := []int{}

	if i < cols-1 {
		colors = append(colors, HouseColorCost[i+1][j])
	}

	if i > 0 {
		colors = append(colors, HouseColorCost[i-1][j])
	}

	if j < rows-1 {

		colors = append(colors, HouseColorCost[i][j+1])
	}

	if j > 0 {
		colors = append(colors, HouseColorCost[i][j-1])
	}

	if i > 0 && j > 0 {
		colors = append(colors, HouseColorCost[i-1][j-1])
	}

	if i < cols-1 && j > 0 {
		colors = append(colors, HouseColorCost[i+1][j-1])
	}

	if i > 0 && j < rows-1 {
		colors = append(colors, HouseColorCost[i-1][j+1])
	}

	if i < cols-1 && j < rows-1 {
		colors = append(colors, HouseColorCost[i+1][j+1])
	}

	var dupes []Point

	for _, c := range colors {
		if c == colorCheck {
			fmt.Println("Found Dupe: ", i, j)
			dupes = append(dupes, Point{i, j, colorCheck})
		}
	}

	return dupes

}

func TestHouseColor(t *testing.T) {

	setup()

}
