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

// Antiwedge performs the antiwedge (dot) product of this vector with another vector.
func (v *Vector) Antiwedge(other *Vector) *Vector {
	return v.Algebra.NewScalar(0.5).Multiply(v.Multiply(other).Add(other.Multiply(v)))
}

// Dot is another name for the antiwedge product
func (v *Vector) Dot(other *Vector) *Vector {
	return v.Antiwedge(other)
}

// Wedge performs the wedge product of this vector with another vector.
func (v *Vector) Wedge(other *Vector) *Vector {
	return v.Algebra.NewScalar(0.5).Multiply(v.Multiply(other).Sub(other.Multiply(v)))
}

// IsSimple tells whether a vector is a simple 1-vector.
func (v *Vector) IsSimple() bool {
	for _, c := range v.Components {
		if len(c.Basis) != 1 {
			return false
		}
	}
	return true
}

// Invert inverts the vector if it can tell that it is invertible.
// Currently only inverts simple vectors.  Returns nil if it
// doesn't know how to invert the vector.
func (v *Vector) Invert() *Vector {
	if v.IsSimple() {
		sumSquares := 0.0
		for _, c := range v.Components {
			sumSquares += c.Value * c.Value
		}
		return v.Algebra.NewScalar(1.0 / sumSquares).Multiply(v)
	} else {
		return nil
	}
}

// Divide inverts the second vector and then multiplies
// by the inverse.
func (v *Vector) Divide(other *Vector) *Vector {
	return v.Multiply(other.Invert())
}
