package svg

import (
	"encoding/xml"
	"image"
)

type svg struct {
	dst            *image.RGBA
	scaleX, scaleY float64
	svgCount       int
}

func parseSVG(b []byte, scaleX, scaleY float64) (image.Image, error) {
	var err error
	s := &svg{
		scaleX: scaleX,
		scaleY: scaleY,
	}

	if err = xml.Unmarshal(b, s); err != nil {
		return nil, err
	}

	return s.dst, nil
}

//UnmarshalXML implements the xml.Unmarshaler interface
func (s *svg) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error
	inSVG := false
loop:
	for {
		var token xml.Token
		if token, err = decoder.Token(); err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			if !inSVG {
				if tok.Name.Local != "svg" {
					continue loop
				}
				inSVG = true
			}
			attrs := make(map[string]string)
			for _, attr := range tok.Attr {
				attrs[attr.Name.Local] = attr.Value
			}
			switch tok.Name.Local {
			case "svg":
				s.svgCount++
				setViewport(attrs["width"], attrs["height"], attrs["viewBox"])
			}
		case xml.EndElement:
			switch tok.Name.Local {
			case "svg":
				s.svgCount--
				if s.svgCount == 0 {
					break loop
				}
			}
		}
	}
	return nil
}
