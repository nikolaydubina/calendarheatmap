package parsers

import (
	"fmt"
	"strings"
	"time"
)

// row represents single row data in input file
type rowInput struct {
	Date  time.Time
	Count int
}

// RowDaySecondsCountParser is example parser for custom format
type RowDaySecondsCountParser struct{}

// Parse extracts day and count from counts of symbol
func (p *RowDaySecondsCountParser) Parse(data []byte) (int, map[int]int, error) {
	var rows []rowInput
	for _, row := range strings.Split(string(data), "\n") {
		if len(row) == 0 {
			continue
		}
		items := strings.Split(row, " ")
		if len(items) != 3 {
			return 0, nil, fmt.Errorf("number of items in row is not 3, items %#v", items)
		}
		timeString, countString := items[0]+" "+items[1], items[2]

		date, err := time.Parse("2006-01-02 15:04", timeString)
		if err != nil {
			return 0, nil, fmt.Errorf("can not parse time: %w", err)
		}
		count := strings.Count(countString, "P")
		rows = append(rows, rowInput{Date: date, Count: count})
	}

	if len(rows) == 0 {
		return 0, nil, fmt.Errorf("at least one row is required")
	}

	year := rows[0].Date.Year()
	countByDay := make(map[int]int, 366)
	for _, row := range rows {
		countByDay[row.Date.YearDay()] += row.Count
	}

	return year, countByDay, nil
}
