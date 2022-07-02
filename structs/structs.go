package structs

type Rectangle struct {
	width  float64
	height float64
}

func RectArea(rect Rectangle) (area float64) {
	return rect.height * rect.width
}
func RectPerim(rect Rectangle) (perim float64) {
	return 2 * (rect.height + rect.width)
}
