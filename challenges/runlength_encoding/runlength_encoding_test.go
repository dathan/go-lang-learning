package challenges_test

import (
	"strconv"
	"testing"
)

/*
 * Run-length encoding is a fast and simple method of encoding strings.
 The basic idea is to represent repeated successive characters as a single count and character.
 For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".

Implement run-length encoding and decoding.
You can assume the string to be encoded have no digits and consists solely of alphabetic characters.
You can assume the string to be decoded is valid.
*/
func RunLengthEncoding(input string) string {

	ret := ""
	var strCounterMap map[string]int = make(map[string]int)
	var prevLetter string
	for i := 0; i < len(input); i++ {
		if _, ok := strCounterMap[string(input[i])]; !ok {
			if prevLetter != "" {
				ret += strconv.Itoa(strCounterMap[prevLetter]) + prevLetter
				strCounterMap = make(map[string]int)
			}
			prevLetter = string(input[i])

		}
		strCounterMap[string(input[i])]++
	}

	// handle the last element
	ret += strconv.Itoa(strCounterMap[prevLetter]) + prevLetter

	return ret
}

func TestRunLengthEncoding(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"encoded", args{input: "AAAABBBCCDAA"}, "4A3B2C1D2A"},
		{"encoded", args{input: "AAAABBBCCDAA"}, "4A3B2C1D2A"},
		{"encoded", args{input: "AAAABBBCCDAAEEEGGC"}, "4A3B2C1D2A3E2G1C"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunLengthEncoding(tt.args.input); got != tt.want {
				t.Errorf("RunLengthEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
