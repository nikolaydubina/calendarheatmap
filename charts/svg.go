package charts

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"text/template"
)

// Day is SVG template day parameters
type Day struct {
	Count int
	Date  string
	Color string
	Show  bool
}

// WeekdayLabel is SVG template weekday label parameters
type WeekdayLabel struct {
	Label string
	Show  bool
}

// MonthLabel is SVG template month label parameters
type MonthLabel struct {
	Label   string
	XOffset int
}

// Params is total SVG template parameters
type Params struct {
	Days           [53][7]Day
	LabelsMonths   [12]MonthLabel
	LabelsWeekdays [7]WeekdayLabel
	LabelsColor    string
}

func writeSVGColor(c color.RGBA) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func writeSVG(conf HeatmapConfig, w io.Writer) {
	fullYearTemplate := template.Must(template.New("fullyear").Funcs(template.FuncMap{
		"mul": func(a int, b int) int { return a * b },
		"add": func(a int, b int) int { return a + b },
		"sub": func(a int, b int) int { return a - b },
	}).Parse(fullyear))

	locale := conf.Locale
	if locale == "" {
		locale = "en_US"
	}
	labelsProvider := NewLabelsProvider(locale)

	labelsWeekdays := [7]WeekdayLabel{}
	for i, w := range weekdayOrder {
		labelsWeekdays[i] = WeekdayLabel{labelsProvider.GetWeekday(w), conf.ShowWeekdays[w]}
	}

	labelsMonths := [12]MonthLabel{}
	for i, v := range labelsProvider.months {
		labelsMonths[i-1].Label = v
	}

	month := 0
	days := [53][7]Day{}
	for iter := NewDayIterator(conf.Counts, image.Point{}, 0, 0); !iter.Done(); iter.Next() {
		days[iter.Col][iter.Row] = Day{
			Count: iter.Count(),
			Date:  iter.Time().Format("2006-01-02"),
			Color: writeSVGColor(conf.ColorScale.GetColor(iter.Value())),
			Show:  true,
		}

		// Note, day is from 1~31
		if iter.Row == 0 && iter.Time().Day() <= 7 {
			labelsMonths[month].XOffset = iter.Col
			month++
		}
	}

	fullYearTemplate.Execute(w, Params{
		Days:           days,
		LabelsMonths:   labelsMonths,
		LabelsWeekdays: labelsWeekdays,
		LabelsColor:    writeSVGColor(conf.TextColor),
	})
}
