package challenges_test

import (
	"testing"
)

/**
a word, phrase, or name formed by rearranging the letters of another, such as cinema, formed from iceman.
*/

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	if s == "" {
		return true
	}

	var uniqSCounter map[byte]int = make(map[byte]int)
	var uniqTCounter map[byte]int = make(map[byte]int)
	for left := 0; left < len(s); left++ {
		uniqSCounter[s[left]]++
		uniqTCounter[t[left]]++
	}

	for i := 0; i < len(t); i++ {
		v := t[i]
		if _, ok := uniqSCounter[v]; !ok || uniqSCounter[v] != uniqTCounter[v] {
			return false
		}
	}

	return true

}

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{"anagram", "nagaram"}, true},
		{"true", args{"ab", "ba"}, true},
		{"false", args{"rat", "car"}, false},
		{"true", args{"cinema", "iceman"}, true},
		{"acbb", args{"ac", "bb"}, false},
		{"nlcx", args{"nl", "cx"}, false},
		{"ab", args{"a", "b"}, false},
		{"aacc:ccac", args{"aacc", "ccac"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
