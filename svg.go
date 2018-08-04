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
		attrs["fill"] = "black"
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
				attrs["fill"] = "black"
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
			if tok.Name.Local == "g" || tok.Name.Local == "text" {
				c, ok := cmds[tok.Name.Local]
				if !ok {
					return errors.New(tok.Name.Local + " was not recognized as an svg element")
				}
				s.commands = append(s.commands, Command{
					C:     c,
					Style: attrs,
				})
			}
		case xml.EndElement:
			if !inSVG {
				continue loop
			}
			if tok.Name.Local == "text" {
				continue loop
			}
			c, ok := cmds[tok.Name.Local]
			if !ok {
				return errors.New(tok.Name.Local + " was not recognized as an svg element")
			}
			attrs := make(map[string]string)
			for k, v := range ctx[len(ctx)-1] {
				attrs[k] = v
			}
			if tok.Name.Local == "g" {
				attrs["end"] = ""
			}
			s.commands = append(s.commands, Command{
				C:     c,
				Style: attrs,
			})
			if tok.Name.Local == svgstr {
				s.svgCount--
				if s.svgCount == 0 {
					break loop
				}
			}
			ctx = ctx[:len(ctx)-1]
		case xml.CharData:
			if s.commands[len(s.commands)-1].C == TEXT {
				attrs := make(map[string]string)
				for k, v := range ctx[len(ctx)-1] {
					attrs[k] = v
				}
				attrs["InnerText"] = string([]byte(tok))
				s.commands = append(s.commands, Command{
					C:     InnerText,
					Style: attrs,
				})
			}
		}
	}
	return nil
}
