package prettyfyne

import (
	"golang.org/x/image/colornames"
	"image/color"
	"image/color/palette"
	"math"
)

func PalletMap() map[string][]color.Color {
	p := make(map[string][]color.Color)
	p["Default"] = []color.Color{color.RGBA{R:0, G:0, B:0, A:0 }} // gets overridden by current theme value in widget
	p["Material Light"] = append(MaterialLight, MaterialTypographyLight...)
	p["Web Safe"] = palette.WebSafe
	p["Plan9"] = palette.Plan9
	p["Greyscale"] = makeGrey()
	p["SVG 1.1"] = makeSvgColors()
	return p
}

func makeGrey() []color.Color {
	c := make([]color.Color, 0)
	for i := uint8(0); i <= math.MaxUint8-1; i++ {
		c = append(c, color.RGBA{
			R: i,
			G: i,
			B: i,
			A: 255,
		})
	}
	return c
}

func makeSvgColors() []color.Color {
	c := make([]color.Color, 0)
	for _, v := range colornames.Names {
		c = append(c, colornames.Map[v])
	}
	return c
}

var MaterialLight = []color.Color{
	color.RGBA{0xa6, 0xd4, 0xfa, 0xff},
	color.RGBA{0xf6, 0xa5, 0xc0, 0xff},
	color.RGBA{0xe5, 0x73, 0x73, 0xff},
	color.RGBA{0xff, 0xb7, 0x4d, 0xff},
	color.RGBA{0x64, 0xb5, 0xf6, 0xff},
	color.RGBA{0x81,0xc7,0x84, 0xff},
}

var MaterialTypographyLight = []color.Color {
	color.RGBA{0, 0, 0, 222},
	color.RGBA{0, 0, 0, 138},
	color.RGBA{0, 0, 0, 66},
	color.RGBA{0xfa, 0xfa, 0xfa, 0xff},
	color.RGBA{0, 0, 0, 30},
	color.RGBA{0, 0, 0, 10},
	color.RGBA{0, 0, 0, 97},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}