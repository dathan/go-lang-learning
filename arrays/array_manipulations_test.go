package arrays

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"intersect-ok", args{[]int{1, 4, 5, 7, 8, 10}, []int{2, 3, 5, 7}}, []int{5, 7}},
		{"intersect-ok", args{[]int{1, 10}, []int{10, 10, 10}}, []int{10, 10, 10}},
		{"intersect-ok", args{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 4, 4, 5}}, []int{2, 3, 4, 4, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
