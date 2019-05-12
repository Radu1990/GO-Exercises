package tests

import "testing"

func TestInterfaceValuesWithNil(t *testing.T) {
	var i J

	var te *U
	i = te
	describeTwo(i)
	i.M()

	i = &U{"hello"}
	describeTwo(i)
	i.M()
}

// Calling a method on a nil interface is a run-time error because
// there is no type inside the interface tuple to indicate which concrete method to call.
// will leave this here BUT commented for the sake of build

/*
func TestNilInterfaceValues(t *testing.T) {
	var i I
	describe(i)
	i.M()
}
*/
