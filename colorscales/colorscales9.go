package colorscales

import (
	"image/color"
	"math"
)

type ColorScale9 [9]color.RGBA

func (c ColorScale9) GetColor(val float64) color.RGBA {
	maxIdx := 8
	idx := int(math.Round(float64(maxIdx) * val))
	return c[idx]
}

var PuBu9 = ColorScale9{
	color.RGBA{255, 247, 251, 255},
	color.RGBA{236, 231, 242, 255},
	color.RGBA{208, 209, 230, 255},
	color.RGBA{166, 189, 219, 255},
	color.RGBA{116, 169, 207, 255},
	color.RGBA{54, 144, 192, 255},
	color.RGBA{5, 112, 176, 255},
	color.RGBA{4, 90, 141, 255},
	color.RGBA{2, 56, 88, 255},
}

var GnBu9 = ColorScale9{
	color.RGBA{247, 252, 240, 255},
	color.RGBA{224, 243, 219, 255},
	color.RGBA{204, 235, 197, 255},
	color.RGBA{168, 221, 181, 255},
	color.RGBA{123, 204, 196, 255},
	color.RGBA{78, 179, 211, 255},
	color.RGBA{43, 140, 190, 255},
	color.RGBA{8, 104, 172, 255},
	color.RGBA{8, 64, 129, 255},
}

var YlGn9 = ColorScale9{
	color.RGBA{255, 255, 229, 255},
	color.RGBA{247, 252, 185, 255},
	color.RGBA{217, 240, 163, 255},
	color.RGBA{173, 221, 142, 255},
	color.RGBA{120, 198, 121, 255},
	color.RGBA{65, 171, 93, 255},
	color.RGBA{35, 132, 67, 255},
	color.RGBA{0, 104, 55, 255},
	color.RGBA{0, 69, 41, 255},
}
