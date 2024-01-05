package geoplay

// Add adds the vector to another vector and returns the result.
func (v *Vector) Add(other *Vector) *Vector {
	newV := v.Algebra.NewVector()
	newV.addComponentsInternal(v.Components...)
	newV.addComponentsInternal(other.Components...)
	return newV
}

// Sub subtracts the other vector from this vector and returns the result.
func (v *Vector) Sub(other *Vector) *Vector {
	newV := v.Algebra.NewVector()
	newV.addComponentsInternal(v.Components...)
	for _, c := range other.Components {
		newV.addComponentInternal(c.DupWithNewValue(-c.Value))
	}
	return newV
}

// AddComponent adds a single (or a variadic list) of individual components
// to this vector.  If the basis already exists it just adds
// the value to it.
func (v *Vector) AddComponent(cList ...Component) *Vector {
	newV := v.Dup()
	newV.addComponentsInternal(cList...)
	return newV
}

// Multiply does the full geometric product of this vector with
// another vector and returns the result.
func (v *Vector) Multiply(other *Vector) *Vector {
	newV := v.Algebra.NewVector()
	for _, c1 := range v.Components {
		for _, c2 := range other.Components {
			newComponent := c1.Multiply(c2)
			newV.addComponentInternal(newComponent)
		}
	}

	return newV
}

// Complement returns the complement of the vector - where each
// basis is replaced by the complement basis.
func (v *Vector) Complement() *Vector {
	algebra := v.Algebra
	newV := algebra.NewVectorWithComponents(v.Components...)
	for idx := range newV.Components {
		newV.Components[idx].Basis = algebra.ComplementBasis(newV.Components[idx].Basis)
	}
	return newV
}

// Dot performs the dot product of this vector with another vector.
func (v *Vector) Dot(other *Vector) *Vector {
	return v.Algebra.NewScalar(0.5).Multiply(v.Multiply(other).Add(other.Multiply(v)))
}

// Wedge performs the wedge product of this vector with another vector.
func (v *Vector) Wedge(other *Vector) *Vector {
	return v.Algebra.NewScalar(0.5).Multiply(v.Multiply(other).Sub(other.Multiply(v)))
}
