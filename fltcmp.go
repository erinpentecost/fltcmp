// Package fltcmp implements ULP float comparison.
// https://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/
package fltcmp

import (
	"math"
)

// FloatDiff finds the ULP difference between two floats.
func FloatDiff(a, b float64) uint64 {
	// Exact case short-circuit.
	if a == b {
		return 0
	}

	// Convert to ints
	uA := math.Float64bits(a)
	uB := math.Float64bits(b)

	// Fail when signs are different?
	if signbit64(uA) != signbit64(uB) {
		// Just try again after shifting everything.
		return FloatDiff(math.Abs(a)+math.Abs(b), 0.0)
	}

	// Find difference in ULPs.
	return diff64(uA, uB)
}

// AlmostEqual tells you how close two floats are.
// Make maxUlpsDiff 1 if they need be really, really close.
// Returns true if the ULP is within the threshold.
func AlmostEqual(a, b float64, maxUlpsDiff uint64) bool {
	return FloatDiff(a, b) <= maxUlpsDiff
}

func diff64(a, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}

func signbit64(x uint64) bool {
	return x&(1<<63) != 0
}

// FloatDiff32 finds the ULP difference between two floats.
func FloatDiff32(a, b float32) uint32 {
	// Exact case short-circuit.
	if a == b {
		return 0
	}

	// Convert to ints
	uA := math.Float32bits(a)
	uB := math.Float32bits(b)

	// Fail when signs are different?
	if signbit32(uA) != signbit32(uB) {
		// Just try again after shifting everything.
		return FloatDiff32(abs32(a)+abs32(b), 0.0)
	}

	// Find difference in ULPs.
	return diff32(uA, uB)
}

// AlmostEqual32 tells you how close two floats are.
// Make maxUlpsDiff 1 if they need be really, really close.
// Returns true if the ULP is within the threshold.
func AlmostEqual32(a, b float32, maxUlpsDiff uint32) bool {
	return FloatDiff32(a, b) <= maxUlpsDiff
}

func diff32(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func signbit32(x uint32) bool {
	return x&(1<<31) != 0
}

func abs32(x float32) float32 {
	if x >= 0.0 {
		return x
	}
	return x * (-1.0)
}
