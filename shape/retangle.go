package shape

type Rectangle struct {
	Width  float32
	Length float32
}

func (retangle *Rectangle) Area() float32 {
	return retangle.Width * retangle.Length
}

func (retangle *Rectangle) Perimeter() float32 {
	return 2 * (retangle.Width + retangle.Length)
}
