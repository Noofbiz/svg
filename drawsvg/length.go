package drawsvg

import (
	"math"
	"strconv"
	"strings"
)

type svgLength struct {
	num  float64
	unit svgLengthUnit
}

type svgLengthUnit uint8

const (
	px svgLengthUnit = iota
	ex
	em
	pt
	pc
	centimeters
	milimeters
	inches
	percent
	unknown
)

func getLengthUnits(s string, t convertType) (svgLengthUnit, int) {
	s = strings.ToLower(s)
	switch true {
	case strings.HasSuffix(s, "em"):
		return em, 2
	case strings.HasSuffix(s, "ex"):
		return ex, 2
	case strings.HasSuffix(s, "px"):
		return px, 2
	case strings.HasSuffix(s, "pt"):
		return pt, 2
	case strings.HasSuffix(s, "pc"):
		return pc, 2
	case strings.HasSuffix(s, "cm"):
		return centimeters, 2
	case strings.HasSuffix(s, "mm"):
		return milimeters, 2
	case strings.HasSuffix(s, "in"):
		return inches, 2
	case strings.HasSuffix(s, "%"):
		return percent, 1
	default:
		if t == width {
			return defaultXUnit, 0
		}
		return defaultYUnit, 0
	}
}

func getSVGLength(cssNumber string, t convertType) (svgLength, error) {
	units, unitLen := getLengthUnits(cssNumber, t)

	f, err := strconv.ParseFloat(cssNumber[0:len(cssNumber)-unitLen], 64)
	if err != nil {
		return svgLength{0.0, defaultXUnit}, err
	}

	return svgLength{f, units}, nil
}

type convertType uint8

const (
	width convertType = iota
	height
	none
)

//convert changes an svgLength into pixels.
func (s svgLength) convert(t convertType) float64 {
	var dpi, p float64
	if t == width {
		dpi = wDPI
		p = userWidth.num
	} else if t == height {
		dpi = hDPI
		p = userHeight.num
	} else {
		dpi = math.Sqrt(wDPI*wDPI + hDPI*hDPI)
		p = math.Sqrt(userWidth.num*userWidth.num + userHeight.num*userHeight.num)
	}
	switch s.unit {
	case em:
		return s.num * fontSize.convert(height)
	case ex:
		return s.num * fontSize.convert(height) / 2
	case px:
		return s.num
	case pc:
		return s.num * 12.0 * (1.0 / 72.0) * dpi
	case pt:
		return s.num * (1.0 / 72.0) * dpi
	case centimeters:
		return s.num * 2.54 * dpi
	case milimeters:
		return s.num * 0.254 * dpi
	case inches:
		return s.num * dpi
	case percent:
		return s.num * (1.0 / 100.0) * p
	}
	return 0
}
