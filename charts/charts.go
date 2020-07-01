package charts

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/nikolaydubina/plotstats/colorscales"
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

const (
	numWeeksYear = 52
	numWeekCols  = numWeeksYear + 1 // 53 * 7 = 371 > 366
	margin       = 7                // should be odd number for best result if using month separator
	boxSize      = 25
)

var borderColor = color.RGBA{200, 200, 200, 255}

// MakeYearDayHeatmapHoriz draw every day of a year as square
// filled with color proportional to counter from the max.
func MakeYearDayHeatmapHoriz(year int, countByDay map[int]int, colorScale colorscales.ColorScale, drawMonthSeparator bool) image.Image {
	maxCount := 0
	for _, q := range countByDay {
		if q > maxCount {
			maxCount = q
		}
	}

	width := numWeekCols * (boxSize + margin)
	height := 7 * (boxSize + margin)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	x, y := 0, 0
	yearStartDate := time.Date(year, 1, 1, 1, 1, 1, 1, time.UTC)
	vIdx := weekdaysPos[yearStartDate.Weekday()]

	for day := yearStartDate; day.Year() == year; day = day.Add(time.Hour * 24) {
		y = (boxSize + margin) * vIdx

		r := image.Rect(x, y, x+boxSize, y+boxSize)
		val := float64(countByDay[day.YearDay()]) / float64(maxCount)
		color := colorScale.GetColor(val)
		draw.Draw(img, r, &image.Uniform{color}, image.ZP, draw.Src)

		if drawMonthSeparator {
			if day.Day() == 1 && day.Month() != time.January {
				marginSep := margin / 2

				closeLeft := image.Point{X: x - marginSep - 1, Y: y - marginSep - 1}
				closeRight := image.Point{X: x + boxSize + marginSep, Y: y - marginSep - 1}
				farLeft := image.Point{X: x - marginSep - 1, Y: height - margin - 1}
				farRight := image.Point{X: x + boxSize + marginSep, Y: 0}

				drawLineAxis(img, farLeft, closeLeft, borderColor) // left line
				if vIdx != 0 {
					drawLineAxis(img, closeRight, farRight, borderColor)  // right line
					drawLineAxis(img, closeLeft, closeRight, borderColor) // top line
				}
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
