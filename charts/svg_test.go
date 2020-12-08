package charts_test

import (
	"image/color"
	"testing"

	"github.com/nikolaydubina/calendarheatmap/charts"
	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

func TestBasicDataSVG(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "svg",
		}
		save(t, conf, "testdata/basic.svg")
	})

	t.Run("korean", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Locale:             "ko_KR",
			Format:             "svg",
		}
		save(t, conf, "testdata/korean.svg")
	})

	t.Run("empty data", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             map[string]int{},
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "svg",
		}
		save(t, conf, "testdata/empty_data.svg")
	})

	t.Run("nil data", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             nil,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "svg",
		}
		save(t, conf, "testdata/nil_data.svg")
	})
}
