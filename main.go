package main

import (
	"encoding/json"
	"flag"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/nikolaydubina/calendarheatmap/charts"
)

func main() {
	var (
		colorScale   string
		labels       bool
		locale       string
		monthSep     bool
		outputFormat string
		assetsPath   string
	)

	flag.BoolVar(&labels, "labels", true, "labels for weekday and months")
	flag.BoolVar(&monthSep, "monthsep", true, "render month separator")
	flag.StringVar(&colorScale, "colorscale", "green-blue-9.csv", "filename of colorscale")
	flag.StringVar(&locale, "locale", "en_US", "locale of labels (en_US, ko_KR)")
	flag.StringVar(&outputFormat, "output", "png", "output format (png, jpeg, gif, svg)")
	flag.StringVar(&assetsPath, "assetspath", "", "absolute path, or relative path for executable, of calendarheatmap repo assets, if not set will try CALENDARHEATMAP_ASSETS env variable, if not will try 'assets'")
	flag.Parse()

	if assetsPath == "" {
		assetsPath = os.Getenv("CALENDAR_HEATMAP_ASSETS_PATH")
		if assetsPath == "" {
			assetsPath = "assets"
		}
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var counts map[string]int
	if err := json.Unmarshal(data, &counts); err != nil {
		log.Fatal(err)
	}

	colorscale, err := charts.NewBasicColorscaleFromCSVFile(path.Join(assetsPath, "colorscales", colorScale))
	if err != nil {
		log.Fatal(err)
	}

	fontFace, err := charts.LoadFontFaceFromFile(path.Join(assetsPath, "fonts", "Sunflower-Medium.ttf"))
	if err != nil {
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
