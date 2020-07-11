// This is example on how to read data, calculate statistics
// and draw it with this module.
package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/nikolaydubina/calendarheatmap/charts"
	"github.com/nikolaydubina/calendarheatmap/colorscales"
	"github.com/nikolaydubina/calendarheatmap/parsers"
)

func main() {
	filenameInput := flag.String("input", "input.txt", "input filename")
	filenameChart := flag.String("output", "chart.png", "output filename, will export as PNG")
	monthSep := flag.Bool("monthsep", true, "render month separator")
	colorScale := flag.String("colorscale", "PuBu9", "refer to colorscales for examples")
	labels := flag.Bool("labels", true, "labels for weekday and months")
	outputFormat := flag.String("output-format", "png", "output format (png, jpeg, gif)")
	inputFormat := flag.String("input-format", "json-basic", "format of input file refer to `/parsers` for examples")
	flag.Parse()

	data, err := ioutil.ReadFile(*filenameInput)
	if err != nil {
		log.Fatal(fmt.Errorf("cant not read file: %w", err))
	}

	var parser parsers.Parser
	switch *inputFormat {
	case "row-day-seconds-count":
		parser = &parsers.RowDaySecondsCountParser{}
	case "json-basic":
		parser = &parsers.BasicJSONParser{}
	default:
		log.Fatal("unnknown parser format")
		return
	}
	year, countByDay, err := parser.Parse(data)

	img := charts.NewHeatmap(charts.HeatmapConfig{
		Year:               year,
		CountByDay:         countByDay,
		ColorScale:         colorscales.LoadColorScale(*colorScale),
		DrawMonthSeparator: *monthSep,
		DrawLabels:         *labels,
		Margin:             3,
		BoxSize:            15,
		TextWidthLeft:      35,
		TextHeightTop:      20,
		TextColor:          color.RGBA{100, 100, 100, 255},
		BorderColor:        color.RGBA{200, 200, 200, 255},
	})
	f, err := os.Create(*filenameChart)
	if err != nil {
		log.Fatal(fmt.Errorf("can not create file: %w", err))
	}
	defer f.Close()

	switch *outputFormat {
	case "png":
		if err := png.Encode(f, img); err != nil {
			log.Fatal(fmt.Errorf("can not encode png: %w", err))
		}
	case "jpeg":
		if err := jpeg.Encode(f, img, nil); err != nil {
			log.Fatal(fmt.Errorf("can not encode jpeg: %w", err))
		}
	case "gif":
		if err := gif.Encode(f, img, nil); err != nil {
			log.Fatal(fmt.Errorf("can not encode gifg: %w", err))
		}
	default:
		log.Fatal("unknown output format")
	}
}
