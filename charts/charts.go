package charts

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"time"

	"golang.org/x/image/font"
)

// ColorScale is interface for extracting color based on value from 0 to 1
type ColorScale interface {
	GetColor(val float64) color.RGBA
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
	Counts             map[string]int
	ColorScale         ColorScale
	DrawMonthSeparator bool
	DrawLabels         bool
	BoxSize            int
	Margin             int
	TextWidthLeft      int
	TextHeightTop      int
	TextColor          color.RGBA
	BorderColor        color.RGBA
	Locale             string
	Format             string
	FontFace           font.Face
}

// WriteHeatmap writes image with heatmap and additional elements
func WriteHeatmap(conf HeatmapConfig, w io.Writer) error {
	if conf.Format == "svg" {
		writeSVG(conf, w)
	}

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
				Width:   5,
				Color:   conf.BorderColor,
			},
		)
	}

	locale := "en_US"
	if conf.Locale != "" {
		locale = conf.Locale
	}
	labelsProvider := NewLabelsProvider(locale)

	if conf.DrawLabels {
		visitors = append(visitors, &MonthLabelsVisitor{FontFace: conf.FontFace, Img: img, YOffset: 50, Color: conf.TextColor, LabelsProvider: labelsProvider})
	}

	for iter := NewDayIterator(conf.Counts, offset, conf.BoxSize, conf.Margin); !iter.Done(); iter.Next() {
		for _, v := range visitors {
			v.Visit(iter)
		}
	}

	if conf.DrawLabels {
		drawWeekdayLabels(
			conf.FontFace,
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
			labelsProvider,
		)
	}

	switch conf.Format {
	case "png":
		if err := png.Encode(w, img); err != nil {
			return err
		}
	case "jpeg":
		if err := jpeg.Encode(w, img, nil); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(w, img, nil); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unexpected format")
	}

	return nil
}

// DayVisitor is interface to update image based on current box
type DayVisitor interface {
	Visit(iter *DayIterator)
}

// DayBoxVisitor draws signle heatbox
type DayBoxVisitor struct {
	Img        *image.RGBA
	ColorScale ColorScale
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

		xL := p.X - marginSep - d.Width/2
		xR := p.X + d.BoxSize + marginSep

		// left vertical line
		draw.Draw(
			d.Img,
			image.Rect(xL, p.Y, xL+d.Width, d.MaxY),
			&image.Uniform{d.Color},
			image.ZP,
			draw.Src,
		)
		if day.Weekday() != weekdayOrder[0] {
			// right vertical line
			draw.Draw(
				d.Img,
				image.Rect(xR, d.MinY, xR+d.Width, p.Y-marginSep),
				&image.Uniform{d.Color},
				image.ZP,
				draw.Src,
			)
			// horizontal line
			draw.Draw(
				d.Img,
				image.Rect(xL, p.Y-marginSep, xR+d.Width, p.Y-marginSep-d.Width),
				&image.Uniform{d.Color},
				image.ZP,
				draw.Src,
			)
			// connect left vertical line and horizontal one
			draw.Draw(
				d.Img,
				image.Rect(xL, p.Y-marginSep-d.Width, xL+d.Width, p.Y),
				&image.Uniform{d.Color},
				image.ZP,
				draw.Src,
			)
		}
	}
}

// MonthLabelsVisitor draws month label on top of first row 0 of month
type MonthLabelsVisitor struct {
	Img            *image.RGBA
	YOffset        int
	Color          color.RGBA
	LabelsProvider LabelsProvider
	FontFace       font.Face
}

// Visit on every iteration
func (d *MonthLabelsVisitor) Visit(iter *DayIterator) {
	day := iter.Time()
	// Note, day is from 1~31
	if iter.Row == 0 && day.Day() <= 7 {
		p := iter.Point()
		drawText(
			d.FontFace,
			d.Img,
			image.Point{X: p.X, Y: p.Y - d.YOffset},
			d.LabelsProvider.GetMonth(day.Month()),
			d.Color,
		)
	}
}

// drawWeekdayLabel draws column of same width labels for weekdays
// All weekday labels assumed to have same width, which really depends on font.
// offset argument is top right corner of where to insert column of weekday labels.
func drawWeekdayLabels(fontFace font.Face, img *image.RGBA, offset image.Point, weekdays map[time.Weekday]bool, boxSize int, margin int, color color.RGBA, lp LabelsProvider) {
	width := 250
	height := 100
	y := offset.Y + height
	for _, w := range weekdayOrder {
		if weekdays[w] {
			drawText(fontFace, img, image.Point{X: offset.X - width, Y: y}, lp.GetWeekday(w), color)
		}
		y += boxSize + margin
	}
}
