package main

import (
	"encoding/json"
	"flag"
	"image/color"
	"io/ioutil"
	"log"
	"os"

	"github.com/nikolaydubina/calendarheatmap/charts"
	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

func main() {
	var (
		colorScale   string
		colorScaleAlt   string
		labels       bool
		locale       string
		monthSep     bool
		outputFormat string
		maxCount int
	)

	flag.BoolVar(&labels, "labels", true, "labels for weekday and months")
	flag.BoolVar(&monthSep, "monthsep", true, "render month separator")
	flag.StringVar(&colorScale, "colorscale", "PuBu9", "refer to colorscales for examples")
	flag.StringVar(&colorScaleAlt, "colorscalealt", "YlGn9", "alternative colorscale used for negative counts")
	flag.StringVar(&locale, "locale", "en_US", "locale of labels (en_US, ko_KR)")
	flag.StringVar(&outputFormat, "output", "png", "output format (png, jpeg, gif, svg)")
	flag.IntVar(&maxCount, "maxcount", 0, "maximum count possible for each day (use 0 to calculate it based on input data)")
	flag.Parse()

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
		MaxCount:           maxCount,
		ColorScale:         colorscales.LoadColorScale(colorScale),
		ColorScaleAlt:         colorscales.LoadColorScale(colorScaleAlt),
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
	}
	charts.WriteHeatmap(conf, os.Stdout)
}
