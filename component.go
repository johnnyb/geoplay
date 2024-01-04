package geoplay

import "fmt"

// Component is a component of a vector.  It has the value and the Basis elements.
// A scalar is just a Component that has no basis vectors.
type Component struct {
	Algebra *Algebra
	Value   float64
	Basis   []*Basis
}

func (c Component) Multiply(other Component) Component {
	newValue := c.Value * other.Value
	tmpBasis := append(c.Basis, other.Basis...)

	multiplier, newBasis := c.Algebra.SimplifyBasis(tmpBasis)
	if multiplier == -1 {
		newValue = -newValue
	}

	result := Component{
		Algebra: c.Algebra,
		Value:   newValue,
		Basis:   newBasis,
	}

	return result
}

func (c Component) String() string {
	basisStr := ""
	for _, basis := range c.Basis {
		if basisStr != "" {
			basisStr += "^"
		}
		basisStr += basis.Name
	}
	if basisStr == "" {
		return fmt.Sprintf("%f", c.Value)
	}

	return fmt.Sprintf("%f %s", c.Value, basisStr)
}

func (c Component) HasBasis(basis []*Basis) bool {
	if len(c.Basis) != len(basis) {
		return false
	}
	for idx := range c.Basis {
		if c.Basis[idx] != basis[idx] {
			return false
		}
	}
	return true
}
