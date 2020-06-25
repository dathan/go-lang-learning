package challenges_test

import "testing"

/**
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.


Note: You may assume the string contain only lowercase English letters.
*/
func firstUniqChar(s string) int {
	// counts the number of chars
	var countCharMap map[rune]int = make(map[rune]int)
	for _, sb := range s { // can use two pointers to fill the map up faster
		countCharMap[sb]++
	}

	for i, sb := range s { // can use two pointers to fill the map up faster
		if v, _ := countCharMap[sb]; v == 1 {
			return i
		}
	}

	return -1
}

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{"leetcode"}, 0},
		{"test", args{"loveleetcode"}, 2},
		{"test", args{"aaa"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
