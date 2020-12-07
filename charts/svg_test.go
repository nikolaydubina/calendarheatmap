package charts

import (
	"fmt"
	"image/color"
	"os"
	"testing"

	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

func saveSVG(t *testing.T, conf HeatmapConfig, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		t.Errorf(fmt.Errorf("can not save: %w", err).Error())
	}
	NewHeatmapSVG(conf, f)
	if err := f.Close(); err != nil {
		t.Errorf(fmt.Errorf("can not close: %w", err).Error())
	}
}

func TestBasicDataSVG(t *testing.T) {
	os.Setenv("CALENDAR_HEATMAP_ASSETS_PATH", "assets")
	countByDay := map[int]int{
		137: 8, 138: 13, 139: 5, 140: 8, 141: 5, 142: 5, 143: 3, 144: 5, 145: 6,
		146: 3, 147: 5, 148: 8, 149: 2, 150: 2, 151: 8, 152: 5, 153: 1, 154: 3,
		155: 1, 156: 3, 157: 1, 158: 3, 159: 5, 161: 1, 162: 2, 164: 9, 165: 7,
		166: 4, 167: 1, 169: 1, 172: 2, 173: 1, 175: 2, 176: 2, 177: 3, 178: 3,
		179: 2, 180: 1, 181: 1, 182: 2,
	}

	t.Run("basic", func(t *testing.T) {
		conf := HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		}
		saveSVG(t, conf, "testdata/basic.svg")
	})

	t.Run("korean", func(t *testing.T) {
		conf := HeatmapConfig{
			Year:               2020,
			CountByDay:         countByDay,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Locale:             "ko_KR",
		}
		saveSVG(t, conf, "testdata/korean.svg")
	})

	t.Run("empty data", func(t *testing.T) {
		conf := HeatmapConfig{
			Year:               2020,
			CountByDay:         map[int]int{},
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		}
		saveSVG(t, conf, "testdata/empty_data.svg")
	})

	t.Run("nil data", func(t *testing.T) {
		conf := HeatmapConfig{
			Year:               2020,
			CountByDay:         nil,
			ColorScale:         colorscales.PuBu9,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
		}
		saveSVG(t, conf, "testdata/nil_data.svg")
	})
}
