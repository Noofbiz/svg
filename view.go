package svg

type coordinateSystem struct {
	width    svgNum
	height   svgNum
	rotation float64
}

var (
	viewport, user coordinateSystem
)

func (s *svg) setViewport(x, y, width, height, viewBox string) {

}
