package svg

type Command struct {
	C     Cmd
	Style map[string]string
}

type Cmd uint

const (
	SVG Cmd = iota
	A
	ALT_GLYPH
	ALT_GLYPH_DEF
	ALT_GLYPH_ITEM
	ANIMATE
	ANIMATE_COLOR
	ANIMATE_MOTION
	ANIMATE_TRANSFORM
	CIRCLE
	CLIP_PATH
	COLOR_PROFILE
	CURSOR
	DEFS
	DESC
	ELLIPSE
	FE_BLEND
	FE_COLOR_MATRIX
	FE_COMPONENT_TRANSFER
	FE_COMPOSITE
	FE_CONVOLVE_MATRIX
	FE_DIFFUSE_LIGHTING
	FE_DISPLACEMENT_MAP
	FE_DISTANT_LIGHT
	FE_FLOOD
	FE_FUNC_A
	FE_FUNC_B
	FE_FUNC_G
	FE_FUNC_R
	FE_GAUSSIAN_BLUR
	FE_IMAGE
	FE_MERGE
	FE_MERGE_NODE
	FE_MORPHOLOGY
	FE_OFFSET
	FE_POINT_LIGHT
	FE_SPECULAR_LIGHTING
	FE_SPOT_LIGHT
	FE_TILE
	FE_TURBULENCE
	FILTER
	FONT
	FONT_FACE
	FONT_FACE_FORMAT
	FONT_FACE_NAME
	FONT_FACE_SRC
	FONT_FACE_URI
	FOREIGN_OBJECT
	G
	GLYPH
	GLYPH_REF
	H_KERN
	IMAGE
	LINE
	LINEAR_GRADIENT
	MARKER
	MASK
	META_DATA
	MISSING_GLYPH
	M_PATH
	PATH
	PATTERN
	POLYGON
	POLYLINE
	RADIAL_GRADIENT
	RECT
	SCRIPT
	SET
	STOP
	STYLE
	SWITCH
	SYMBOL
	TEXT
	TEXT_PATH
	TITLE
	T_REF
	T_SPAN
	USE
	VIEW
	V_KERN
)

var cmds = map[string]Cmd{
	"svg":                 SVG,
	"a":                   A,
	"altGlyph":            ALT_GLYPH,
	"altGlyphDef":         ALT_GLYPH_DEF,
	"altGlyphItem":        ALT_GLYPH_ITEM,
	"animate":             ANIMATE,
	"animateColor":        ANIMATE_COLOR,
	"animateMotion":       ANIMATE_MOTION,
	"animateTransform":    ANIMATE_TRANSFORM,
	"circle":              CIRCLE,
	"clipPath":            CLIP_PATH,
	"color-profile":       COLOR_PROFILE,
	"cursor":              CURSOR,
	"defs":                DEFS,
	"desc":                DESC,
	"ellipse":             ELLIPSE,
	"feBlend":             FE_BLEND,
	"feColorMatrix":       FE_COLOR_MATRIX,
	"feComponentTransfer": FE_COMPONENT_TRANSFER,
	"feComposite":         FE_COMPOSITE,
	"feConvolveMatrix":    FE_CONVOLVE_MATRIX,
	"feDiffuseLighting":   FE_DIFFUSE_LIGHTING,
	"feDisplacementMap":   FE_DISPLACEMENT_MAP,
	"feDistantLight":      FE_DISTANT_LIGHT,
	"feFlood":             FE_FLOOD,
	"feFuncA":             FE_FUNC_A,
	"feFuncB":             FE_FUNC_B,
	"feFuncG":             FE_FUNC_G,
	"feFuncR":             FE_FUNC_R,
	"feGaussianBlur":      FE_GAUSSIAN_BLUR,
	"feImage":             FE_IMAGE,
	"feMerge":             FE_MERGE,
	"feMergeNode":         FE_MERGE_NODE,
	"feMorphology":        FE_MORPHOLOGY,
	"feOffset":            FE_OFFSET,
	"fePointLight":        FE_POINT_LIGHT,
	"feSpecularLighting":  FE_SPECULAR_LIGHTING,
	"feSpotLight":         FE_SPOT_LIGHT,
	"feTile":              FE_TILE,
	"feTurbulence":        FE_TURBULENCE,
	"filter":              FILTER,
	"font":                FONT,
	"font_face":           FONT_FACE,
	"font_face-format":    FONT_FACE_FORMAT,
	"font-face-name":      FONT_FACE_NAME,
	"font-face-src":       FONT_FACE_SRC,
	"font-face-uri":       FONT_FACE_URI,
	"foreignObject":       FOREIGN_OBJECT,
	"g":                   G,
	"glyph":               GLYPH,
	"glyphRef":            GLYPH_REF,
	"hkern":               H_KERN,
	"image":               IMAGE,
	"line":                LINE,
	"linearGradient":      LINEAR_GRADIENT,
	"marker":              MARKER,
	"mask":                MASK,
	"metadata":            META_DATA,
	"missing-glyph":       MISSING_GLYPH,
	"mpath":               M_PATH,
	"path":                PATH,
	"pattern":             PATTERN,
	"polygon":             POLYGON,
	"polyline":            POLYLINE,
	"radialGradient":      RADIAL_GRADIENT,
	"rect":                RECT,
	"script":              SCRIPT,
	"set":                 SET,
	"stop":                STOP,
	"style":               STYLE,
	"switch":              SWITCH,
	"symbol":              SYMBOL,
	"text":                TEXT,
	"textPath":            TEXT_PATH,
	"title":               TITLE,
	"tref":                T_REF,
	"tSpan":               T_SPAN,
	"use":                 USE,
	"view":                VIEW,
	"vkern":               V_KERN,
}
