package challenges_test

import (
	"fmt"
	"testing"
)

/*
Matching Pairs
Given two strings s and t of length N, find the maximum number of possible matching pairs in strings s and t after swapping exactly two characters within s.
A swap is switching s[i] and s[j], where s[i] and s[j] denotes the character that is present at the ith and jth index of s, respectively. The matching pairs of the two strings are defined as the number of indices for which s[i] and t[i] are equal.
Note: This means you must swap two characters at different indices.
Signature
int matchingPairs(String s, String t)
Input
s and t are strings of length N
N is between 2 and 1,000,000
Output
Return an integer denoting the maximum number of matching pairs
Example 1
s = "abcd"
t = "adcb"
output = 4
Explanation:
Using 0-based indexing, and with i = 1 and j = 3, s[1] and s[3] can be swapped, making it  "adcb".
Therefore, the number of matching pairs of s and t will be 4.
Example 2
s = "mno"
t = "mno"
output = 1
Explanation:
Two indices have to be swapped, regardless of which two it is, only one letter will remain the same. If i = 0 and j=1, s[0] and s[1] are swapped, making s = "nmo", which shares only "o" with t
*/
func matchingPairs(s string, t string) int {
	// some notes s & t are the same length

	sb := []byte(s)
	tb := []byte(t)
	swapped := false
	for i := 0; i < len(sb); i++ {

		if sb[i] != tb[i] {
			fmt.Printf("%d) SB: %s != TB: %s\n", i, string(sb[i]), string(tb[i]))
			for j := i + 1; j < len(s); j++ {
				if sb[j] == tb[i] {
					fmt.Printf("Swaping sb[j]: %s with sb[i] %s\n", string(sb[j]), string(sb[i]))
					sb[j], sb[i] = sb[i], sb[j]
					fmt.Printf("SWAPED: %s\n", string(sb))
					swapped = true
					break
				}
			}
		}
	}
	if !swapped {
		fmt.Printf("NO SWAP swaping 1st two chars\n")
		sb[0], sb[1] = sb[1], sb[0]
	}

	pairCounter := 0
	for i := 0; i < len(sb); i++ {
		if sb[i] == tb[i] {
			pairCounter++
		}
	}
	fmt.Printf("SB (string): %s PC: %d\n", string(sb), pairCounter)
	return pairCounter
}

func Test_matchingPairs(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"testok", args{"abcd", "adcb"}, 4},
		{"testok", args{"mno", "mno"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchingPairs(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("matchingPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
