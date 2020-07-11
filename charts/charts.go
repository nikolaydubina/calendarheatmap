package charts

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

var monthLabel = map[time.Month]string{
	time.January:   "Jan",
	time.February:  "Feb",
	time.March:     "Mar",
	time.April:     "Apr",
	time.May:       "May",
	time.June:      "Jun",
	time.July:      "Jul",
	time.August:    "Aug",
	time.September: "Sep",
	time.October:   "Oct",
	time.November:  "Nov",
	time.December:  "Dec",
}

var weekdayLabel = map[time.Weekday]string{
	time.Monday:    "Mon",
	time.Tuesday:   "Tue",
	time.Wednesday: "Wed",
	time.Thursday:  "Thu",
	time.Friday:    "Fri",
	time.Saturday:  "Sat",
	time.Sunday:    "Sun",
}

var weekdayOrder = [7]time.Weekday{
	time.Monday,
	time.Tuesday,
	time.Wednesday,
	time.Thursday,
	time.Friday,
	time.Saturday,
	time.Sunday,
}

const (
	numWeeksYear = 52
	numWeekCols  = numWeeksYear + 1 // 53 * 7 = 371 > 366
)

// HeatmapConfig contains config of calendar heatmap image
type HeatmapConfig struct {
	Year               int
	CountByDay         map[int]int
	ColorScale         colorscales.ColorScale
	DrawMonthSeparator bool
	DrawLabels         bool
	BoxSize            int
	Margin             int
	TextWidthLeft      int
	TextHeightTop      int
	TextColor          color.RGBA
	BorderColor        color.RGBA
}

// NewHeatmap creates image with heatmap and additional elements
func NewHeatmap(conf HeatmapConfig) image.Image {
	width := conf.TextWidthLeft + numWeekCols*(conf.BoxSize+conf.Margin)
	height := conf.TextHeightTop + 7*(conf.BoxSize+conf.Margin)
	offset := image.Point{X: conf.TextWidthLeft, Y: conf.TextHeightTop}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	visitors := []DayVisitor{
		&DayBoxVisitor{img, conf.ColorScale, conf.BoxSize},
	}

	if conf.DrawMonthSeparator {
		visitors = append(
			visitors,
			&MonthSeparatorVisitor{
				Img:     img,
				MinY:    conf.TextHeightTop,
				MaxY:    height - conf.Margin,
				Margin:  conf.Margin,
				BoxSize: conf.BoxSize,
				Width:   1,
				Color:   conf.BorderColor,
			},
		)
	}

	if conf.DrawLabels {
		visitors = append(visitors, &MonthLabelsVisitor{Img: img, YOffset: 5, Color: conf.TextColor})
	}

	for iter := NewDayIterator(conf.Year, offset, conf.CountByDay, conf.BoxSize, conf.Margin); !iter.Done(); iter.Next() {
		for _, v := range visitors {
			v.Visit(iter)
		}
	}

	if conf.DrawLabels {
		drawWeekdayLabels(
			img,
			offset,
			map[time.Weekday]bool{
				time.Monday:    true,
				time.Wednesday: true,
				time.Friday:    true,
			},
			conf.BoxSize,
			conf.Margin,
			conf.TextColor,
		)
	}

	return img
}

// DayVisitor is interface to update image based on current box
type DayVisitor interface {
	Visit(iter *DayIterator)
}

// DayBoxVisitor draws signle heatbox
type DayBoxVisitor struct {
	Img        *image.RGBA
	ColorScale colorscales.ColorScale
	BoxSize    int
}

// Visit called on every iteration
func (d *DayBoxVisitor) Visit(iter *DayIterator) {
	p := iter.Point()
	r := image.Rect(p.X, p.Y, p.X+d.BoxSize, p.Y+d.BoxSize)
	color := d.ColorScale.GetColor(iter.Value())
	draw.Draw(d.Img, r, &image.Uniform{color}, image.ZP, draw.Src)
}

// MonthSeparatorVisitor draws month separator
type MonthSeparatorVisitor struct {
	Img     *image.RGBA
	MinY    int
	MaxY    int
	Margin  int
	BoxSize int
	Width   int
	Color   color.RGBA
}

// Visit called on every iteration
func (d *MonthSeparatorVisitor) Visit(iter *DayIterator) {
	day := iter.Time()
	if day.Day() == 1 && day.Month() != time.January {
		p := iter.Point()

		marginSep := d.Margin / 2

		xL := p.X - marginSep - d.Width
		xR := p.X + d.BoxSize + marginSep

		// left vertical line
		drawLineAxis(
			d.Img,
			image.Point{X: xL, Y: p.Y},
			image.Point{X: xL, Y: d.MaxY - d.Width},
			d.Color,
		)
		if day.Weekday() != weekdayOrder[0] {
			// right vertical line
			drawLineAxis(
				d.Img,
				image.Point{X: xR, Y: d.MinY},
				image.Point{X: xR, Y: p.Y - marginSep - d.Width},
				d.Color,
			)
			// right vertical line
			drawLineAxis(
				d.Img,
				image.Point{X: xL, Y: p.Y - marginSep - d.Width},
				image.Point{X: xR, Y: p.Y - marginSep - d.Width},
				d.Color,
			)
			// connect left vertical line and horizontal one
			drawLineAxis(
				d.Img,
				image.Point{X: xL, Y: p.Y - marginSep - d.Width},
				image.Point{X: xL, Y: p.Y},
				d.Color,
			)
		}
	}
}

// MonthLabelsVisitor draws month label on top of first row 0 of month
type MonthLabelsVisitor struct {
	Img     *image.RGBA
	YOffset int
	Color   color.RGBA
}

// Visit on every iteration
func (d *MonthLabelsVisitor) Visit(iter *DayIterator) {
	day := iter.Time()
	// Note, day is from 1~31
	if iter.Row == 0 && day.Day() <= 7 {
		p := iter.Point()
		drawText(
			d.Img,
			image.Point{X: p.X, Y: p.Y - d.YOffset},
			monthLabel[day.Month()],
			d.Color,
		)
	}
}

// drawWeekdayLabel draws column of same width labels for weekdays
// All weekday labels assumed to have same width, which really depends on font.
// offset argument is top right corner of where to insert column of weekday labels.
func drawWeekdayLabels(img *image.RGBA, offset image.Point, weekdays map[time.Weekday]bool, boxSize int, margin int, color color.RGBA) {
	width := 25
	height := 10
	y := offset.Y + height
	for _, w := range weekdayOrder {
		if weekdays[w] {
			drawText(img, image.Point{X: offset.X - width, Y: y}, weekdayLabel[w], color)
		}
		y += boxSize + margin
	}
}
