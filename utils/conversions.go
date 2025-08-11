package utils

import (
	"image/color"
	"math"
	"strconv"
	"strings"

	"github.com/crazy3lf/colorconv"
)

type RGB struct {
	R uint8
	G uint8
	B uint8
}

type HSL struct {
	H float64
	S float64
	L float64
}

func radToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func hexToRgb(h string) (RGB, error) {
	noPrefixHex := strings.Replace(h, "#", "", 1)
	var rgb RGB
	values, err := strconv.ParseUint(string(noPrefixHex), 16, 32)
	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		R: uint8(values >> 16),
		G: uint8(values >> 8 & 0xFF),
		B: uint8(values & 0xFF),
	}
	return rgb, nil
}

func rgbToRgba(rgb RGB) color.RGBA {
	return color.RGBA{
		R: rgb.R,
		G: rgb.G,
		B: rgb.B,
		A: 255,
	}
}

func rgbToHsl(rgb RGB) HSL {
	_h, _s, _l := colorconv.RGBToHSL(rgb.R, rgb.G, rgb.B)
	return HSL{
		H: _h,
		S: _s,
		L: _l,
	}
}
