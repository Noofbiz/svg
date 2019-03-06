package drawsvg

import "github.com/golang/freetype/truetype"

func setupPresentationAttributes(m map[string]string) error {
	for k, v := range m {
		switch k {
		case "fill":
			p, err := getSVGPaint(v, paintTypeFill)
			if err != nil {
				return err
			}
			fillColor = p
		case "stroke":
			p, err := getSVGPaint(v, paintTypeStroke)
			if err != nil {
				return err
			}
			strokeColor = p
		case "stroke-width":
			l, err := getSVGLength(v, none)
			if err != nil {
				return err
			}
			dctx.SetLineWidth(l.convert(none))
		case "font-size":
			if isRelativeOrAbsoluteFontSize(v) {
				//do stuff for relative or absolute
				continue
			}
			l, err := getSVGLength(v, width)
			if err != nil {
				return err
			}
			fontSize = l
		case "font-family":
			fnt, err := getSVGFontFamily(v)
			if err != nil {
				return err
			}
			face := truetype.NewFace(fnt, &truetype.Options{
				Size: fontSize.convert(height),
			})
			dctx.SetFontFace(face)
		}
	}
	return nil
}
