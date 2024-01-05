package geoplay

import (
	"fmt"
	"testing"
)

func TestSimpleVector(t *testing.T) {
	// Not really a test per se, just getting something to run
	a := NewAlgebraWithNumDimensions(3)
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
}
