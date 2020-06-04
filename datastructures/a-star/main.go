package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/as27/gop5js"
)

//Globals
var rows int = 50
var cols int = 50
var calls int = 0

// open and closed sets
var openSet []*Spot
var closedSet []*Spot

// Start and End
var start *Spot
var end *Spot

// width and height
var spotWidth float64 = 0
var spotHeight float64 = 0

//https://golang.org/ref/spec#Assignments
// Addressable is a specific term in the Go language specification meaning whether an address can be accessed for an expression. You can only assign to addressable expressions. Currently map members are not addressable.
// If you use a pointer instead then you can assign via pointer indirection.
// var gScore map[*Spot]int = make(map[*Spot]int) // we have to use the pointer since Spot is not comparable.

//matrix of spots, these are the only allocated spots, everything else refers to these posts.
var grid [][]Spot = make([][]Spot, cols) // this makes the 1st array

// setup the grid and neighbors
func setup() {

	fmt.Println("Starting A* Setup")

	// set up the grid's columns
	for i := 0; i < cols; i++ {
		grid[i] = make([]Spot, rows)
	}

	// populate the grid with spots and make the array
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			grid[i][j] = NewSpot(i, j)
		}
	}

	// add neighbors by pointing to other parts of the grid
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			grid[i][j].AddNeighbors(grid)
		}
	}

	spotWidth = float64(400) / float64(cols)
	spotHeight = float64(400) / float64(rows)

	start = &grid[0][0]
	end = &grid[cols-1][rows-1]
	start.wall = false
	end.wall = false

	openSet = append(openSet, start)
	closedSet = []*Spot{}

}

// this is called in a while loop -- this will implement the while loop for the a-star
func draw() {
	calls++
	var current *Spot
	fmt.Printf("Length of openSet: %d -- called: %d\n", len(openSet), calls)
	if len(openSet) == 0 {
		fmt.Println("NO SOLUTION??")
		gop5js.NoLoop()
		return
	}

	// the direction the path takes
	var winnerIndex int = 0

	// NOTE: not using the for loop format with a range.
	// The range is like a cache copy of the original global
	// Looking for the best next option on where to go
	for i := 0; i < len(openSet); i++ {
		if openSet[i].f < openSet[winnerIndex].f {
			winnerIndex = i
		}
	}

	// if current == goal done (print path)
	current = openSet[winnerIndex] // NOTE: has the lowest fscore
	fmt.Printf("working on openSet (%d, %d): %d  winnerIndex: %d\n", current.i, current.j, len(openSet), winnerIndex)

	closedSet = append(closedSet, current)
	openSet = unsetArr(openSet, winnerIndex)

	// test for are we done?
	if current == end { // comparing the memory address.
		fmt.Println("DONE")
		gop5js.NoLoop()
	}

	for i, neighbor := range current.neighbors {
		// if the item is in the openSet and not in the closedSet
		if ok, _ := in_array(neighbor, closedSet); !ok && neighbor.wall == false {
			tempG := current.g + heuristic(neighbor, current)
			newPath := false
			fmt.Printf("The neighbor has not been visited I: %d J: %d\n", neighbor.i, neighbor.j)
			if ok, _ := in_array(neighbor, openSet); ok {
				fmt.Printf("OPENSET HAS THE NEIGHBOR\n")
				if tempG < neighbor.g {
					current.neighbors[i].g = tempG
					newPath = true
				}

			} else {
				current.neighbors[i].g = tempG
				newPath = true
				fmt.Printf("Adding an openSet: x: %d y: %d\n", neighbor.i, neighbor.j)
				openSet = append(openSet, neighbor)
			}

			if newPath {
				fmt.Println("NEW PATH FOUND:")
				current.neighbors[i].h = heuristic(neighbor, end)
				current.neighbors[i].f = neighbor.g + neighbor.h
				current.neighbors[i].previous = current
			}
		}
	}

	showTheGrid(current)

	//if (current.i == end.i) && (current.j == end.j) { // cannot compare structs in this case since all fields are not comparable.

}

func showTheGrid(current *Spot) {
	// show the grid
	fmt.Println("Showing the grid")
	gop5js.Background("255")

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			grid[i][j].Show("color(255, 255, 255)")
		}
	}

	for i := range closedSet {
		closedSet[i].Show("color(255, 0, 0, 50)") // red
	}

	for i := range openSet {
		openSet[i].Show("color(0, 255, 0, 50)") // green
	}

	// Find the path
	var path []Spot
	var temp Spot = *current

	path = append(path, temp)
	// go back from the previous to the start
	for temp.previous != nil {
		path = append(path, temp)
		temp = *temp.previous
	}
	/*
		for i, p := range path {
			if p.i == end.i && p.j == end.j {
				fmt.Printf("Drawing from the end: (%d, %d)\n", end.i, end.j)
			}
			if p.i == start.i && start.j == p.j {
				fmt.Printf("Drawing to the start: (%d, %d)\n", start.i, start.j)
			}
			path[i].Show("color(0, 0, 255)") // blue
		}*/
	current.Show("color(0,0,255)")

	//start.Show("color(0, 0, 255)")
	//end.Show("color(0, 0, 255)")

	gop5js.NoFill()
	gop5js.Stroke("255, 0, 200")
	gop5js.StrokeWeight((spotWidth / float64(5)))

	gop5js.BeginShape()

	for _, p := range path {
		x := float64(p.i)*(spotWidth) + spotHeight/float64(2)
		y := float64(p.j)*(spotHeight) + spotWidth/float64(2)
		//fmt.Printf("X: %.1f Y: %.1f\n", x, y)
		gop5js.Vertex(float64(x), float64(y))
	}

	gop5js.EndShape()

}

func setIsSet(arr []Spot, index int) bool {
	return (int(len(arr)) > index)
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) { // == true
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func heuristic(a *Spot, b *Spot) float64 {
	p1 := Point{a.j, b.i}
	p2 := Point{a.i, b.j}
	return p1.Distance(p2)
}

func unsetArr(slice []*Spot, s int) []*Spot {
	fmt.Printf("REMOVING: %d LEN: %d\n", s, len(slice))
	return append(slice[:s], slice[s+1:]...)
}

func arrayUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}

func main() {

	setup()
	gop5js.CanvasHeight = 500
	gop5js.CanvasWidth = 500
	gop5js.Draw = draw
	err := gop5js.Serve()
	if err != nil {
		panic(errors.Unwrap(err))
	}
}
