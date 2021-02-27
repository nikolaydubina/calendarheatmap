package charts

import (
	"encoding/csv"
	"errors"
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
)

// BasicColorScale is color scale with variable number of colors
type BasicColorScale []color.RGBA

// GetColor returns color based on float value from 0 to 1
func (c BasicColorScale) GetColor(val float64) color.RGBA {
	maxIdx := len(c) - 1
	idx := int(math.Round(float64(maxIdx) * val))
	return c[idx]
}

// NewBasicColorscaleFromCSVFile loads basic colorscale from CSV file
func NewBasicColorscaleFromCSVFile(path string) (BasicColorScale, error) {
	colorscaleReader, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can not open file: %w", err)
	}
	rows, err := csv.NewReader(colorscaleReader).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("can not read CSV: %w", err)
	}
	if len(rows) == 0 {
		return nil, errors.New("empty colorscales file")
	}

	colmap := make(map[string]int, 3)
	for i, name := range rows[0] {
		colmap[name] = i
	}
	if _, ok := colmap["R"]; !ok {
		return nil, errors.New("missing R column")
	}
	if _, ok := colmap["G"]; !ok {
		return nil, errors.New("missing G column")
	}
	if _, ok := colmap["B"]; !ok {
		return nil, errors.New("missing B column")
	}

	colorscale := make(BasicColorScale, 0)
	for _, vals := range rows[1:] {
		r, err := strconv.Atoi(vals[colmap["R"]])
		if err != nil {
			return nil, errors.New("bad value for color")
		}
		g, err := strconv.Atoi(vals[colmap["G"]])
		if err != nil {
			return nil, errors.New("bad value for color")
		}
		b, err := strconv.Atoi(vals[colmap["B"]])
		if err != nil {
			return nil, errors.New("bad value for color")
		}
		colorscale = append(colorscale, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})
	}
	return colorscale, nil
}
