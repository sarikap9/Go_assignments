package main

import (
	"testing"
)

func TestGenDisplaceFn(t *testing.T) {
	fn := GenDisplaceFn(10, 2, 1)

	tests := []struct {
		time     float64
		expected float64
	}{
		{3, 52.0},
		{5, 136.0},
		{0, 1.0},
	}

	for _, tt := range tests {
		result := fn(tt.time)
		if result != tt.expected {
			t.Errorf("For time %f, expected %f but got %f", tt.time, tt.expected, result)
		}
	}
}
