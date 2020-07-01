package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nikolaydubina/plotstats/charts"
	"github.com/nikolaydubina/plotstats/colorscales"
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
	filenameChart := flag.String("output", "chart.png", "output filename")
	monthSep := flag.Bool("monthsep", true, "redner month separator")
	colorScale := flag.String("colorscale", "PuBu9", "refer to colorscales for examples")
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

	img := charts.MakeYearDayHeatmapHoriz(
		year,
		countByDay,
		colorscales.LoadColorScale(*colorScale),
		*monthSep,
	)
	f, err := os.Create(*filenameChart)
	if err != nil {
		log.Fatal(fmt.Errorf("can not create file: %w", err))
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(fmt.Errorf("can not encode png: %w", err))
	}
}
