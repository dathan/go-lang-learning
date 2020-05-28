package main

import (
	"fmt"

	"github.com/as27/gop5js"
)

// a grid is a matrix representing the board.
// the grid contains spot objects that represent the state of that element of the grid
// Spots can have an obsticle and the algorithm has to route around it
// We are doing a best first search where the algorithm

//The Spot on the grid
type Spot struct {
	x         int64
	y         int64
	h         int64       // hueristic
	g         int64       // cost
	f         int64       // f(n) = g(n) + h(n) where g(n) is the cost and h(n) is the hueristic and n is the next node
	wall      bool        // if this spot is blocked
	neighbors []Spot      // neighbors to this spot alos have spots
	previous  interface{} // TODO: Revisit
}

//Globals
var rows int64 = 50
var cols int64 = 50
var spotWidth float64 = 0
var spotHeight float64 = 0
var start = 0
var end = 0

//matrix of spots
var grid [][]Spot = make([][]Spot, rows) // this makes the 1st array

// return a new spot
func NewSpot(i, j int64) Spot {
	fmt.Printf("NewSpot(%f,%f)\n", i, j)
	s := &Spot{}
	s.x = i
	s.y = j

	return *s

}

func (s *Spot) Show(color string) {
	gop5js.Fill(color)

	x := float64(s.x)
	y := float64(s.y)

	if s.wall {
		gop5js.Fill("0")
		gop5js.NoStroke() // https://p5js.org/reference/#/p5/noStroke
		gop5js.Ellipse(x*spotWidth+x/2, y*spotHeight+spotHeight/2, spotWidth/2, spotHeight/2)
	} else {
		//fmt.Printf("X: %d Y: %d\n", x, y)
		gop5js.Rect(x*spotWidth, y*spotHeight, spotWidth-1, spotHeight-1)
	}
}

func (s *Spot) AddNeighbors(grid [][]Spot) {
	var i = s.x
	var j = s.y

	if i < cols-1 {
		s.neighbors = append(s.neighbors, grid[i+1][j])
	}

	if i > 0 {
		s.neighbors = append(s.neighbors, grid[i-1][j])
	}

	if j < rows-1 {
		s.neighbors = append(s.neighbors, grid[i][j+1])
	}

	if j > 0 {
		s.neighbors = append(s.neighbors, grid[i][j-1])
	}

	if i > 0 && j > 0 {
		s.neighbors = append(s.neighbors, grid[i-1][j-1])
	}

	if i < cols-1 && j > 0 {
		s.neighbors = append(s.neighbors, grid[i+1][j-1])
	}

	if i > 0 && j < rows-1 {
		s.neighbors = append(s.neighbors, grid[i-1][j+1])
	}

	if i < cols-1 && j < rows-1 {
		s.neighbors = append(s.neighbors, grid[i+1][j+1])
	}

}

func main() {
	setup()
	gop5js.CanvasHeight = 500
	gop5js.CanvasWidth = 500
	gop5js.Draw = draw
	gop5js.Serve()
}

// setup the grid and neighbors
func setup() {

	for i := int64(0); i < cols; i++ {
		grid[i] = make([]Spot, rows)
	}
	// populate the grid with spots and make the array
	for i := int64(0); i < cols; i++ {
		for j := int64(0); j < rows; j++ {
			grid[i][j] = NewSpot(i, j)
		}
	}

	for i := int64(0); i < cols; i++ {
		for j := int64(0); j < rows; j++ {
			grid[i][j].AddNeighbors(grid)
		}
	}

	spotWidth = float64(gop5js.CanvasWidth) / float64(cols)
	spotHeight = float64(gop5js.CanvasHeight) / float64(rows)

}

// this is called in a while loop
func draw() {
	gop5js.Background("127")

	for i := int64(0); i < cols; i++ {
		for j := int64(0); j < rows; j++ {
			grid[i][j].Show("255")
		}
	}
	//gop5js.NoLoop()
	return

	/* 	for i := 0; i < len(closedSet); i++ {
	   		closedSet[i].Show("255,0,0")
	   	}

	   	for i := 0; i < len(openSet); i++ {
	   		openSet[i].Show("0, 255, 0")
	   	}
	*/
}
