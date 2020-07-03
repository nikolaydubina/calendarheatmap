package charts

import (
	"image"
	"testing"
)

func TestBasicDayIter(t *testing.T) {
	t.Run("num days correct", func(t *testing.T) {
		iter := NewDayIterator(
			2019,
			image.Point{X: 0, Y: 0},
			map[int]int{},
			5,
			3,
		)
		if iter == nil {
			t.Errorf("should not be nil")
		}
		if iter.Done() {
			t.Errorf("should not be done on start")
		}
		cnt := 1
		for ; !iter.Done(); iter.Next() {
			cnt++
		}
		cnt = cnt - 1
		if cnt != 365 {
			t.Errorf("2019 has 365 days, got %d", cnt)
		}
		if iter.Time().YearDay() != 1 || iter.Time().Year() != 2020 {
			t.Errorf("has to be day 1 of next year")
		}
	})

	t.Run("num days correct, leap year", func(t *testing.T) {
		iter := NewDayIterator(
			2000,
			image.Point{X: 0, Y: 0},
			map[int]int{},
			5,
			3,
		)
		if iter == nil {
			t.Errorf("should not be nil")
		}
		if iter.Done() {
			t.Errorf("should not be done on start")
		}
		cnt := 1
		for ; !iter.Done(); iter.Next() {
			cnt++
		}
		cnt = cnt - 1
		if cnt != 366 {
			t.Errorf("2000 has 366 days, got %d", cnt)
		}
		if iter.Time().YearDay() != 1 || iter.Time().Year() != 2001 {
			t.Errorf("has to be day 1 of next year")
		}
	})

	t.Run("value check, float", func(t *testing.T) {
		iter := NewDayIterator(
			2000,
			image.Point{X: 0, Y: 0},
			map[int]int{2: 2, 5: 1},
			5,
			3,
		)
		for ; !iter.Done(); iter.Next() {
			var exp float64
			switch iter.Time().YearDay() {
			case 2:
				exp = 1
			case 5:
				exp = 0.5
			}
			if iter.Value() != exp {
				t.Errorf("wrong day value")
			}
		}
	})

	t.Run("value check, empty counters", func(t *testing.T) {
		iter := NewDayIterator(
			2000,
			image.Point{X: 0, Y: 0},
			map[int]int{},
			5,
			3,
		)
		for ; !iter.Done(); iter.Next() {
			if iter.Value() != 0 {
				t.Errorf("wrong day value")
			}
		}
	})

	t.Run("value check, nil counters", func(t *testing.T) {
		iter := NewDayIterator(
			2000,
			image.Point{X: 0, Y: 0},
			nil,
			5,
			3,
		)
		for ; !iter.Done(); iter.Next() {
			if iter.Value() != 0 {
				t.Errorf("wrong day value")
			}
		}
	})
}
