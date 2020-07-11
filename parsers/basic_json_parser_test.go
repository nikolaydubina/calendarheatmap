package parsers

import (
	"reflect"
	"testing"
)

func Test_BasicJSONParser_Basic(t *testing.T) {

	t.Run("empty map", func(t *testing.T) {
		inputData := []byte(`{}`)

		parser := BasicJSONParser{}
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
		inputData := []byte(`{
			"2020-01-02": 7,
			"2020-02-05": 2,
			"2020-02-04": 3
		}`)

		expectedCountByDay := map[int]int{
			2:  7,
			36: 2,
			35: 3,
		}

		parser := BasicJSONParser{}
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

func Test_BasicJSONParser_Errors(t *testing.T) {

	t.Run("empty", func(t *testing.T) {
		inputData := []byte(``)

		parser := BasicJSONParser{}
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
		parser := BasicJSONParser{}
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

	t.Run("bad format: bad json", func(t *testing.T) {
		inputData := []byte(`{
			"2020-01-02: 7,
			"2020-02-05": 2,
			"2020-02-04": 3
		}`)
		parser := BasicJSONParser{}
		year, _, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
	})

	t.Run("bad format: wrong time format", func(t *testing.T) {
		inputData := []byte(`{
			"2020/01/02": 7
		}`)
		parser := BasicJSONParser{}
		year, _, err := parser.Parse(inputData)
		if err == nil {
			t.Fatal("expect error")
		}
		if year != 0 {
			t.Fatalf("wrong year %d", year)
		}
	})
}
