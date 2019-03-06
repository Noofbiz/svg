package drawsvg

import "image/color"

type paintType uint8

const (
	paintTypeFill paintType = iota
	paintTypeStroke
)

func getSVGPaint(p string, t paintType) (color.Color, error) {
	var s color.Color
	switch p {
	case "none":
		s = color.Transparent
	case "currentColor", "inherit":
		switch t {
		case paintTypeFill:
			s = fillColor
		case paintTypeStroke:
			s = strokeColor
		}
	default:
		c, err := getSVGColor(p)
		if err != nil {
			return nil, err
		}
		s = c
	}
	return s, nil
}
