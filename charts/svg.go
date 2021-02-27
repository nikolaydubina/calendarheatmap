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

// Params is total SVG template parameters
type Params struct {
	Days           [53][7]Day
	LabelsMonths   [12]string
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

	days := [53][7]Day{}

	for iter := NewDayIterator(conf.Counts, image.Point{}, 0, 0); !iter.Done(); iter.Next() {
		days[iter.Col][iter.Row] = Day{
			Count: iter.Count(),
			Date:  iter.Time().Format("2006-01-02"),
			Color: writeSVGColor(conf.ColorScale.GetColor(iter.Value())),
			Show:  true,
		}
	}

	locale := conf.Locale
	if locale == "" {
		locale = "en_US"
	}
	labelsProvider := NewLabelsProvider(locale)

	labelsMonths := [12]string{}
	for i, v := range labelsProvider.months {
		labelsMonths[i-1] = v
	}

	labelsWeekdays := [7]WeekdayLabel{}
	for i, v := range labelsProvider.weekdays {
		labelsWeekdays[i] = WeekdayLabel{v, true}
	}

	fullYearTemplate.Execute(w, Params{
		Days:           days,
		LabelsMonths:   labelsMonths,
		LabelsWeekdays: labelsWeekdays,
		LabelsColor:    writeSVGColor(conf.TextColor),
	})
}
