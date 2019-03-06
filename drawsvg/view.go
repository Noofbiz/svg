package drawsvg

import (
	"errors"
	"strings"
	"unicode"
)

var (
	viewBoxWidth, viewBoxHeight svgLength
	userWidth, userHeight       svgLength
	userOffX, userOffY          svgLength
	defaultXUnit                = unknown
	defaultYUnit                = unknown
	hDPI                        = 90.0
	wDPI                        = 90.0
	fontSize                    = svgLength{num: 10.0, unit: pt}
)

func viewPortSetup(w, h, vb string) error {
	var err error
	if userWidth, err = getSVGLength(w, width); err != nil {
		return err
	}

	if userHeight, err = getSVGLength(h, height); err != nil {
		return err
	}

	defaultXUnit = userHeight.unit
	defaultYUnit = userWidth.unit

	userOffX = svgLength{num: 0.0, unit: defaultXUnit}
	userOffY = svgLength{num: 0.0, unit: defaultYUnit}

	switch defaultXUnit {
	case px:
		wDPI = 90
	case pc:
		wDPI = 6
	case pt:
		wDPI = 72
	case centimeters:
		wDPI = 0.393701
	case milimeters:
		wDPI = 0.0393701
	case inches:
		wDPI = 1
	case unknown:
		wDPI = 90
		defaultXUnit = px
	}

	switch defaultYUnit {
	case px:
		hDPI = 90
	case pc:
		hDPI = 6
	case pt:
		hDPI = 72
	case centimeters:
		hDPI = 0.393701
	case milimeters:
		hDPI = 0.0393701
	case inches:
		hDPI = 1
	case unknown:
		hDPI = 90
		defaultYUnit = px
	}

	if err = setupViewBox(vb); err != nil {
		return err
	}

	return nil
}

func setupViewBox(vb string) error {
	if vb == "" {
		viewBoxWidth = userWidth
		viewBoxHeight = userHeight
		return nil
	}

	vbDim := strings.FieldsFunc(vb, func(r rune) bool {
		if unicode.IsSpace(r) || r == ',' {
			return true
		}
		return false
	})

	x0, err := getSVGLength(vbDim[0], width)
	if err != nil {
		return err
	}

	y0, err := getSVGLength(vbDim[1], height)
	if err != nil {
		return err
	}

	w, err := getSVGLength(vbDim[2], width)
	if err != nil {
		return err
	}

	h, err := getSVGLength(vbDim[3], height)
	if err != nil {
		return err
	}

	if w.num <= 0 || h.num <= 0 {
		return errors.New("malformed svg viewbox")
	}

	viewBoxWidth = svgLength{w.convert(width) - x0.convert(width), px}
	viewBoxHeight = svgLength{h.convert(height) - y0.convert(height), px}

	return nil
}
