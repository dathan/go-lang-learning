package challenges_test

import (
	"sort"
	"testing"
)

/**
Given an array of time intervals (start, end) for classroom lectures (possibly overlapping),

--- find the minimum number of rooms required. ---

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

*/
type StartStop struct {
	Start int
	Stop  int
}

func TestOverlap(t *testing.T) {

	input := []StartStop{
		{30, 70},
		{0, 50},
		{60, 150},
	}

	max_count := optimal(input)
	t.Logf("Minimum rooms needed: %d\n", max_count)
}

type startstop []StartStop

func (s startstop) Less(i, j int) bool { return s[i].Start < s[j].Start }
func (s startstop) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s startstop) Len() int           { return len(s) }

func optimal(input []StartStop) int {

	/**
	sort the input, then increment the count if start doesn't overlap
	*/
	sort.Sort(startstop(input))
	count := 1
	lastStop := input[0].Stop

	for j := 1; j < len(input); j++ {

		if lastStop < input[j].Start {
			count++
			lastStop = input[j].Stop
		}
	}

	return count
}

func bruteForce(input []StartStop) int {
	max_count := 0
	for i, timelap := range input {
		count := 1
		for j := 0; j < len(input); j++ {

			if i == j {
				continue
			}

			if timelap.Start < input[j].Start &&
				timelap.Start < input[j].Stop &&
				timelap.Stop < input[j].Stop &&
				timelap.Stop < input[j].Start {
				count++
			}

			if count > max_count {
				max_count = count
			}
		}
	}
	return max_count
}
