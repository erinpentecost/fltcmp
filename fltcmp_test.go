package fltcmp_test

import (
	"fmt"
	"testing"

	"github.com/erinpentecost/fltcmp"
	"github.com/erinpentecost/fltcmp/fltassert"
	"github.com/stretchr/testify/assert"
)

type Caser interface {
	Test(t *testing.T)
}

type Case32 struct {
	a float32
	b float32
	u uint32
	e bool
}

func (c Case32) Test(t *testing.T) {
	if c.e {
		assert.True(t, fltcmp.AlmostEqual32(c.a, c.b, c.u), fmt.Sprintf("%v != %v with ulp %v", c.a, c.b, c.u))
	} else {
		assert.True(t, fltcmp.AlmostEqual32(c.a, c.b, c.u), fmt.Sprintf("%v != %v with ulp %v", c.a, c.b, c.u))
	}
}

type Case64 struct {
	a float64
	b float64
	u uint64
	e bool
}

func (c Case64) Test(t *testing.T) {
	if c.e {
		assert.True(t, fltcmp.AlmostEqual(c.a, c.b, c.u), fmt.Sprintf("%v != %v with ulp %v", c.a, c.b, c.u))
	} else {
		assert.False(t, fltcmp.AlmostEqual(c.a, c.b, c.u), fmt.Sprintf("%v == %v with ulp %v", c.a, c.b, c.u))
	}
}

func TestCases(t *testing.T) {
	// I need way more tests.
	cases := []Caser{
		Case32{2.0, 2.0, 1, true},
		Case64{2.0, 2.0, 1, true},
		Case32{0.0, -0.0, 1, true},
		Case64{0.0, -0.0, 1, true},
		Case32{1e-450, -1e-450, 2, true},
		Case64{1e-450, -1e-450, 2, true},
	}

	for _, c := range cases {
		c.Test(t)
	}
}

func TestDiff(t *testing.T) {
	nearZero := 8.881784197001252e-16
	distance := fltcmp.FloatDiff(nearZero, 0.0)
	assert.Equal(t, uint64(4382002437431492608), distance)
}

func TestAssert(t *testing.T) {
	fltassert.Equal32(t, 394769.0, 394769.0, 1, "assertion32")
	fltassert.Equal(t, 394769.0, 394769.0, 1, "assertion64")
}
