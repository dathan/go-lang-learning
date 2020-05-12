package arrays

import "testing"

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "slice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceExample()
		})
	}
}
