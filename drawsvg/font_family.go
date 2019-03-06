package drawsvg

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	find "github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

func getSVGFontFamily(s string) (*truetype.Font, error) {
	var fnt *truetype.Font
	var err error
	fonts := strings.Split(s, ",")
	for _, v := range fonts {
		v = strings.TrimSpace(v)
		if isGenericFont(v) {
			fnt, err = getGenericFont(v)
			return fnt, err
		}
		fnt, err = systemFontPath(v)
		if err != nil {
			if strings.HasPrefix(err.Error(), "cannot find font") {
				//check for embedded fonts or web fonts?
			}
		}
		if fnt != nil {
			break
		}
	}
	return fnt, nil
}

var genericFonts = []string{
	"serif",
	"sans-serif",
	"cursive",
	"fantasy",
	"monospace",
}

var loadedTTFs = make(map[string]*truetype.Font)

func isGenericFont(s string) bool {
	for _, v := range genericFonts {
		if v == s {
			return true
		}
	}
	return false
}

func getGenericFont(s string) (*truetype.Font, error) {
	f, ok := loadedTTFs[s]
	if ok {
		return f, nil
	}
	switch s {
	case "serif":
		fnt, err := systemFontPath("Times New Roman")
		if err != nil {
			if strings.HasPrefix(err.Error(), "cannot find font") {
				fnt, err = systemFontPath("FreeSerif")
			}
		}
		return fnt, err
	case "sans-serif":
		fnt, err := loadFont(goregular.TTF, s)
		return fnt, err
	case "cursive":
		fnt, err := systemFontPath("Brush Script")
		if err != nil {
			if strings.HasPrefix(err.Error(), "cannot find font") {
				fnt, err = systemFontPath("Apple Chancery")
				if strings.HasPrefix(err.Error(), "cannot find font") {
					fnt, err = systemFontPath("URW Chancery L")
				}
			}
		}
		return fnt, err
	case "fantasy":
		fnt, err := systemFontPath("Papyrus")
		return fnt, err
	case "monospace":
		fnt, err := loadFont(gomono.TTF, s)
		return fnt, err
	}
	return nil, errors.New(s + " is not a generic font")
}

func systemFontPath(name string) (*truetype.Font, error) {
	path, err := find.Find(name + ".ttf")
	if err != nil {
		return nil, err
	}
	fnt, err := loadFromPath(path, name)
	return fnt, err
}

func loadFromPath(path, name string) (*truetype.Font, error) {
	fnt, ok := loadedTTFs[name]
	if ok {
		return fnt, nil
	}
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	fnt, err = loadFont(b, name)
	return fnt, err
}

func loadFont(ttf []byte, s string) (*truetype.Font, error) {
	fnt, err := truetype.Parse(ttf)
	if err != nil {
		return nil, err
	}
	loadedTTFs[s] = fnt
	return fnt, nil
}
