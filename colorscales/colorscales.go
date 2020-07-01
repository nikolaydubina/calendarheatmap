package colorscales

import (
	"image/color"
)

// ColorScale is interface for extracting color from float
type ColorScale interface {
	GetColor(val float64) color.RGBA
}

// LoadColorScale loads colorscale struct based on associated string name
func LoadColorScale(name string) ColorScale {
	switch name {
	case "PuBu9":
		return PuBu9
	case "GnBu9":
		return GnBu9
	case "YlGn9":
		return YlGn9
	default:
		panic("unknown colorscale " + name)
	}
}
