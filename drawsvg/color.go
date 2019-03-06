package drawsvg

import (
	"errors"
	"image/color"
	"strings"

	"golang.org/x/image/colornames"
)

func getSVGColor(c string) (color.Color, error) {
	out, ok := colornames.Map[c]
	if ok {
		return out, nil
	}
	if strings.HasPrefix(c, "#") {
		return hexColor(c), nil
	} else if strings.HasPrefix(c, "rgb") {
		return rgbColor(c), nil
	} else {
		return nil, errors.New("provided string " + c + " is not a recognized color")
	}
}

// TODO: make these work :P
func hexColor(c string) color.Color {
	return color.Black
}

func rgbColor(c string) color.Color {
	return color.Black
}
