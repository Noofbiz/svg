package svg

import (
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
)

type svg struct {
	commands     []Command
	svgCount     int
	useChannels  bool
	cmdCh        chan Command
	errCh        chan error
	getInnerText bool
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

//ParseSVGChannel parses the SVG file on a separate goroutine and returns a channel
//to receive the draw commands on.
func ParseSVGChannel(r io.Reader) (chan Command, chan error) {
	s := new(svg)
	s.cmdCh = make(chan Command, 10)
	s.errCh = make(chan error)
	s.useChannels = true

	go func(r io.Reader, s *svg) {
		var buf []byte
		var err error
		if buf, err = ioutil.ReadAll(r); err != nil {
			s.errCh <- err
			return
		}

		if err = xml.Unmarshal(buf, s); err != nil {
			s.errCh <- err
			return
		}
	}(r, s)

	return s.cmdCh, s.errCh
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
		if s.useChannels {
			s.cmdCh <- Command{
				C:     cmds[svgstr],
				Style: attrs,
			}
		} else {
			s.commands = append(s.commands, Command{
				C:     cmds[svgstr],
				Style: attrs,
			})
		}
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
			if shouldGrabStartTag(tok.Name.Local) {
				c, ok := cmds[tok.Name.Local]
				if !ok {
					return errors.New(tok.Name.Local + " was not recognized as an svg element")
				}
				if tok.Name.Local == "text" {
					s.getInnerText = true
				}
				if s.useChannels {
					s.cmdCh <- Command{
						C:     c,
						Style: attrs,
					}
				} else {
					s.commands = append(s.commands, Command{
						C:     c,
						Style: attrs,
					})
				}
			}
		case xml.EndElement:
			if !inSVG {
				continue loop
			}
			if shouldGetInnerText(tok.Name.Local) {
				s.getInnerText = false
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
			if shouldMarkEndTag(tok.Name.Local) {
				attrs["end"] = ""
			}
			if s.useChannels {
				s.cmdCh <- Command{
					C:     c,
					Style: attrs,
				}
			} else {
				s.commands = append(s.commands, Command{
					C:     c,
					Style: attrs,
				})
			}
			if tok.Name.Local == svgstr {
				s.svgCount--
				if s.svgCount == 0 {
					break loop
				}
			}
			ctx = ctx[:len(ctx)-1]
		case xml.CharData:
			if s.getInnerText {
				attrs := make(map[string]string)
				for k, v := range ctx[len(ctx)-1] {
					attrs[k] = v
				}
				attrs["InnerText"] = string([]byte(tok))
				if s.useChannels {
					s.cmdCh <- Command{
						C:     InnerText,
						Style: attrs,
					}
				} else {
					s.commands = append(s.commands, Command{
						C:     InnerText,
						Style: attrs,
					})
				}
			}
		}
	}
	return nil
}

func shouldGrabStartTag(s string) bool {
	tags := []string{
		"g",
		"text",
		"defs",
		"symbol",
	}
	return inList(s, tags)
}

func shouldGetInnerText(s string) bool {
	tags := []string{
		"text",
	}
	return inList(s, tags)
}

func shouldMarkEndTag(s string) bool {
	tags := []string{
		"g",
		"defs",
		"symbol",
	}
	return inList(s, tags)
}

func inList(s string, list []string) bool {
	for _, l := range list {
		if l == s {
			return true
		}
	}
	return false
}
