package charts

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

// drawText inserts text into provided image at bottom left coordinate
func drawText(img *image.RGBA, offset image.Point, text string, color color.RGBA) {
	assetsPath := getEnv("CALENDAR_HEATMAP_ASSETS_PATH", "charts/assets")
	if assetsPath == "" {
		log.Fatalf("assets path is not set")
	}
	fontBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/fonts/Sunflower-Medium.ttf", assetsPath))
	if err != nil {
		log.Fatalf("can not open font file with error: %#v", err)
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("can not parse font file: %v", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    26,
		DPI:     280,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("can not create font face: %v", err)
	}

	point := fixed.Point26_6{
		X: fixed.Int26_6(offset.X * 64),
		Y: fixed.Int26_6(offset.Y * 64),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: face,
		Dot:  point,
	}
	d.DrawString(text)
}
