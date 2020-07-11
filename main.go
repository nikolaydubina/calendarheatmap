// This is example on how to read data, calculate statistics
// and draw it with this module.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nikolaydubina/calendarheatmap/charts"
	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

// Row represents single row data in input file
type Row struct {
	Date  time.Time
	Count int
}

func loadRows(filename string) ([]Row, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("cant not open file: %w", err)
	}
	defer file.Close()

	rows := make([]Row, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		if len(items) != 3 {
			return nil, fmt.Errorf("number of items in row is not 3")
		}
		timeString, countString := items[0]+" "+items[1], items[2]

		date, err := time.Parse("2006-01-02 15:04", timeString)
		if err != nil {
			return nil, fmt.Errorf("can not parse time: %w", err)
		}
		count := strings.Count(countString, "P")
		rows = append(rows, Row{Date: date, Count: count})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner got error: %w", err)
	}
	return rows, nil
}

func main() {
	filenameLogs := flag.String("input", "input.txt", "file should contain lines in format: 2020-05-16 20:43 PPPP")
	filenameChart := flag.String("output", "chart.png", "output filename, will export as PNG")
	monthSep := flag.Bool("monthsep", true, "render month separator")
	colorScale := flag.String("colorscale", "PuBu9", "refer to colorscales for examples")
	labels := flag.Bool("labels", true, "labels for weekday and months")
	outputFormat := flag.String("output-format", "png", "output format (png, jpeg, gif)")
	flag.Parse()

	rows, err := loadRows(*filenameLogs)
	if err != nil {
		log.Fatal(err)
	}

	year := rows[0].Date.Year()
	countByDay := make(map[int]int, 366)
	for _, row := range rows {
		countByDay[row.Date.YearDay()] += row.Count
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
		TextHightTop:       20,
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
	}
}
