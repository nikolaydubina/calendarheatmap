package charts

import (
	"fmt"
	"image"
	"io"
	"text/template"
)

type Day struct {
	Count int
	Date  string
	Color string
	Show  bool
}

type WeekdayLabel struct {
	Label string
	Show  bool
}

type Params struct {
	Days           [53][7]Day
	LabelsMonths   [12]string
	LabelsWeekdays [7]WeekdayLabel
	LabelColor     string
}

func NewHeatmapSVG(conf HeatmapConfig, w io.Writer) {
	fullYearTemplate := template.Must(template.New("fullyear").Funcs(template.FuncMap{
		"mul": func(a int, b int) int { return a * b },
		"add": func(a int, b int) int { return a + b },
		"sub": func(a int, b int) int { return a - b },
	}).Parse(fullyear))

	days := [53][7]Day{}

	for iter := NewDayIterator(conf.Year, image.Point{}, conf.CountByDay, 0, 0); !iter.Done(); iter.Next() {
		color := conf.ColorScale.GetColor(iter.Value())
		days[iter.Col][iter.Row] = Day{
			Count: iter.Count(),
			Date:  iter.Time().Format("2006-01-02"),
			Color: fmt.Sprintf("rgb(%d,%d,%d)", color.R, color.G, color.B),
			Show:  true,
		}
	}

	locale := "en_US"
	if conf.Locale != "" {
		locale = conf.Locale
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

	params := Params{
		Days:           days,
		LabelsMonths:   labelsMonths,
		LabelsWeekdays: labelsWeekdays,
		LabelColor:     fmt.Sprintf("rgb(%d,%d,%d)", conf.TextColor.R, conf.TextColor.G, conf.TextColor.B),
	}

	fullYearTemplate.Execute(w, params)
}
