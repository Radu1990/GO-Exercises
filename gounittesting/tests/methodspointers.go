package tests

// Tip: You can declare a method on non-struct types, too. (eg. float64)

// The Scale method must have a pointer receiver
// to change the Vertex value declared in the main function.
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
