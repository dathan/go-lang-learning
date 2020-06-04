package main

import (
	"math"
	"math/rand"

	"github.com/as27/gop5js"
)

// a point definition
// Point represents a point in space.
type Point struct {
	X int
	Y int
}

// New returns a Point based on X and Y positions on a graph.
func New(x int, y int) Point {
	return Point{x, y}
}

// Distance finds the length of the hypotenuse between two points.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}

// a grid is a matrix representing the board.
// the grid contains spot objects that represent the state of that element of the grid
// Spots can have an obsticle and the algorithm has to route around it
// We are doing a best first search where the algorithm

//The Spot on the grid
type Spot struct {
	i         int
	j         int
	h         float64 // hueristic
	g         float64 // cost
	f         float64 // f(n) = g(n) + h(n) where g(n) is the cost and h(n) is the hueristic and n is the next node
	wall      bool    // if this spot is blocked
	neighbors []*Spot // neighbors to this spot alos have spots
	previous  *Spot   // TODO: Revisit
}

// return a new spot
func NewSpot(i, j int) Spot {
	s := &Spot{}
	s.i = i
	s.j = j
	s.f = 0
	s.g = 0
	s.h = 0
	s.neighbors = []*Spot{}
	s.previous = nil
	s.wall = false
	if rand.Float64() < 0.2 {
		s.wall = true
	}

	return *s

}

// Show the color
func (s *Spot) Show(color string) {

	i := float64(s.i)
	j := float64(s.j)

	if s.wall {
		gop5js.Fill("0")
		gop5js.NoStroke() // https://p5js.org/reference/#/p5/noStroke
		// put a circle (ellipse) in this grid position
		gop5js.Ellipse(i*spotWidth+spotWidth/2, j*spotHeight+spotHeight/2, spotWidth/2, spotHeight/2)
		return
	}

	gop5js.Fill(color)
	gop5js.Rect(i*spotWidth, j*spotHeight, spotWidth, spotHeight)
}

// add the neighbors spot so each point knows about what its connected to
func (s *Spot) AddNeighbors(grid [][]Spot) {

	var i = s.i
	var j = s.j

	// note since the grid address is assigned to the objects neighbors a copy is avoided enabling a single loop to create it.
	if i < cols-1 {
		s.neighbors = append(s.neighbors, &grid[i+1][j])
	}

	if i > 0 {
		s.neighbors = append(s.neighbors, &grid[i-1][j])
	}

	if j < rows-1 {
		s.neighbors = append(s.neighbors, &grid[i][j+1])
	}

	if j > 0 {
		s.neighbors = append(s.neighbors, &grid[i][j-1])
	}

	if i > 0 && j > 0 {
		s.neighbors = append(s.neighbors, &grid[i-1][j-1])
	}

	if i < cols-1 && j > 0 {
		s.neighbors = append(s.neighbors, &grid[i+1][j-1])
	}

	if i > 0 && j < rows-1 {
		s.neighbors = append(s.neighbors, &grid[i-1][j+1])
	}

	if i < cols-1 && j < rows-1 {
		s.neighbors = append(s.neighbors, &grid[i+1][j+1])
	}

}
