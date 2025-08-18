package gcolors

import (
	_ "embed"
	"encoding/json"
	"math"
)

//go:embed data/colors_named.json
var colorsData string

type ColorKvPair struct {
	NAME string
	RGB  [3]int
}

type DistanceColorNameKv struct {
	DISTANCE float64
	NAME     string
}

func parseNamedColorsData() ([]ColorKvPair, error) {
	var tempColorData []map[string][3]int

	err := json.Unmarshal([]byte(colorsData), &tempColorData)
	if err != nil {
		return []ColorKvPair{}, err
	}

	parsedData := make([]ColorKvPair, 0, len(tempColorData))
	for _, colorMap := range tempColorData {
		for name, rgb := range colorMap {
			parsedData = append(parsedData, ColorKvPair{NAME: name, RGB: rgb})
		}
	}

	return parsedData, nil
}

func findClosestDistanceInRgbSpace(inputRGB RGB) (string, error) {
	colors, err := parseNamedColorsData()
	if err != nil {
		return "", err
	}
	best := DistanceColorNameKv{DISTANCE: 1000, NAME: ""}

	for _, color := range colors {

		distance := math.Sqrt(
			math.Pow(float64(inputRGB.R)-float64(color.RGB[0]), 2.0) +
				math.Pow(float64(inputRGB.G)-float64(color.RGB[1]), 2.0) +
				math.Pow(float64(inputRGB.B)-float64(color.RGB[2]), 2.0),
		)
		if distance < best.DISTANCE {
			best.DISTANCE = distance
			best.NAME = color.NAME
		}
	}
	return best.NAME, nil
}
