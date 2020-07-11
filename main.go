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
	colorScale := flag.String("colorscale", "PuBu9", "refer to colorscales for examples")
	inputFormat := flag.String("input", "json-basic", "format of input file, refer to parsers module")
	labels := flag.Bool("labels", true, "labels for weekday and months")
	monthSep := flag.Bool("monthsep", true, "render month separator")
	outputFormat := flag.String("output", "png", "output format (png, jpeg, gif)")
	flag.Parse()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(fmt.Errorf("cant not read from stdin: %w", err))
	}

	var parser parsers.Parser
	switch *inputFormat {
	case "row-day-seconds-count":
		parser = &parsers.RowDaySecondsCountParser{}
	case "json-basic":
		parser = &parsers.BasicJSONParser{}
	default:
		log.Fatal("unknown parser format")
		return
	}
	year, countByDay, err := parser.Parse(data)
	if err != nil {
		log.Fatal("error parsing data: %w", err)
	}

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

	outWriter := os.Stdout
	switch *outputFormat {
	case "png":
		if err := png.Encode(outWriter, img); err != nil {
			log.Fatal(fmt.Errorf("can not encode png: %w", err))
		}
	case "jpeg":
		if err := jpeg.Encode(outWriter, img, nil); err != nil {
			log.Fatal(fmt.Errorf("can not encode jpeg: %w", err))
		}
	case "gif":
		if err := gif.Encode(outWriter, img, nil); err != nil {
			log.Fatal(fmt.Errorf("can not encode gifg: %w", err))
		}
	default:
		log.Fatal("unknown output format")
	}
}
