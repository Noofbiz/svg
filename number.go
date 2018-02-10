package svg

type svgNum struct {
	num  int
	unit cssUnit
}

type cssUnit int

const (
	px cssUnit = iota
	inches
	centimeters
	milimeters
)
