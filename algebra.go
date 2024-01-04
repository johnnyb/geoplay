package geoplay

import "fmt"

// Algebra is the struct that holds the fundamental list of basis vectors.
type Algebra struct {
	Basis              []*Basis
	BasisPositionCache map[*Basis]int
}

func NewAlgebraWithNumDimensions(numDimensions int) *Algebra {
	dimensionNames := []string{}
	for i := 0; i < numDimensions; i++ {
		dimensionNames = append(dimensionNames, fmt.Sprintf("e%d", i))
	}
	return NewAlgebra(dimensionNames...)
}

// NewAlgebra creates a new Clifford Algebra with the given dimension names.
func NewAlgebra(dimensions ...string) *Algebra {
	algebra := &Algebra{
		Basis: []*Basis{},
	}
	for _, dim := range dimensions {
		basis := &Basis{
			Algebra: algebra,
			Name:    dim,
		}
		algebra.Basis = append(algebra.Basis, basis)
	}
	algebra.cacheBasisPositions()
	return algebra
}

func (algebra *Algebra) cacheBasisPositions() {
	positions := map[*Basis]int{}
	for idx, basis := range algebra.Basis {
		positions[basis] = idx
	}
	algebra.BasisPositionCache = positions
}

// NewVector creates a new, simple vector with the given component values (assumed to align with the order of the basis vectors in the algebra).
func (algebra *Algebra) NewVector(componentValues ...float64) *Vector {
	vec := &Vector{
		Algebra:    algebra,
		Components: []Component{},
	}
	for i := 0; i < len(componentValues); i++ {
		basis := algebra.Basis[i]
		component := Component{
			Algebra: algebra,
			Value:   componentValues[i],
			Basis:   []*Basis{basis},
		}
		vec.Components = append(vec.Components, component)
	}

	return vec
}

func (algebra *Algebra) SimplifyBasis(basis []*Basis) (int, []*Basis) {
	multiplier, basis := algebra.SortBasis(basis)
	for i := 0; i < len(basis)-1; i++ {
		if basis[i] == basis[i+1] {
			if i == 0 {
				basis = basis[i+2:]
			} else {
				basis = append(basis[:i], basis[i+2:]...)
			}
			i = i - 1 // Reset i so it gets incremented back to this position
		}
	}

	return multiplier, basis
}

func (algebra *Algebra) GetBasisNamed(name string) *Basis {
	for _, basis := range algebra.Basis {
		if basis.Name == name {
			return basis
		}
	}
	return nil
}

func (algebra *Algebra) SortBasis(basis []*Basis) (int, []*Basis) {
	var multiplier int = 1

	// Get a new copy of the array
	newBasis := append([]*Basis{}, basis...)

	// Swap positions until they are in order
	for i := 0; i < len(newBasis)-1; i++ {
		b1 := newBasis[i]
		b2 := newBasis[i+1]
		if algebra.BasisPositionCache[b1] > algebra.BasisPositionCache[b2] {
			newBasis[i+1] = b1
			newBasis[i] = b2
			multiplier = -multiplier

			i -= 2
			if i < -1 {
				i = -1
			}
		}
	}

	return multiplier, newBasis
}
