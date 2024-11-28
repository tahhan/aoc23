package day06_2

import "testing"

func TestBoats2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Boats2(); got != tt.want {
				t.Errorf("Boats2() = %v, want %v", got, tt.want)
			}
		})
	}
}
