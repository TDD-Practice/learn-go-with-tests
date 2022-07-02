package structs

type Line struct {
	point1 float64
	point2 float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (rect Rectangle) Area() (area float64) {
	return rect.height * rect.width
}
func (rect Rectangle) Perim() (perim float64) {
	return 2 * (rect.height + rect.width)
}

func (circ Circle) Perim() (area float64) {
	return 3.1416 * 2 * circ.radius
}

func (circ Circle) Area() (area float64) {
	return 3.1416 * circ.radius * circ.radius
}
