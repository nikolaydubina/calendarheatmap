package charts_test

import (
	"bytes"
	"encoding/json"
	"image/color"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"golang.org/x/image/font"

	"github.com/nikolaydubina/calendarheatmap/charts"
)

func loadData(t *testing.T, filepath string) map[string]int {
	var counts map[string]int
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(data, &counts); err != nil {
		t.Error(err)
	}
	return counts
}

func loadFontFace(t *testing.T, filepath string) font.Face {
	fontFace, err := charts.LoadFontFaceFromFile(filepath)
	if err != nil {
		t.Error(err)
	}
	return fontFace
}

func loadColorscale(t *testing.T, filepath string) charts.ColorScale {
	colorscale, err := charts.NewBasicColorscaleFromCSVFile(filepath)
	if err != nil {
		t.Fail()
	}
	return colorscale
}

func TestCharts(t *testing.T) {
	tests := []struct {
		name         string
		outputPath   string
		expectedPath string
		conf         charts.HeatmapConfig
	}{
		{
			name:         "basic-png",
			outputPath:   path.Join("testdata", "basic-png-output.png"),
			expectedPath: path.Join("testdata", "basic-png-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "basic-jpeg",
			outputPath:   path.Join("testdata", "basic-jpeg-output.jpeg"),
			expectedPath: path.Join("testdata", "basic-jpeg-expected.jpeg"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "jpeg",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "basic-svg",
			outputPath:   path.Join("testdata", "basic-svg-output.svg"),
			expectedPath: path.Join("testdata", "basic-svg-expected.svg"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "svg",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "no-data",
			outputPath:   path.Join("testdata", "basic-no-data-output.png"),
			expectedPath: path.Join("testdata", "basic-no-data-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             nil,
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "no-labels",
			outputPath:   path.Join("testdata", "basic-no-labels-output.png"),
			expectedPath: path.Join("testdata", "basic-no-labels-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         false,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "no-separator",
			outputPath:   path.Join("testdata", "basic-no-separator-output.png"),
			expectedPath: path.Join("testdata", "basic-no-separator-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: false,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "korean",
			outputPath:   path.Join("testdata", "basic-korean-output.png"),
			expectedPath: path.Join("testdata", "basic-korean-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "ko_KR",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       map[time.Weekday]bool{time.Monday: true, time.Wednesday: true, time.Friday: true},
			},
		},
		{
			name:         "no-weekdays",
			outputPath:   path.Join("testdata", "basic-no-weekdays-output.png"),
			expectedPath: path.Join("testdata", "basic-no-weekdays-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays:       nil,
			},
		},
		{
			name:         "all-weekdays",
			outputPath:   path.Join("testdata", "basic-all-weekdays-output.png"),
			expectedPath: path.Join("testdata", "basic-all-weekdays-expected.png"),
			conf: charts.HeatmapConfig{
				Counts:             loadData(t, path.Join("testdata", "basic.json")),
				ColorScale:         loadColorscale(t, path.Join("..", "assets", "colorscales", "purple-blue-9.csv")),
				DrawMonthSeparator: true,
				DrawLabels:         true,
				BoxSize:            150,
				Margin:             30,
				TextWidthLeft:      350,
				TextHeightTop:      200,
				TextColor:          color.RGBA{100, 100, 100, 255},
				BorderColor:        color.RGBA{200, 200, 200, 255},
				Locale:             "en_US",
				Format:             "png",
				FontFace:           loadFontFace(t, path.Join("..", "assets", "fonts", "Sunflower-Medium.ttf")),
				ShowWeekdays: map[time.Weekday]bool{
					time.Monday:    true,
					time.Tuesday:   true,
					time.Wednesday: true,
					time.Thursday:  true,
					time.Friday:    true,
					time.Saturday:  true,
					time.Sunday:    true,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// output
			outputfile, err := os.Create(tc.outputPath)
			if err != nil {
				t.Error(err)
			}
			if err := charts.WriteHeatmap(tc.conf, outputfile); err != nil {
				t.Error(err)
			}
			if err := outputfile.Close(); err != nil {
				t.Error(err)
			}

			// compare to expected
			expected, err := ioutil.ReadFile(tc.expectedPath)
			if err != nil {
				t.Error(err)
			}
			actual, err := ioutil.ReadFile(tc.outputPath)
			if err != nil {
				t.Error(err)
			}
			if !bytes.Equal(expected, actual) {
				t.Fail()
			}
		})
	}
}
