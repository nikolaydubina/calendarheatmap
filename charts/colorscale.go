package charts

import (
	"encoding/csv"
	"errors"
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
)

// BasicColorScale is color scale with variable number of colors
type BasicColorScale []color.RGBA

// GetColor returns color based on float value from 0 to 1
func (c BasicColorScale) GetColor(val float64) color.RGBA {
	if len(c) == 0 {
		return color.RGBA{}
	}
	maxIdx := len(c) - 1
	idx := int(math.Round(float64(maxIdx) * val))
	return c[idx]
}

func uint8FromStr(s string) (uint8, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("can not convert to int")
	}
	if v < 0 {
		return 0, errors.New("less than 0")
	}
	if v > math.MaxUint8 {
		return 0, errors.New("higher than max")
	}
	return uint8(v), nil
}

// NewBasicColorscaleFromCSV creates colorscale from reader
func NewBasicColorscaleFromCSV(reader io.Reader) (BasicColorScale, error) {
	rows, err := csv.NewReader(reader).ReadAll()
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

	colorscale := make(BasicColorScale, len(rows)-1)
	for i, vals := range rows[1:] {
		r, err := uint8FromStr(vals[colmap["R"]])
		if err != nil {
			return nil, fmt.Errorf("bad value for color: %w", err)
		}
		g, err := uint8FromStr(vals[colmap["G"]])
		if err != nil {
			return nil, fmt.Errorf("bad value for color: %w", err)
		}
		b, err := uint8FromStr(vals[colmap["B"]])
		if err != nil {
			return nil, fmt.Errorf("bad value for color: %w", err)
		}
		colorscale[i] = color.RGBA{r, g, b, 255}
	}
	return colorscale, nil
}

// NewBasicColorscaleFromCSVFile loads basic colorscale from CSV file
func NewBasicColorscaleFromCSVFile(path string) (BasicColorScale, error) {
	colorscaleReader, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can not open file: %w", err)
	}
	return NewBasicColorscaleFromCSV(colorscaleReader)
}
