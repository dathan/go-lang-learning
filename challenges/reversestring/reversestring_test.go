package challenges_test

import (
	"testing"
)

// a is a byte slice which is passwd by references. The string is read only and cannot be passed. TODO look to why
func reverse(a []byte) {
	// I like this syntax as it makes sense. the left,right in the left most part of the for loop is cool
	//https://github.com/golang/go/wiki/SliceTricks
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		str []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"try", args{[]byte("hello")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse([]byte(tt.args.str))
			t.Logf("RESULTS: %s\n", string(tt.args.str))
		})
	}
}
