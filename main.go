package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	_ "embed"

	"github.com/nikolaydubina/calendarheatmap/charts"
)

//go:embed assets/fonts/Sunflower-Medium.ttf
var defaultFontFaceBytes []byte

//go:embed assets/colorscales/green-blue-9.csv
var defaultColorScaleBytes []byte

func main() {
	var (
		colorScale   string
		labels       bool
		locale       string
		monthSep     bool
		outputFormat string
	)

	flag.BoolVar(&labels, "labels", true, "labels for weekday and months")
	flag.BoolVar(&monthSep, "monthsep", true, "render month separator")
	flag.StringVar(&colorScale, "colorscale", "green-blue-9.csv", "filename of colorscale")
	flag.StringVar(&locale, "locale", "en_US", "locale of labels (en_US, ko_KR)")
	flag.StringVar(&outputFormat, "output", "png", "output format (png, jpeg, gif, svg)")
	flag.Parse()

	var colorscale charts.BasicColorScale
	if assetsPath := os.Getenv("CALENDAR_HEATMAP_ASSETS_PATH"); assetsPath != "" {
		var err error
		colorscale, err = charts.NewBasicColorscaleFromCSVFile(path.Join(assetsPath, "colorscales", colorScale))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var err error
		if colorScale != "green-blue-9.csv" {
			log.Printf("defaulting to colorscale %s since CALENDAR_HEATMAP_ASSETS_PATH is not set", "green-blue-9.csv")
		}
		colorscale, err = charts.NewBasicColorscaleFromCSV(bytes.NewBuffer(defaultColorScaleBytes))
		if err != nil {
			log.Fatal(err)
		}
	}

	fontFace, err := charts.LoadFontFace(defaultFontFaceBytes)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var counts map[string]int
	if err := json.Unmarshal(data, &counts); err != nil {
		log.Fatal(err)
	}

	conf := charts.HeatmapConfig{
		Counts:             counts,
		ColorScale:         colorscale,
		DrawMonthSeparator: monthSep,
		DrawLabels:         labels,
		Margin:             30,
		BoxSize:            150,
		TextWidthLeft:      350,
		TextHeightTop:      200,
		TextColor:          color.RGBA{100, 100, 100, 255},
		BorderColor:        color.RGBA{200, 200, 200, 255},
		Locale:             locale,
		Format:             outputFormat,
		FontFace:           fontFace,
		ShowWeekdays: map[time.Weekday]bool{
			time.Monday:    true,
			time.Wednesday: true,
			time.Friday:    true,
		},
	}
	charts.WriteHeatmap(conf, os.Stdout)
}
