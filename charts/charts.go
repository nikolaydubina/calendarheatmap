package charts

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/nikolaydubina/calendarheatmap/colorscales"
)

var weekdaysPos = map[time.Weekday]int{
	time.Monday:    0,
	time.Tuesday:   1,
	time.Wednesday: 2,
	time.Thursday:  3,
	time.Friday:    4,
	time.Saturday:  5,
	time.Sunday:    6,
}

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

var weekdayOrder = [7]time.Weekday{
	time.Monday,
	time.Tuesday,
	time.Wednesday,
	time.Thursday,
	time.Friday,
	time.Saturday,
	time.Sunday,
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

const (
	numWeeksYear  = 52
	numWeekCols   = numWeeksYear + 1 // 53 * 7 = 371 > 366
	margin        = 3                // should be odd number for best result if using month separator
	boxSize       = 15
	textWidthLeft = 35
	textHightTop  = 20
)

var textColor = color.RGBA{100, 100, 100, 255}
var borderColor = color.RGBA{200, 200, 200, 255}

// HeatmapConfig contains config of calendar heatmap image
type HeatmapConfig struct {
	Year               int
	CountByDay         map[int]int
	ColorScale         colorscales.ColorScale
	DrawMonthSeparator bool
	DrawLabels         bool
}

// NewHeatmap draws every day of a year as square
// filled with color proportional to counter from the max.
func NewHeatmap(conf HeatmapConfig) image.Image {
	maxCount := 0
	for _, q := range conf.CountByDay {
		if q > maxCount {
			maxCount = q
		}
	}

	// one margin less on the right and bottom side to avoid whitespace
	width := numWeekCols*(boxSize+margin) - margin + textWidthLeft
	height := 7*(boxSize+margin) - margin + textHightTop
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	x, y := 0, 0

	// text margins
	x += textWidthLeft
	y += textHightTop

	yearStartDate := time.Date(conf.Year, 1, 1, 1, 1, 1, 1, time.UTC)
	vIdx := weekdaysPos[yearStartDate.Weekday()]

	// draw each cell as block
	for day := yearStartDate; day.Year() == conf.Year; day = day.Add(time.Hour * 24) {
		y = (boxSize+margin)*vIdx + textHightTop

		r := image.Rect(x, y, x+boxSize, y+boxSize)
		val := float64(conf.CountByDay[day.YearDay()]) / float64(maxCount)
		color := conf.ColorScale.GetColor(val)
		draw.Draw(img, r, &image.Uniform{color}, image.ZP, draw.Src)

		if conf.DrawMonthSeparator {
			drawMonthSeparator(img, day, image.Point{X: x, Y: y}, textHightTop, height)
		}

		if conf.DrawLabels {
			drawMonthLabels(img, day, image.Point{X: x, Y: y})
		}

		vIdx++
		if vIdx == 7 {
			vIdx = 0
			x += boxSize + margin
		}
	}

	if conf.DrawLabels {
		drawWeekdayLabels(
			img,
			image.Point{X: textWidthLeft, Y: textHightTop},
			map[time.Weekday]bool{
				time.Monday:    true,
				time.Wednesday: true,
				time.Friday:    true,
			},
		)
	}

	return img
}

// drawMonthSeparator progressively draws months separator into provided image
func drawMonthSeparator(img *image.RGBA, day time.Time, point image.Point, minY int, maxY int) {
	if day.Day() == 1 && day.Month() != time.January {
		marginSep := margin / 2

		closeLeft := image.Point{X: point.X - marginSep - 1, Y: point.Y - marginSep - 1}
		closeRight := image.Point{X: point.X + boxSize + marginSep, Y: point.Y - marginSep - 1}
		farLeft := image.Point{X: point.X - marginSep - 1, Y: maxY}
		farRight := image.Point{X: point.X + boxSize + marginSep, Y: minY}

		drawLineAxis(img, farLeft, closeLeft, borderColor) // left line
		if day.Weekday() != time.Monday {
			drawLineAxis(img, closeRight, farRight, borderColor)  // right line
			drawLineAxis(img, closeLeft, closeRight, borderColor) // top line
		}
	}
}

// drawMonthLabels progressively draws text of months based on current day into provided image
func drawMonthLabels(img *image.RGBA, day time.Time, point image.Point) {
	if day.Day() == 1 {
		monthLabelX := point.X
		if weekdaysPos[day.Weekday()] != 0 {
			monthLabelX += boxSize + margin
		}
		drawText(img, image.Point{X: monthLabelX, Y: textHightTop - 5}, monthLabel[day.Month()])
	}
}

// drawWeekdayLabel draws column of same width labels for weekdays
// All weekday labels assumed to have same width, which really depends on font.
// offset argument is top right corner of where to insert column of weekday labels.
func drawWeekdayLabels(img *image.RGBA, offset image.Point, weekdays map[time.Weekday]bool) {
	width := 25
	height := 10
	x := offset.X - width
	y := offset.Y + height
	for _, w := range weekdayOrder {
		if weekdays[w] {
			drawText(img, image.Point{X: x, Y: y}, weekdayLabel[w])
		}
		y += boxSize + margin
	}
}
