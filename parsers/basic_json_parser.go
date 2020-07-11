package parsers

import (
	"encoding/json"
	"fmt"
	"time"
)

// BasicJSONParser is example parser for JSON encoding
type BasicJSONParser struct{}

// Parse unmarshalls data and makes necessary object
func (p *BasicJSONParser) Parse(data []byte) (int, map[int]int, error) {
	var counts map[string]int

	if err := json.Unmarshal(data, &counts); err != nil {
		return 0, nil, fmt.Errorf("can not unmarshal data: %w", err)
	}

	if len(counts) == 0 {
		return 0, nil, fmt.Errorf("object is empty")
	}

	var year int
	countByDay := make(map[int]int, 366)

	for dayString, count := range counts {
		date, err := time.Parse("2006-01-02", dayString)
		if err != nil {
			return 0, nil, fmt.Errorf("can not parse time: %w", err)
		}
		countByDay[date.YearDay()] += count
		year = date.Year()
	}

	return year, countByDay, nil
}
