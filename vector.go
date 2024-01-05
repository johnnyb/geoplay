package geoplay

// Vector is the struct used to handle all multicomponent values.
type Vector struct {
	Algebra    *Algebra
	Components []Component
}

func (v *Vector) Dup() *Vector {
	comps := append([]Component{}, v.Components...)
	newV := Vector{
		Algebra:    v.Algebra,
		Components: comps,
	}

	return &newV
}

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

func (v *Vector) Add(other *Vector) *Vector {
	newV := v.Algebra.NewVector()
	newV.addComponentsInternal(v.Components...)
	newV.addComponentsInternal(other.Components...)
	return newV
}

func (v *Vector) Sub(other *Vector) *Vector {
	newV := v.Algebra.NewVector()
	newV.addComponentsInternal(v.Components...)
	for _, c := range other.Components {
		newV.addComponentInternal(c.DupWithNewValue(-c.Value))
	}
	return newV
}

func (v *Vector) AddComponent(cList ...Component) *Vector {
	newV := v.Dup()
	newV.addComponentsInternal(cList...)
	return newV
}

func (v *Vector) ScalarComponent() float64 {
	return v.ComponentForBasis(nil).Value
}

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

func (v *Vector) addComponentsInternal(cList ...Component) {
	for _, c := range cList {
		v.addComponentInternal(c)
	}
}

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

func (v *Vector) Complement() *Vector {
	algebra := v.Algebra
	newV := algebra.NewVectorWithComponents(v.Components...)
	for idx := range newV.Components {
		newV.Components[idx].Basis = algebra.ComplementBasis(newV.Components[idx].Basis)
	}
	return newV
}

func (v *Vector) DotProduct(other *Vector) *Vector {
	return v.Algebra.NewScalar(0.5).Multiply(v.Multiply(other).Add(other.Multiply(v)))
}

func (v *Vector) WedgeProduct(other *Vector) *Vector {
	return v.Algebra.NewScalar(0.5).Multiply(v.Multiply(other).Sub(other.Multiply(v)))
}

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
