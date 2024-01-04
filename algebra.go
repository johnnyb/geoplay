package geoplay

import "fmt"

type Algebra struct {
	Basis []*Basis
}

func NewAlgebra(dimensions int) *Algebra {
	algebra := &Algebra{
		Basis: []*Basis{},
	}
	for i := 0; i < dimensions; i++ {
		basis := &Basis{Name: fmt.Sprintf("e%d", i)}
		algebra.Basis = append(algebra.Basis, basis)
	}
	return algebra
}

func (algebra *Algebra) NewVector(componentValues ...float64) *Vector {
	vec := &Vector{
		Algebra:    algebra,
		Components: []Component{},
	}
	for i := 0; i < len(componentValues); i++ {
		basis := algebra.Basis[i]
		component := Component{
			Value: componentValues[i],
			Basis: []*Basis{basis},
		}
		vec.Components = append(vec.Components, component)
	}

	return vec
}
