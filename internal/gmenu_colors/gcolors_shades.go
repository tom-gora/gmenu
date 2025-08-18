package gcolors

import (
	"slices"

	"github.com/lucasb-eyer/go-colorful"
)

var black = colorful.Color{
	R: 0,
	G: 0,
	B: 0,
}

var white = colorful.Color{
	R: 1,
	G: 1,
	B: 1,
}

func getRatios(count uint32) []float64 {
	if count == 0 {
		return []float64{}
	}
	r := 1.0 / (float64(count) + 1.0)

	ratios := make([]float64, count)

	var base float64 = 0

	for i := uint32(0); i < count; i++ {
		base += r
		ratios[i] = base
	}
	return ratios
}

func getShades(hex string, count uint32) []colorful.Color {
	ratios := getRatios(count)
	c, _ := colorful.Hex(hex)
	results := make([]colorful.Color, count)
	for i, r := range ratios {
		results[i] = c.BlendLab(black, r)
	}
	return results
}

func getTints(hex string, count uint32) []colorful.Color {
	ratios := getRatios(count)
	c, _ := colorful.Hex(hex)
	results := make([]colorful.Color, count)
	for i, r := range ratios {
		results[i] = c.BlendLab(white, r)
	}
	return results
}

func GetShadesAndTintsSpectrum(hex string, offset uint32) []string {
	shades := getShades(hex, 10)
	tints := getTints(hex, 10)

	slices.Reverse(tints)

	results := make([]string, offset*2+1)

	for i, t := range tints {
		results[i] = t.Hex()
	}
	results[offset] = hex

	for i, s := range shades {
		results[(i+int(offset))+1] = s.Hex()
	}
	return results
}
