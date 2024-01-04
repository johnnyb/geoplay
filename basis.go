package geoplay

// Basis is the struct that is used to refer to the basis for a vector component.
// A basis should *always* be passed by pointer, as basis comparisons are
// generally done as pointer comparisons, not name comparisons (importing into a different algebra is based on name, however).
type Basis struct {
	Algebra *Algebra
	Name    string
}
