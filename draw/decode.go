package draw

import (
	"errors"
	"image"
	"io"

	"github.com/Noofbiz/svg"
)

// Decode takes in an svg file and outputs an image.Image
// scaleX and scaleY are the scale you want the returned image to be from the
// original svg file. If either one is zero or negative, it'll scale both axes
// based on the positive factor. If both are zero or negative, it'll return an
// error.
func Decode(r io.Reader, scaleX, scaleY float64) (image.Image, error) {
	var err error

	if scaleX <= 0 && scaleY <= 0 {
		return nil, errors.New("one scale needs to be greater than zero")
	}

	if scaleX <= 0 {
		scaleX = scaleY
	} else if scaleY <= 0 {
		scaleY = scaleX
	}

	var cmd []svg.Command
	if cmd, err = svg.ParseSVG(r); err != nil {
		return nil, err
	}

	var img image.Image
	if img, err = drawSVG(cmd, scaleX, scaleY); err != nil {
		return nil, err
	}

	return img, nil
}
