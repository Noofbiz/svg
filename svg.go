package svg

import (
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
)

type svg struct {
	commands []Command
	svgCount int
}

const svgstr string = "svg"

var ctx = make([]map[string]string, 0)

// ParseSVG parses an SVG file and returns the draw commands required to
// rasterize the SVG
func ParseSVG(r io.Reader) ([]Command, error) {
	var err error
	s := new(svg)

	var buf []byte
	if buf, err = ioutil.ReadAll(r); err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(buf, s); err != nil {
		return nil, err
	}

	return s.commands, nil
}

//UnmarshalXML implements the xml.Unmarshaler interface
func (s *svg) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error
	inSVG := false
	if start.Name.Local == svgstr {
		inSVG = true
		attrs := make(map[string]string)
		for _, attr := range start.Attr {
			attrs[attr.Name.Local] = attr.Value
		}
		s.commands = append(s.commands, Command{
			C:     cmds[svgstr],
			Style: attrs,
		})
		ctx = append(ctx, attrs)
		s.svgCount++
	}
loop:
	for {
		var token xml.Token
		if token, err = decoder.Token(); err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			attrs := make(map[string]string)
			if !inSVG {
				if tok.Name.Local != svgstr {
					continue loop
				}
				inSVG = true
			}
			if len(ctx) != 0 {
				for k, v := range ctx[len(ctx)-1] {
					attrs[k] = v
				}
			}
			for _, attr := range tok.Attr {
				attrs[attr.Name.Local] = attr.Value
			}
			ctx = append(ctx, attrs)
			if tok.Name.Local == svgstr {
				s.svgCount++
			}
		case xml.EndElement:
			if !inSVG {
				continue loop
			}
			c, ok := cmds[tok.Name.Local]
			if !ok {
				return errors.New(tok.Name.Local + " was not recognized as an svg element")
			}
			s.commands = append(s.commands, Command{
				C:     c,
				Style: ctx[len(ctx)-1],
			})

			if tok.Name.Local == svgstr {
				s.svgCount--
				if s.svgCount == 0 {
					break loop
				}
			}
			ctx = ctx[0 : len(ctx)-1]
		}
	}
	return nil
}
