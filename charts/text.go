package charts

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

// LoadFontFaceFromFile loads font face from file
func LoadFontFaceFromFile(fontPath string) (font.Face, error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, fmt.Errorf("can not open font file with error: %w", err)
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("can not parse font file: %w", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    26,
		DPI:     280,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, fmt.Errorf("can not create font face: %w", err)
	}
	return face, nil
}

// drawText inserts text into provided image at bottom left coordinate
func drawText(fontFace font.Face, img *image.RGBA, offset image.Point, text string, color color.RGBA) {
	if fontFace == nil {
		return
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: fontFace,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(offset.X * 64),
			Y: fixed.Int26_6(offset.Y * 64),
		},
	}
	d.DrawString(text)
}
