package challenges_test

import (
	"strconv"
	"testing"
)

const MAX_32BITINT_SIZE int = 2147483647

func reverse(x int) int {

	multiply := 1
	if x < 0 {
		multiply = -1
	}
	x *= multiply
	xb := []byte(strconv.Itoa(x))

	for left, right := 0, len(xb)-1; left < right; left, right = left+1, right-1 {

		if xb[right] == 0 {
			continue
		}

		xb[left], xb[right] = xb[right], xb[left]
	}

	ret, _ := strconv.Atoi(string(xb))
	//fmt.Printf("RET: %v\n", string(xb))
	if ret > MAX_32BITINT_SIZE {
		return 0
	}
	ret *= multiply
	return ret
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"inital", args{123}, 321},
		{"inital", args{-123}, -321},
		{"inital", args{MAX_32BITINT_SIZE + 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
