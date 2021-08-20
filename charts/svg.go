package charts

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"io"
	"strings"
	"text/template"
	"time"
)

//go:embed template.svg
var svgTemplate string

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
	BoxAreaHeight   int
	BoxAreaWidth    int
	BoxSize         int
	Margin          int
	Days            [53][7]Day
	LabelsMonths    []MonthLabel
	LabelsWeekdays  []WeekdayLabel
	LabelsColor     string
	MonthSeparators []string
}

func writeSVGColor(c color.RGBA) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func writeSVG(conf HeatmapConfig, w io.Writer) {
	fullYearTemplate := template.Must(template.New("fullyear").Funcs(template.FuncMap{
		"mul": func(a int, b int) int { return a * b },
		"add": func(a int, b int) int { return a + b },
	}).Parse(svgTemplate))

	labelsProvider := NewLabelsProvider(conf.Locale)

	var labelsWeekdays []WeekdayLabel
	var labelsMonths []MonthLabel

	if conf.DrawLabels {
		labelsWeekdays = make([]WeekdayLabel, 7)
		for i, w := range weekdayOrder {
			labelsWeekdays[i] = WeekdayLabel{labelsProvider.GetWeekday(w), conf.ShowWeekdays[w]}
		}

		labelsMonths = make([]MonthLabel, 12)
		for i, v := range labelsProvider.months {
			labelsMonths[i-1].Label = v
		}
	}

	monthSeparators := []string{}

	month := 0
	days := [53][7]Day{}
	boxw := 18
	margin := 4

	for iter := NewDayIterator(conf.Counts, image.Point{}, boxw, margin); !iter.Done(); iter.Next() {
		days[iter.Col][iter.Row] = Day{
			Count: iter.Count(),
			Date:  iter.Time().Format("2006-01-02"),
			Color: writeSVGColor(conf.ColorScale.GetColor(iter.Value())),
			Show:  true,
		}

		// Update month label offset
		// Note, day is from 1~31
		if len(labelsMonths) > 0 {
			if iter.Row == 0 && iter.Time().Day() <= 7 {
				labelsMonths[month].XOffset = iter.Col
				month++
			}
		}

		if conf.DrawMonthSeparator {
			day := iter.Time()
			if day.Day() == 1 && day.Month() != time.January {
				p := iter.Point()

				points := []string{
					// left vertical line
					fmt.Sprintf("%d,%d", p.X-2, 7*(boxw+margin)-margin),
					fmt.Sprintf("%d,%d", p.X-2, p.Y),
				}

				if day.Weekday() != weekdayOrder[0] {
					points = append(points,
						// horizontal line
						fmt.Sprintf("%d,%d", p.X-2, p.Y-2),
						fmt.Sprintf("%d,%d", p.X+boxw+2, p.Y-2),

						// right vertical line
						fmt.Sprintf("%d,%d", p.X+boxw+2, p.Y-2),
						fmt.Sprintf("%d,%d", p.X+boxw+2, 0),
					)
				}

				line := fmt.Sprintf(`<polyline style="fill:none;stroke:%s;stroke-width:1" points="%s"></polyline>`, writeSVGColor(conf.BorderColor), strings.Join(points, " "))
				monthSeparators = append(monthSeparators, line)
			}
		}
	}

	fullYearTemplate.Execute(w, Params{
		BoxSize:         boxw,
		Margin:          margin,
		BoxAreaWidth:    54 * (boxw + margin),
		BoxAreaHeight:   8 * (boxw + margin),
		Days:            days,
		LabelsMonths:    labelsMonths,
		LabelsWeekdays:  labelsWeekdays,
		LabelsColor:     writeSVGColor(conf.TextColor),
		MonthSeparators: monthSeparators,
	})
}
