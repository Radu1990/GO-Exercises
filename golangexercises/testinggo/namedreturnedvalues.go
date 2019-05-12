package testinggo

// Split is a sneacky method
func Split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x

	return x, y
}
