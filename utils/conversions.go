package utils

import (
	"image/color"
	"math"
	"strconv"
	"strings"

	"github.com/alltom/oklab"
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

func parseRgbStringBackToRGB(color string) (uint8, uint8, uint8) {
	color = strings.TrimPrefix(color, "rgb")
	color = strings.TrimPrefix(color, "a")
	color = strings.TrimPrefix(color, "(")
	color = strings.TrimSuffix(color, ")")

	parts := strings.Split(color, ",")
	var r, g, b uint8

	if len(parts) >= 3 {
		rInt, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		gInt, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		bInt, _ := strconv.Atoi(strings.TrimSpace(parts[2]))

		r = uint8(rInt)
		g = uint8(gInt)
		b = uint8(bInt)
	}

	return r, g, b
}

func parseHslStringBackToRGB(color string) (uint8, uint8, uint8) {
	color = strings.TrimPrefix(color, "hsl")
	color = strings.TrimPrefix(color, "a")
	color = strings.TrimPrefix(color, "(")
	color = strings.TrimSuffix(color, ")")

	parts := strings.Split(color, " ")
	var h, s, l float64

	h, _ = strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	sStr := strings.TrimSpace(parts[1])
	sStr = strings.TrimSuffix(sStr, "%")
	s, _ = strconv.ParseFloat(sStr, 64)

	lStr := strings.TrimSpace(parts[2])
	lStr = strings.TrimSuffix(lStr, "%")
	l, _ = strconv.ParseFloat(lStr, 64)

	r, g, b, _ := colorconv.HSLToRGB(h, s/100, l/100)
	return r, g, b
}

func parseOklabStringBackToRGB(color string) (uint8, uint8, uint8) {
	color = strings.TrimPrefix(color, "oklab(")
	color = strings.TrimSuffix(color, ")")

	parts := strings.Split(color, " ")
	var _l, _a, _b float64

	if len(parts) >= 3 {
		lStr := strings.TrimSpace(parts[0])
		lStr = strings.TrimSuffix(lStr, "%")
		_l, _ = strconv.ParseFloat(lStr, 64)

		_a, _ = strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		_b, _ = strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	}

	_oklab := oklab.Oklab{L: _l / 100, A: _a, B: _b}
	r32, g32, b32, _ := _oklab.RGBA()
	r := uint8(r32 / 256)
	g := uint8(g32 / 256)
	b := uint8(b32 / 256)

	return r, g, b
}

func parseOklchStringBackToRGB(color string) (uint8, uint8, uint8) {
	color = strings.TrimPrefix(color, "oklch(")
	color = strings.TrimSuffix(color, ")")

	parts := strings.Split(color, " ")
	var l, c, h float64

	if len(parts) >= 3 {
		lStr := strings.TrimSpace(parts[0])
		lStr = strings.TrimSuffix(lStr, "%")
		l, _ = strconv.ParseFloat(lStr, 64)

		cStr := strings.TrimSpace(parts[1])
		cStr = strings.TrimSuffix(cStr, "%")
		c, _ = strconv.ParseFloat(cStr, 64)

		hStr := strings.TrimSpace(parts[2])
		hStr = strings.TrimSuffix(hStr, "deg")
		h, _ = strconv.ParseFloat(hStr, 64)
	}

	_oklch := oklab.Oklch{L: l / 100, C: c / 100, H: h}
	r32, g32, b32, _ := _oklch.RGBA()
	r := uint8(r32 >> 24)
	g := uint8(g32 >> 24)
	b := uint8(b32 >> 24)

	return r, g, b
}
