package geoplay

// Vector is the struct used to handle all multicomponent values.
type Vector struct {
	Algebra    *Algebra    // The Algebra that the Vector is a part of
	Components []Component // The Components of the Vector
}

// Dup duplicates the vector.
func (v *Vector) Dup() *Vector {
	comps := append([]Component{}, v.Components...)
	newV := Vector{
		Algebra:    v.Algebra,
		Components: comps,
	}

	return &newV
}

// ScalarComponent is a helper function extracts the scalar
// part of a vector.
func (v *Vector) ScalarComponent() float64 {
	return v.ComponentForBasis(nil).Value
}

// ComponentForBasis extracts the given component from the
// Vector.  It returns a Component with a Value of 0 if none
// is found within the Vector.
func (v *Vector) ComponentForBasis(basis []*Basis) Component {
	for _, comp := range v.Components {
		if comp.HasBasis(basis) {
			return comp
		}
	}

	return Component{
		Algebra: v.Algebra,
		Value:   0,
		Basis:   basis,
	}
}

// addComponentsInternal is a helper function which performs
// addComponentInternal on a variadic list.
func (v *Vector) addComponentsInternal(cList ...Component) {
	for _, c := range cList {
		v.addComponentInternal(c)
	}
}

// addComponentInternal adds a given component to the vector,
// but *modifies* the vector itself.  This is only really valid
// when initially creating a vector internal to the system before
// sending the result outside of the system.
func (v *Vector) addComponentInternal(c Component) {
	for idx, comp := range v.Components {
		if comp.HasBasis(c.Basis) {
			v.Components[idx] = Component{
				Algebra: comp.Algebra,
				Value:   comp.Value + c.Value,
				Basis:   comp.Basis,
			}
			if v.Components[idx].Value == 0 {
				if idx == 0 {
					v.Components = v.Components[1:]
				} else {
					v.Components = append(v.Components[0:idx], v.Components[idx+1:]...)
				}
			}
			return
		}
	}
	v.Components = append(v.Components, c)
}

// String yields the string representation of the vector.
func (v *Vector) String() string {
	str := ""
	for _, comp := range v.Components {
		if str != "" {
			str += " + "
		}
		str += comp.String()
	}
	return str
}
