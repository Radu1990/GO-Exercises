package tests

import (
	"fmt"
	"image"
)


func computeImage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.At(100, 100).RGBA())
}
