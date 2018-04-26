package draw

import (
	"fmt"
	"image"

	"github.com/Noofbiz/svg"
	"github.com/llgcode/draw2d/draw2dimg"
)

var dctx *draw2dimg.GraphicContext
var img *image.RGBA

func drawSVG(cmds []svg.Command, scaleX, scaleY float64) (image.Image, error) {
	initStyle := cmds[0].Style
	viewPortSetup(initStyle["width"], initStyle["height"], initStyle["viewbox"])

	img = image.NewRGBA(image.Rect(0, 0, int(userWidth.convert(width)), int(userHeight.convert(height))))
	dctx = draw2dimg.NewGraphicContext(img)
	for _, cmd := range cmds {
		err := drawFuncs[cmd.C](cmd.Style)
		if err != nil {
			return nil, err
		}
	}
	dctx.Scale(scaleX, scaleY)
	return img, nil
}

var drawFuncs = make(map[svg.Cmd]func(map[string]string) error)

func init() {
	drawFuncs[svg.SVG] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.A] = func(m map[string]string) error {
		fmt.Println("hello")
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
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.DESC] = func(m map[string]string) error {
		fmt.Println("hello")
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
		fmt.Println("hello")
		return nil
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
		fmt.Println("hello")
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
		fmt.Println("hello")
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
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.TEXT] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.TEXT_PATH] = func(m map[string]string) error {
		fmt.Println("hello")
		return nil
	}
	drawFuncs[svg.TITLE] = func(m map[string]string) error {
		fmt.Println("hello")
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
		fmt.Println("hello")
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
}
