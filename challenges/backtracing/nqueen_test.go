package backtracing

import (
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

/*

The N queens puzzle is the classic backtracking problem. The question is this:

You have an N by N board.
Write a function that returns the number of possible arrangements of the board where N queens
can be placed on the board without threatening each other,
i.e. no two queens share the same row, column, or diagonal.

*/

var N int = 32
var board [][]int = make([][]int, N)

//not used
func setup() {
	//cols
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
	}

	//rows
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = 0
		}
	}

	// let's look at the 1D array (1st row)
	row := board[0:N][0]
	spew.Dump(row)

}

/*
We can represent our board as just a 1D array of integers from 1..N,
where the value at the index i that represents the column the queen on row i is on.
Since we're working incrementally, we don't even need to have the whole board initialized.
We can just append and pop as we go down the stack.
*/

func TestNQueen(t *testing.T) {
	log.Println("Starting test")

	var expectedResultsMap map[int]int = map[int]int{
		0: 1,
		1: 1,
		2: 0,
		3: 0,
		4: 2,
		5: 10,
		6: 4,
		7: 40,
		8: 92,
		9: 352,
	}

	for i := 0; i < 10; i++ {
		board := []int{}
		r := n_queen(i, board)
		if res, ok := expectedResultsMap[i]; !ok || res != r {
			t.Fail()
			t.Errorf("RESULTS[%d]: %d\n", i, r)
			continue
		}
		log.Printf("RESULTS[%d]: %d\n", i, r)
	}
}

// this is called recursively
func n_queen(n int, board []int) int {

	if n == len(board) {
		return 1
	}
	//fmt.Printf("NQUEEN: %d \n", n)
	count := 0
	for i := 0; i < n; i++ {
		board = append(board, i) // push
		//fmt.Printf("Add board: %v col: %v\n", board, i)
		if is_valid(board) {
			//fmt.Printf("BOARD IS VALID!!!\n")
			count += n_queen(n, board)
		}
		board = board[:len(board)-1] // pop (mem leak?)
		//fmt.Printf("Add board: %v\n", board)

	}

	return count
}

// the 1d array represents a diag from 0,0 to N-1,N-1 (note len is the number of elements 1 based)
func is_valid(board []int) bool {
	current_row, current_col := len(board)-1, board[len(board)-1] // the row is the element, the array value is the col
	// Since the last element in the array is the current. We need to compare then 2nd to the last element going backwards.
	for i := len(board) - 2; i >= 0; i-- { // DFS on single array. Looking up the diaganoal from the conext of the last element.
		diff := abs((board[i] - current_col))
		//fmt.Printf("ROW[%d] column-check: %d current-col: %d diff:%d\n", i, board[i], current_col, diff)
		if diff == 0 || diff == (current_row-i) {
			return false
		}
	}
	// above computation is small since we are building the board as we evaluate.
	return true
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

/**
Can we construct a partial solution?

Yes, we can tentatively place queens on the board.
**/

/**
Can we verify if the partial solution is invalid?

Yes, we can check a solution is invalid if two queens threaten each other.
To speed this up, we can assume that all queens already placed so far do not threaten each other,
so we only need to check if the last queen we added attacks any other queen.

Can we verify if the solution is complete?

Yes, we know a solution is complete if all N queens have been placed.

Now that we are confident that we can use backtracking,
let's apply it to this problem.
We'll loop through the first row and try placing a queen in column 0..N, and then the second,
and so on up until N. What differs here from brute force is that we'll be adding the queens
incrementally instead of all at once.

**/
