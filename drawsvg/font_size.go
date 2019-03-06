package drawsvg

var relativeFontSizes = []string{
	"larger",
	"smaller",
}

var absoluteFontSizes = []string{
	"xx-small",
	"x-small",
	"small",
	"medium",
	"large",
	"x-large",
	"xx-large",
}

func isRelativeOrAbsoluteFontSize(s string) bool {
	for _, a := range absoluteFontSizes {
		if a == s {
			return true
		}
	}
	for _, r := range relativeFontSizes {
		if r == s {
			return true
		}
	}
	return false
}
