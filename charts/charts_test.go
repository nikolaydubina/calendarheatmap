package charts

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

func savePNG(t *testing.T, img image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		t.Errorf(fmt.Errorf("can not save: %w", err).Error())
	}
	if err := png.Encode(f, img); err != nil {
		t.Errorf(fmt.Errorf("can not encode png: %w", err).Error())
	}
	if err := f.Close(); err != nil {
		t.Errorf(fmt.Errorf("can not close: %w", err).Error())
	}
}

func TestBasicData(t *testing.T) {
	countByDay := map[int]int{
		137: 8, 138: 13, 139: 5, 140: 8, 141: 5, 142: 5, 143: 3, 144: 5,
		145: 6, 146: 3, 147: 5, 148: 8, 149: 2, 150: 3, 151: 8, 152: 5,
		153: 1, 154: 3, 155: 1, 156: 3, 157: 1, 158: 3, 159: 5, 161: 1,
		162: 2, 164: 9, 165: 7, 166: 4, 167: 1, 169: 1, 172: 2, 173: 1,
		175: 2, 176: 2, 177: 3, 178: 3, 179: 2, 180: 1, 181: 1, 182: 2,
	}

	t.Run("basic", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/basic.png")
	})

	t.Run("colorscale_1", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.GnBu9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/colorscale_1.png")
	})

	t.Run("colorscale_2", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.YlGn9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/colorscale_2.png")
	})

	t.Run("no separator", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: false,
			DrawLabels:         true,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/noseparator.png")
	})

	t.Run("no labels", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/nolabels.png")
	})

	t.Run("no separator, no labels", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/noseparator_nolabels.png")
	})

	t.Run("empty data", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         map[int]int{},
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/empty_data.png")
	})

	t.Run("nil data", func(t *testing.T) {
		img := NewHeatmap(HeatmapConfig{
			Year:               2020,
			CountByDay:         nil,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             3,
			BoxSize:            15,
			TextWidthLeft:      35,
			TextHightTop:       20,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		})
		savePNG(t, img, "testdata/nil_data.png")
	})
}
