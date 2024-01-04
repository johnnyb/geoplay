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
	fmt.Printf("%+v\n", a)
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", v2)
	fmt.Printf("%v\n", v.Multiply(v2))
	fmt.Printf("%v\n", v2.Multiply(v))
	fmt.Printf("%f\n", v2.Multiply(v).ScalarComponent())
}
