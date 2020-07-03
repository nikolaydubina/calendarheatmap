package charts

import (
	"image"
	"image/color"
	"testing"
)

func TestBasicDrawAxis(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	t.Run("along x", func(t *testing.T) {
		drawLineAxis(img, image.Point{X: 0, Y: 0}, image.Point{X: 0, Y: 10}, color.Black)
	})

	t.Run("along y", func(t *testing.T) {
		drawLineAxis(img, image.Point{X: 0, Y: 0}, image.Point{X: 10, Y: 0}, color.Black)
	})

	t.Run("reverse x", func(t *testing.T) {
		drawLineAxis(img, image.Point{X: 0, Y: 10}, image.Point{X: 0, Y: 0}, color.Black)
	})

	t.Run("reverse y", func(t *testing.T) {
		drawLineAxis(img, image.Point{X: 10, Y: 0}, image.Point{X: 0, Y: 0}, color.Black)
	})

	t.Run("dot", func(t *testing.T) {
		drawLineAxis(img, image.Point{X: 0, Y: 0}, image.Point{X: 0, Y: 0}, color.Black)
	})
}
