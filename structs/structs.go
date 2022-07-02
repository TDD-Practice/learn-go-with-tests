package structs

type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) Area() (area float64) {
	return rect.height * rect.width
}
func (rect Rectangle) Perim() (perim float64) {
	return 2 * (rect.height + rect.width)
}
