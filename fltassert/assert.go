package fltassert

import (
	"fmt"

	"github.com/erinpentecost/fltcmp"
	"github.com/stretchr/testify/assert"
)

// Equal32 can be used in unit testing.
func Equal32(t assert.TestingT, expected, actual float32, ulp uint32, msgAndArgs ...interface{}) bool {
	if !fltcmp.AlmostEqual32(expected, actual, ulp) {

		return assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			"actual  : %v", expected, actual), msgAndArgs...)
	}

	return true

}

// Equal can be used in unit testing.
func Equal(t assert.TestingT, expected, actual float64, ulp uint64, msgAndArgs ...interface{}) bool {
	if !fltcmp.AlmostEqual(expected, actual, ulp) {

		return assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			"actual  : %v", expected, actual), msgAndArgs...)
	}

	return true
}
