# svg
SVG decoder using draw2d.

Right now it doesn't do anything! Yay!

To use:

```go
package main

import (
	"log"
	"os"

	"github.com/Noofbiz/svg"
)

func main() {
	var err error

	var in *os.File
	if in, err = os.Open("./in.svg"); err != nil {
		log.Fatal(err.Error())
	}

	xScaleFactor := 1.0
	yScaleFactor := 1.0
	img, err := svg.Decode(in, xScaleFactor, yScaleFactor)
	if err != nil {
		log.Fatal(err.Error())
	}
}
```

This will *eventually* return an image.Image of the svg at the provided scale.

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
