package challenges_test

import (
	"strings"
	"testing"
	"unicode"
)

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/

type ByteArr []byte

func (a ByteArr) reverse() string {
	//https://github.com/golang/go/wiki/SliceTricks
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return string(a)

}

func isPalindrome(s string) bool {

	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	newStr := ""
	for _, v := range s {

		if isLetter(v) || unicode.IsNumber(v) {
			newStr += string(v)
		}
	}

	ba := ByteArr(newStr)
	if newStr == ba.reverse() {
		return true
	}

	return false
}

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"yes", args{"Anna"}, true},
		{"yes", args{"My gym"}, true},
		{"yes", args{"I did, did I?"}, true},
		{"no", args{"race a car"}, false},
		{"no", args{"0P"}, false},
		{"yes", args{"abbʼa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
