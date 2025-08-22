package gcolors

import "github.com/lucasb-eyer/go-colorful"

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

	for i := range count {
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
		results[i] = c.BlendRgb(black, r)
	}
	return results
}

func getTints(hex string, count uint32) []colorful.Color {
	ratios := getRatios(count)
	c, _ := colorful.Hex(hex)
	results := make([]colorful.Color, count)
	for i, r := range ratios {
		results[i] = c.BlendRgb(white, r)
	}
	return results
}

func GetShadesStrings(hex string) (ShadesStrings, error) {
	shades := getShades(hex, 10)
	tints := getTints(hex, 10)

	return ShadesStrings{
		Dark10:  shades[9].Hex(),
		Dark9:   shades[8].Hex(),
		Dark8:   shades[7].Hex(),
		Dark7:   shades[6].Hex(),
		Dark6:   shades[5].Hex(),
		Dark5:   shades[4].Hex(),
		Dark4:   shades[3].Hex(),
		Dark3:   shades[2].Hex(),
		Dark2:   shades[1].Hex(),
		Dark1:   shades[0].Hex(),
		Input:   hex,
		Light1:  tints[0].Hex(),
		Light2:  tints[1].Hex(),
		Light3:  tints[2].Hex(),
		Light4:  tints[3].Hex(),
		Light5:  tints[4].Hex(),
		Light6:  tints[5].Hex(),
		Light7:  tints[6].Hex(),
		Light8:  tints[7].Hex(),
		Light9:  tints[8].Hex(),
		Light10: tints[9].Hex(),
	}, nil
}
