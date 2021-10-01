package charts

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"strings"
	"time"
)

// Day is SVG day parameters
type Day struct {
	Count int
	Date  string
	Color string
	Show  bool
}

// WeekdayLabel is SVG weekday label parameters
type WeekdayLabel struct {
	Label string
	Show  bool
}

// MonthLabel is SVG month label parameters
type MonthLabel struct {
	Label   string
	XOffset int
}

// Params is total SVG parameters
type Params struct {
	BoxAreaHeight   int
	BoxAreaWidth    int
	BoxSize         int
	Margin          int
	Days            [][]Day
	LabelsMonths    []MonthLabel
	LabelsWeekdays  []WeekdayLabel
	LabelsColor     string
	MonthSeparators []string
}

func writeSVGColor(c color.RGBA) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

func writeSVG(conf HeatmapConfig, w io.Writer) {
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
        // wasm tingy crashes due to 2K arguments in function. it inlines them for arrays and flattens structs
        // 2K ~ 53 * 7 * 4
        days := make([][]Day, 53)
        for i := 0; i < 53; i++ {
            days[i] = make([]Day, 7)
        }
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

        result := render(Params{
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
        w.Write([]byte(result))
}

// before this was template, but tingo does not support text/template
func render(params Params) string {
    out := ""
    out += fmt.Sprintf(
        `<svg viewBox="0 0 %d %d" xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">`,
        params.BoxAreaWidth,
        params.BoxAreaHeight,

    )
    out += `<g transform="translate(25, 23)">`

    for w, wo := range params.Days {
        out += fmt.Sprintf(`<g transform="translate(%d, 0)">`, w * (params.BoxSize + params.Margin))

        for d, do := range wo {
            if do.Show {
                out += fmt.Sprintf(
                    `<rect class="day" width="%d" height="%d" x="0" y="%d" fill="%s" data-count="%d" data-date="%s"></rect>`,
                    params.BoxSize,
                    params.BoxSize,
                    d * (params.BoxSize + params.Margin),
                    do.Color,
                    do.Count,
                    do.Date,
                )
            }
        }

        out += `</g>`
    }

    for _, label := range params.LabelsMonths {
        out += fmt.Sprintf(
            `<text x="%d" y="-7" font-size="10" fill="%s">%s</text>`,
            label.XOffset * (params.BoxSize + params.Margin),
            params.LabelsColor,
            label.Label,
        )
    }

    for i, o := range params.LabelsWeekdays {
        style := ""
        if !o.Show {
            style = `style="display: none;"`
        }
        out += fmt.Sprintf(
            `<text text-anchor="start" font-size="10" dx="-25" dy="%d" fill="%s" %s>%s</text>`,
            15 + i * (params.BoxSize + params.Margin),
            params.LabelsColor,
            style,
            o.Label,
        )
    }

    for _, l := range params.MonthSeparators {
        out += l + "\n"
    }

    out += `</g></svg>`

    return out
}
