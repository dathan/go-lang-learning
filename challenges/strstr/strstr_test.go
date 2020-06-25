package challenges_test

import "testing"

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	haystackSize := len(haystack)
	needleSize := len(needle)

	for i := 0; i < haystackSize; i++ {
		compliment := haystackSize - i
		if compliment >= needleSize {
			found := true
			for j := 0; j < needleSize; j++ {
				if needle[j] != haystack[i+j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}

	return -1

}

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"FAIL", args{"a", "a"}, 0},
		//{"OK", args{"hello", "ll"}, 2},
		//{"OK", args{"aaaaa", "bba"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
