package drawsvg

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"strings"
	"time"

	"github.com/Noofbiz/svg"
	"github.com/fogleman/gg"
)

var (
	dctx                   *gg.Context
	fillColor, strokeColor color.Color
	cmdCh                  chan svg.Command
	errCh                  chan error
	definitions            = make(map[string][]svg.Command)
)

func drawSVG(scaleX, scaleY float64) (image.Image, error) {
	initCmd := <-cmdCh
	initStyle := initCmd.Style
	viewPortSetup(initStyle["width"], initStyle["height"], initStyle["viewBox"])

	dctx = gg.NewContext(int(viewBoxWidth.convert(width)*scaleX), int(viewBoxHeight.convert(width)*scaleY))
	dctx.Scale(scaleX, scaleY)
	for {
		select {
		case cmd, ok := <-cmdCh:
			if !ok {
				return dctx.Image(), nil
			}
			setDefaults()
			err := drawFuncs[cmd.C](cmd.Style)
			if err != nil {
				return nil, err
			}
		case err := <-errCh:
			return nil, err
		case <-time.After(500 * time.Millisecond):
			return dctx.Image(), nil
		}
	}
}

var drawFuncs = make(map[svg.Cmd]func(map[string]string) error)

func init() {
	drawFuncs[svg.SVG] = func(m map[string]string) error {
		// TODO: handle for multiple nested svg
		return nil
	}
	drawFuncs[svg.A] = func(m map[string]string) error {
		// a Link would do nothing in an image.Image
		return nil
	}
	drawFuncs[svg.ALT_GLYPH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.ALT_GLYPH_DEF] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.ALT_GLYPH_ITEM] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.ANIMATE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.ANIMATE_COLOR] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.ANIMATE_MOTION] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.ANIMATE_TRANSFORM] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.CIRCLE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.CLIP_PATH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.COLOR_PROFILE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.CURSOR] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.DEFS] = func(m map[string]string) error {
		_, ok := m["end"]
		if ok {
			return nil
		}
		cmds := make([]svg.Command, 0)
		curID := ""
		for {
			select {
			case cmd := <-cmdCh:
				if cmd.C == svg.DEFS {
					_, ok = cmd.Style["end"]
					if ok {
						for _, c := range cmds {
							definitions[curID] = append(definitions[curID], c)
						}
						return nil
					}
				}
				ID, ok := cmd.Style["id"]
				if !ok || ID == curID {
					cmds = append(cmds, cmd)
					continue
				}
				curID = ID
				if len(cmds) == 0 {
					cmds = append(cmds, cmd)
					continue
				}
				for _, c := range cmds {
					definitions[curID] = append(definitions[curID], c)
				}
				cmds = make([]svg.Command, 0)
			case <-time.After(500 * time.Millisecond):
				return errors.New("ran out of commands before end of definitions")
			}
		}
	}
	drawFuncs[svg.DESC] = func(m map[string]string) error {
		// nothing to draw for description
		// TODO: see if there's a way to add metadata to an image.Image?
		return nil
	}
	drawFuncs[svg.ELLIPSE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_BLEND] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_COLOR_MATRIX] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_COMPONENT_TRANSFER] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_COMPOSITE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_CONVOLVE_MATRIX] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_DIFFUSE_LIGHTING] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_DISPLACEMENT_MAP] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_DISTANT_LIGHT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_FLOOD] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_FUNC_A] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_FUNC_B] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_FUNC_G] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_FUNC_R] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_GAUSSIAN_BLUR] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_IMAGE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_MERGE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_MERGE_NODE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_MORPHOLOGY] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_OFFSET] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_POINT_LIGHT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_SPECULAR_LIGHTING] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_SPOT_LIGHT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_TILE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FE_TURBULENCE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FILTER] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FONT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FONT_FACE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FONT_FACE_FORMAT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FONT_FACE_NAME] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FONT_FACE_SRC] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FONT_FACE_URI] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.FOREIGN_OBJECT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.G] = func(m map[string]string) error {
		dctx.Push()
		_, ok := m["end"]
		if ok {
			dctx.Pop()
			return nil
		}
		return setupPresentationAttributes(m)
	}
	drawFuncs[svg.GLYPH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.GLYPH_REF] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.H_KERN] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.IMAGE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.LINE] = func(m map[string]string) error {
		err := setupPresentationAttributes(m)
		if err != nil {
			return err
		}
		var x1len, x2len, y1len, y2len float64
		x1, ok := m["x1"]
		if ok {
			x1svglen, err := getSVGLength(x1, width)
			if err != nil {
				return err
			}
			x1len = x1svglen.convert(width)
		} else {
			x1len = 0
		}
		x2, ok := m["x2"]
		if ok {
			x2svglen, err := getSVGLength(x2, width)
			if err != nil {
				return err
			}
			x2len = x2svglen.convert(width)
		} else {
			x2len = 0
		}
		y1, ok := m["y1"]
		if ok {
			y1svglen, err := getSVGLength(y1, height)
			if err != nil {
				return err
			}
			y1len = y1svglen.convert(height)
		} else {
			y1len = 0
		}
		y2, ok := m["y2"]
		if ok {
			y2svglen, err := getSVGLength(y2, height)
			if err != nil {
				return err
			}
			y2len = y2svglen.convert(height)
		} else {
			y2len = 0
		}
		dctx.MoveTo(x1len, y1len)
		dctx.LineTo(x2len, y2len)
		dctx.SetColor(strokeColor)
		dctx.Stroke()
		return nil
	}
	drawFuncs[svg.LINEAR_GRADIENT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.MARKER] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.MASK] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.META_DATA] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.M_PATH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.MISSING_GLYPH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.PATH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.PATTERN] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.POLYGON] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.POLYLINE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.RADIAL_GRADIENT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.RECT] = func(m map[string]string) error {
		err := setupPresentationAttributes(m)
		if err != nil {
			return err
		}
		x, ok := m["x"]
		var xlen float64
		if ok {
			xsvglen, err := getSVGLength(x, width)
			if err != nil {
				return err
			}
			xlen = xsvglen.convert(width)
		} else {
			xlen = 0
		}
		y, ok := m["y"]
		var ylen float64
		if ok {
			ysvglen, err := getSVGLength(y, height)
			if err != nil {
				return err
			}
			ylen = ysvglen.convert(height)
		} else {
			ylen = 0
		}
		w, ok := m["width"]
		var widthlen float64
		if ok {
			widthsvglen, err := getSVGLength(w, width)
			if err != nil {
				return err
			}
			widthlen = widthsvglen.convert(width)
		} else {
			widthlen = 0
		}
		if widthlen < 0 {
			return errors.New("rectangle width less than zero")
		}
		h, ok := m["height"]
		var heightlen float64
		if ok {
			heightsvglen, err := getSVGLength(h, height)
			if err != nil {
				return err
			}
			heightlen = heightsvglen.convert(height)
		} else {
			heightlen = 0
		}
		if heightlen < 0 {
			return errors.New("rectangle height less than zero")
		}
		rx, ok := m["rx"]
		var rxlen float64
		var rxundefined bool
		if ok {
			rxsvglen, err := getSVGLength(rx, width)
			if err != nil {
				return err
			}
			rxlen = rxsvglen.convert(width)
		} else {
			rxlen = 0
			rxundefined = true
		}
		if rxlen < 0 {
			return errors.New("rectangle rx less than zero")
		}
		ry, ok := m["ry"]
		var rylen float64
		var ryundefined bool
		if ok {
			rysvglen, err := getSVGLength(ry, height)
			if err != nil {
				return err
			}
			rylen = rysvglen.convert(height)
		} else {
			rylen = 0
			ryundefined = true
		}
		if rylen < 0 {
			return errors.New("rectangle ry less than zero")
		}
		if rxundefined && !ryundefined {
			rxlen = rylen
		}
		if ryundefined && !rxundefined {
			rylen = rxlen
		}
		if rxlen > widthlen/2 {
			rxlen = widthlen / 2
		}
		if rylen > heightlen/2 {
			rylen = heightlen / 2
		}
		dctx.NewSubPath()
		dctx.MoveTo(xlen+rxlen, ylen)
		dctx.LineTo(xlen+widthlen-rxlen, ylen)
		dctx.DrawEllipticalArc(xlen+widthlen-rxlen, ylen+rylen, rxlen, rylen, gg.Radians(270), gg.Radians(360))
		dctx.LineTo(xlen+widthlen, ylen+heightlen-rylen)
		dctx.DrawEllipticalArc(xlen+widthlen-rxlen, ylen+heightlen-rylen, rxlen, rylen, gg.Radians(0), gg.Radians(90))
		dctx.LineTo(xlen+rxlen, ylen+heightlen)
		dctx.DrawEllipticalArc(xlen+rxlen, ylen+heightlen-rylen, rxlen, rylen, gg.Radians(90), gg.Radians(180))
		dctx.LineTo(xlen, ylen+rylen)
		dctx.DrawEllipticalArc(xlen+rxlen, ylen+rylen, rxlen, rylen, gg.Radians(180), gg.Radians(270))
		dctx.ClosePath()
		dctx.SetColor(fillColor)
		dctx.Fill()
		dctx.NewSubPath()
		dctx.MoveTo(xlen+rxlen, ylen)
		dctx.LineTo(xlen+widthlen-rxlen, ylen)
		dctx.DrawEllipticalArc(xlen+widthlen-rxlen, ylen+rylen, rxlen, rylen, gg.Radians(270), gg.Radians(360))
		dctx.LineTo(xlen+widthlen, ylen+heightlen-rylen)
		dctx.DrawEllipticalArc(xlen+widthlen-rxlen, ylen+heightlen-rylen, rxlen, rylen, gg.Radians(0), gg.Radians(90))
		dctx.LineTo(xlen+rxlen, ylen+heightlen)
		dctx.DrawEllipticalArc(xlen+rxlen, ylen+heightlen-rylen, rxlen, rylen, gg.Radians(90), gg.Radians(180))
		dctx.LineTo(xlen, ylen+rylen)
		dctx.DrawEllipticalArc(xlen+rxlen, ylen+rylen, rxlen, rylen, gg.Radians(180), gg.Radians(270))
		dctx.ClosePath()
		dctx.SetColor(strokeColor)
		dctx.Stroke()
		return nil
	}
	drawFuncs[svg.SCRIPT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.SET] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.STOP] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.STYLE] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.SWITCH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.SYMBOL] = func(m map[string]string) error {
		_, ok := m["end"]
		if ok {
			return nil
		}
		cmds := make([]svg.Command, 0)
		curID := ""
		for {
			select {
			case cmd := <-cmdCh:
				if cmd.C == svg.SYMBOL {
					_, ok = cmd.Style["end"]
					if ok {
						definitions[curID] = cmds
						return nil
					}
				}
				ID, ok := cmd.Style["id"]
				if !ok {
					cmds = append(cmds, cmd)
					continue
				}
				copy(definitions[curID], cmds)
				curID = ID
				cmds = make([]svg.Command, 0)
			case <-time.After(500 * time.Millisecond):
				return errors.New("ran out of commands before end of definitions")
			}
		}
	}
	drawFuncs[svg.TEXT] = func(m map[string]string) error {
		err := setupPresentationAttributes(m)
		if err != nil {
			return err
		}
		var xlen, ylen float64
		x, ok := m["x"]
		if ok {
			xsvglen, err := getSVGLength(x, none)
			if err != nil {
				return err
			}
			xlen = xsvglen.num
		} else {
			xlen = 0
		}
		y, ok := m["y"]
		if ok {
			ysvglen, err := getSVGLength(y, none)
			if err != nil {
				return err
			}
			ylen = ysvglen.num
		} else {
			ylen = 0
		}
		dctx.MoveTo(xlen, ylen)
		return nil
	}
	drawFuncs[svg.TEXT_PATH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.TITLE] = func(m map[string]string) error {
		// TODO: see if there's a way to set metadata for an image.Image
		return nil
	}
	drawFuncs[svg.T_REF] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.T_SPAN] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.USE] = func(m map[string]string) error {
		ref := m["href"]
		if !strings.HasPrefix(ref, "#") {
			return errors.New("css selectors other than id are not supported")
		}
		cmds := definitions[ref[1:]]
		for _, cmd := range cmds {
			attrs := make(map[string]string)
			for k, v := range m {
				attrs[k] = v
			}
			for k, v := range cmd.Style {
				attrs[k] = v
			}
			if err := drawFuncs[cmd.C](attrs); err != nil {
				return err
			}
		}
		return nil
	}
	drawFuncs[svg.VIEW] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.V_KERN] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.InnerText] = func(m map[string]string) error {
		text, ok := m["InnerText"]
		if !ok {
			return errors.New("no inner text found")
		}
		err := setupPresentationAttributes(m)
		if err != nil {
			return err
		}
		var xlen, ylen float64
		x, ok := m["x"]
		if ok {
			xsvglen, err := getSVGLength(x, none)
			if err != nil {
				return err
			}
			xlen = xsvglen.num
		} else {
			xlen = 0
		}
		y, ok := m["y"]
		if ok {
			ysvglen, err := getSVGLength(y, none)
			if err != nil {
				return err
			}
			ylen = ysvglen.num
		} else {
			ylen = 0
		}
		dctx.SetColor(fillColor)
		dctx.DrawString(text, xlen, ylen)
		return nil
	}
}

func setDefaults() {
	fillColor = color.Transparent
	strokeColor = color.Black
}
