package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	// Define a slice of test cases for Add function
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
	}

	// Run each test case and check if the result matches the expected value.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.x, tt.y); got != tt.want {
				t.Errorf("add(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	// Define a slice of test cases for Subtract function
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"positive", 2, 3, -1},
		{"negative", -2, -3, 1},
	}

	// Run each test case and check if the result matches the expected value.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtract(tt.x, tt.y); got != tt.want {
				t.Errorf("subtract(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	// Define a slice of test cases for Multiply function
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"positive", 2, 3, 6},
		{"negative", -2, -3, 6},
	}

	// Run each test case and check if the result matches the expected value.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.x, tt.y); got != tt.want {
				t.Errorf("multiply(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want float64
	}{
		{"positive", 4, 2, 2.0},
		{"negative", -4, 2, -2.0},
		{"zero_divisor", 0, 1, 0.0},   // test with zero as the divisor
		{"divide_by_zero", 5, 0, 0.0}, // test with division by zero
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.x, tt.y); !almostEqual(got, tt.want) {
				t.Errorf("divide(%d, %d) = %.2f, want %.2f", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}
