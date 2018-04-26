package svg

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

type svg struct {
	commands []Command
	svgCount int
}

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

	for _, c := range s.commands {
		fmt.Printf("Command: %v\nStyle: %v\n\n\n", c.C, c.Style)
	}

	return s.commands, nil
}

//UnmarshalXML implements the xml.Unmarshaler interface
func (s *svg) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error
	inSVG := false
	if start.Name.Local == "svg" {
		inSVG = true
		for _, attr := range start.Attr {
			ctx.style[attr.Name.Local] = attr.Value
		}
		s.commands = append(s.commands, Command{
			C:     cmds["svg"],
			Style: ctx.style,
		})
		ctx.push()
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
			if !inSVG {
				if tok.Name.Local != "svg" {
					continue loop
				}
				inSVG = true
			}
			for _, attr := range tok.Attr {
				ctx.style[attr.Name.Local] = attr.Value
			}
			ctx.push()
			if tok.Name.Local == "svg" {
				s.svgCount++
			}
		case xml.EndElement:
			if !inSVG {
				continue loop
			}
			c, ok := cmds[tok.Name.Local]
			if !ok {
				return fmt.Errorf("%v was not recognized as an svg element", tok.Name.Local)
			}
			s.commands = append(s.commands, Command{
				C:     c,
				Style: ctx.style,
			})

			if tok.Name.Local == "svg" {
				s.svgCount--
				if s.svgCount == 0 {
					break loop
				}
			}
			ctx.pop()
		}
	}
	return nil
}
