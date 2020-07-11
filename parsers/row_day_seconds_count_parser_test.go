package parsers

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {

	t.Run("new lines", func(t *testing.T) {
		inputData := []byte(`

`)

		parser := RowDaySecondsCountParser{}
		year, countByDay, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
		if len(countByDay) != 0 {
			t.Fatal("no counts should be made")
		}
	})

	t.Run("basic", func(t *testing.T) {
		inputData := []byte(`
2020-01-02 17:11 P
2020-01-02 20:43 PPPPPP
2020-02-05 09:52 PP
2020-02-04 00:00 PPP
`)

		expectedCountByDay := map[int]int{
			2:  7,
			36: 2,
			35: 3,
		}

		parser := RowDaySecondsCountParser{}
		year, countByDay, err := parser.Parse(inputData)
		if err != nil {
			t.Fatal(err.Error())
		}
		if year != 2020 {
			t.Fatalf("wrong year %d", year)
		}
		if !reflect.DeepEqual(expectedCountByDay, countByDay) {
			t.Fatal("wrong countByDay")
		}
	})
}

func TestErrors(t *testing.T) {

	t.Run("empty", func(t *testing.T) {
		inputData := []byte(``)

		parser := RowDaySecondsCountParser{}
		year, countByDay, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
		if len(countByDay) != 0 {
			t.Fatal("no counts should be made")
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		parser := RowDaySecondsCountParser{}
		year, countByDay, err := parser.Parse(nil)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
		if len(countByDay) != 0 {
			t.Fatal("no counts should be made")
		}
	})

	t.Run("bad format: separator", func(t *testing.T) {
		inputData := []byte(`
2020-01-02 17:11P
2020-01-02 20:43 PPPPPP
2020-02-05 09:52 PP
2020-02-04 00:00 PPP
`)
		parser := RowDaySecondsCountParser{}
		year, _, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
	})

	t.Run("bad format: missing count", func(t *testing.T) {
		inputData := []byte(`
2020-01-02 20:43
2020-02-05 09:52 PP
2020-02-04 00:00 PPP
`)
		parser := RowDaySecondsCountParser{}
		year, _, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
	})

	t.Run("bad format: wrong time format", func(t *testing.T) {
		inputData := []byte(`
2020/01/02 20:43 PPP
`)
		parser := RowDaySecondsCountParser{}
		year, _, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
	})
}
