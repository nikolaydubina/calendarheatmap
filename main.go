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
	os.Setenv("CALENDAR_HEATMAP_ASSETS_PATH", "charts/assets")

	colorScale := *flag.String("colorscale", "PuBu9", "refer to colorscales for examples")
	labels := *flag.Bool("labels", true, "labels for weekday and months")
	monthSep := *flag.Bool("monthsep", true, "render month separator")
	outputFormat := *flag.String("output", "png", "output format (png, jpeg, gif, svg)")
	locale := *flag.String("locale", "en_US", "locale of labels (en_US, ko_KR)")
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
		ColorScale:         colorscales.LoadColorScale(colorScale),
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
