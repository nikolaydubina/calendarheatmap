package charts

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// drawLineAxis draws line parallel to X or Y axis
func drawLineAxis(img draw.Image, a image.Point, b image.Point, col color.Color) {
	switch {
	// do not attempt to draw dot
	case a == b:
		return
	// vertical
	case a.X == b.X:
		y1, y2 := a.Y, b.Y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for q := y1; q <= y2; q++ {
			img.Set(a.X, q, col)
		}
	// horizontal
	case a.Y == b.Y:
		x1, x2 := a.X, b.X
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for q := x1; q <= x2; q++ {
			img.Set(q, a.Y, col)
		}
	default:
		panic("input line is not parallel to axis. not implemented")
	}
}

// drawText inserts text into provided image at bottom left coordinate
func drawText(img *image.RGBA, offset image.Point, text string) {
	point := fixed.Point26_6{
		X: fixed.Int26_6(offset.X * 64),
		Y: fixed.Int26_6(offset.Y * 64),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}
