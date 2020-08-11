package backtracing

import (
	"errors"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestFlightPath(t *testing.T) {

	spew.Dump(backtraceFlights())
	spew.Dump(nonbacktrace())

}

type SrcDst struct {
	src string
	dst string
}

func backtraceFlights() []string {

	var path []SrcDst = []SrcDst{
		{"HNL", "AKL"},
		{"YUL", "ORD"},
		{"ORD", "SFO"},
		{"SFO", "HNL"},
	}
	// build the path as you go
	// make sure the path is valid call backtraceFlights
	return get_itinerary(path, []string{"YUL"})
}

func get_itinerary(flights []SrcDst, current_path []string) []string {

	if len(flights) == 0 {
		return current_path
	}
	last_stop := current_path[len(current_path)-1]
	for i, path := range flights {
		//i is the number of elements from 0
		flights_minus_current := append([]SrcDst{}, flights[:i]...)
		//from i+1 to the length of the array append
		flights_minus_current = append(flights_minus_current, flights[i+1:]...)

		current_path = append(current_path, path.dst)
		if path.src == last_stop {
			return get_itinerary(flights_minus_current, current_path)
		}

		_, current_path = current_path[len(current_path)-1], current_path[:len(current_path)-1]
	}

	return []string{}
}

func nonbacktrace() (error, []string) {
	var startDestPair map[string]string = map[string]string{
		"HNL": "AKL",
		"YUL": "ORD",
		"ORD": "SFO",
		"SFO": "HNL",
	}

	startingPoint := "YUL"
	origStartPoint := startingPoint
	endPoint := "AKL"

	var path []string = []string{startingPoint}
	var ok bool = true
	var dst string = ""

	for {

		// validate
		dst, ok = startDestPair[startingPoint]
		if !ok {
			path = []string{"NOT POSSIBLE"}
			return errors.New("INVALID DESTINATION"), path
		}

		if dst == origStartPoint {
			path = []string{"NOT POSSIBLE"}
			return errors.New("CYCLE DETECTED"), path
		}

		startingPoint = dst
		path = append(path, dst)

		if dst == endPoint {
			break
		}

	}

	return nil, path

}
