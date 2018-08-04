# svg
A simple SVG decoder.

It has been redesigned to do just that. Pass it an svg file as an io.Reader

```go
var cmd []svg.Command
var err error
if cmd, err = svg.ParseSVG(r); err != nil {
	fmt.Println(err.Error())
}
```

and the `[]svg.Command` is the step-by-step list of commands and styles you can
use to generate the svg image.

The package `svg/drawsvg` uses this list to generate an image.Image by utilizing
`draw2d`, but you can implement your own drawing methods if you want to only use
the `svg` package, for example you can draw directly on an OpenGL surface.

To Do:

[] Setup Viewport

[] Groups

[] Rects

[] Circle

[] Ellipse

[] Polygons

[] Lines

[] Curves

[] More tags coming!
