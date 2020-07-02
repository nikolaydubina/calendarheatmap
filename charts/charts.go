package charts

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/nikolaydubina/calendarheatmap/colorscales"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
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

	for day := yearStartDate; day.Year() == conf.Year; day = day.Add(time.Hour * 24) {
		y = (boxSize+margin)*vIdx + textHightTop

		r := image.Rect(x, y, x+boxSize, y+boxSize)
		val := float64(conf.CountByDay[day.YearDay()]) / float64(maxCount)
		color := conf.ColorScale.GetColor(val)
		draw.Draw(img, r, &image.Uniform{color}, image.ZP, draw.Src)

		if conf.DrawMonthSeparator {
			if day.Day() == 1 && day.Month() != time.January {
				marginSep := margin / 2

				closeLeft := image.Point{X: x - marginSep - 1, Y: y - marginSep - 1}
				closeRight := image.Point{X: x + boxSize + marginSep, Y: y - marginSep - 1}
				farLeft := image.Point{X: x - marginSep - 1, Y: height}
				farRight := image.Point{X: x + boxSize + marginSep, Y: textHightTop}

				drawLineAxis(img, farLeft, closeLeft, borderColor) // left line
				if vIdx != 0 {
					drawLineAxis(img, closeRight, farRight, borderColor)  // right line
					drawLineAxis(img, closeLeft, closeRight, borderColor) // top line
				}
			}
		}

		if conf.DrawLabels {
			switch day.Weekday() {
			case time.Monday:
				addLabel(img, textWidthLeft-25, textHightTop+15, "Mon")
			case time.Wednesday:
				addLabel(img, textWidthLeft-25, textHightTop+15+(boxSize+margin)*2, "Wed")
			case time.Friday:
				addLabel(img, textWidthLeft-25, textHightTop+15+(boxSize+margin)*4, "Fri")
			}

			if day.Day() == 1 {
				monthLabelX := x
				if weekdaysPos[day.Weekday()] != 0 {
					monthLabelX += boxSize + margin
				}
				addLabel(img, monthLabelX, textHightTop-5, monthLabel[day.Month()])
			}
		}

		vIdx++
		if vIdx == 7 {
			vIdx = 0
			x += boxSize + margin
		}
	}

	return img
}

// drawLineAxis draws line parallel to X or Y axis
func drawLineAxis(img draw.Image, a image.Point, b image.Point, col color.Color) {
	switch {
	// do not attempt to draw dot
	case a == b:
		return
	// vertical
	case a.X == b.X:
		y1, y2 := a.Y, b.Y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for q := y1; q <= y2; q++ {
			img.Set(a.X, q, col)
		}
	// horizontal
	case a.Y == b.Y:
		x1, x2 := a.X, b.X
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for q := x1; q <= x2; q++ {
			img.Set(q, a.Y, col)
		}
	default:
		panic("input line is not parallel to axis. not implemented")
	}
}

func addLabel(img *image.RGBA, x, y int, label string) {
	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
