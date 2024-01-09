package geoplay

import (
	"fmt"
	"testing"
)

func Test2DVector(t *testing.T) {
	a := NewAlgebraWithNumDimensions(2)
	a.SetZeroCutoff(0.0000000000000001)
	v := a.NewVector(2, 3)
	v2 := a.NewVector(5, 7)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%v\n", v.Multiply(v2))
}

func TestImaginary(t *testing.T) {
	a := NewAlgebraClifford(0, 1)
	a.SetZeroCutoff(0.0000000000000001)
	v := a.NewVector(2).Add(a.NewScalar(3))
	v2 := a.NewVector(5).Add(a.NewScalar(7))
	v3 := v.Multiply(v2)
	fmt.Printf("%v\n%v\n%v\n", v, v2, v3)
}

func TestSimpleVector(t *testing.T) {
	// Not really a test per se, just getting something to run
	a := NewAlgebraWithNumDimensions(3)
	a.SetZeroCutoff(0.0000000000000001)
	v := a.NewVector(1, 2, 3)
	v2 := a.NewVector(4, 5, 6)
	newV := v.Multiply(v2)
	dp := v.Dot(v2)
	wp := v.Wedge(v2)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", v2)
	fmt.Printf("%v\n", v2.Complement())
	fmt.Printf("%v\n", newV)
	fmt.Printf("%v\n", dp)
	fmt.Printf("%v\n", wp)
	fmt.Printf("%v\n", v2.Multiply(v))
	fmt.Printf("%f\n", v2.Multiply(v).ScalarComponent())
	fmt.Printf("%v\n", newV.Multiply(a.NewScalar(3)))
	fmt.Printf("%v\n", v2.Invert())
	fmt.Printf("%v\n", v.Multiply(v2).Divide(v2))
}
