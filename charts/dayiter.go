package charts

import (
	"image"
	"time"
)

// DayIterator has data for updating image based on a day
type DayIterator struct {
	Year       int
	Row        int
	Col        int
	offset     image.Point
	time       time.Time
	countByDay map[int]int
	maxCount   int
	boxSize    int
	margin     int
}

// NewDayIterator initilizes iterator for a year
func NewDayIterator(year int, offset image.Point, countByDay map[int]int, boxSize int, margin int) *DayIterator {
	row := 0
	yearStartDate := time.Date(year, 1, 1, 1, 1, 1, 1, time.UTC)
	for i, w := range weekdayOrder {
		if w == yearStartDate.Weekday() {
			row = i
		}
	}
	// in case CountByDay is empty, we need to make Value 0/1 -> 0
	maxCount := 1
	for _, q := range countByDay {
		if q > maxCount {
			maxCount = q
		}
	}
	return &DayIterator{
		Year:       year,
		time:       yearStartDate,
		Col:        0,
		Row:        row,
		offset:     offset,
		countByDay: countByDay,
		maxCount:   maxCount,
		boxSize:    boxSize,
		margin:     margin,
	}
}

// Next will update current iterator to next value
func (d *DayIterator) Next() {
	if d.Row == 6 {
		d.Row = 0
		d.Col++
	} else {
		d.Row++
	}
	d.time = d.time.Add(time.Hour * 24)
}

// Done returns true if no entries left, else false
func (d *DayIterator) Done() bool {
	return d.time.Year() > d.Year
}

// Point returns position of top left corner of box for drawing
func (d *DayIterator) Point() image.Point {
	return image.Point{
		X: d.offset.X + d.Col*(d.boxSize+d.margin),
		Y: d.offset.Y + d.Row*(d.boxSize+d.margin),
	}
}

// Time returns time representation of iterator
func (d *DayIterator) Time() time.Time {
	return d.time
}

// Value returns relative value in range 0 ~ 1
func (d *DayIterator) Value() float64 {
	return float64(d.countByDay[d.time.YearDay()]) / float64(d.maxCount)
}
