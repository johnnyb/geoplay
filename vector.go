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

func (v *Vector) AddComponent(c Component) *Vector {
	newV := v.Dup()
	newV.addComponentInternal(c)
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

func (v *Vector) addComponentInternal(c Component) {
	for idx, comp := range v.Components {
		if comp.HasBasis(c.Basis) {
			v.Components[idx] = Component{
				Algebra: comp.Algebra,
				Value:   comp.Value + c.Value,
				Basis:   comp.Basis,
			}
			return
		}
	}
	v.Components = append(v.Components, c)
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
