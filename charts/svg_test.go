package charts_test

import (
	"image/color"
	"testing"

	"github.com/nikolaydubina/calendarheatmap/charts"
)

func TestBasicDataSVG(t *testing.T) {
	var colorscale = charts.BasicColorScale{
		color.RGBA{247, 252, 240, 255},
		color.RGBA{224, 243, 219, 255},
		color.RGBA{204, 235, 197, 255},
		color.RGBA{168, 221, 181, 255},
		color.RGBA{123, 204, 196, 255},
		color.RGBA{78, 179, 211, 255},
		color.RGBA{43, 140, 190, 255},
		color.RGBA{8, 104, 172, 255},
		color.RGBA{8, 64, 129, 255},
	}

	t.Run("basic", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
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
			ColorScale:         colorscale,
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
			ColorScale:         colorscale,
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
			ColorScale:         colorscale,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "svg",
		}
		save(t, conf, "testdata/nil_data.svg")
	})
}
