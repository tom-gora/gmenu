package utils

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"rofiQuickColors/cli"

	"github.com/alltom/oklab"
)

type ColorStrings struct {
	HEX          string
	RGB          string
	RGBA         string
	HSL          string
	HSLA         string
	OKLAB        string
	OKLCH        string
	ClosestNamed string
}

func GatherColorStrings(p cli.PickConf) (ColorStrings, error) {
	COLORhex, err := captureColorFromPicker(p)
	if err != nil {
		return ColorStrings{}, err
	}

	_rgb, err := hexToRgb(COLORhex)
	if err != nil {
		return ColorStrings{}, err
	}

	_rgba := rgbToRgba(_rgb)
	_hsl := rgbToHsl(_rgb)
	_oklab := oklab.OklabModel.Convert(_rgba).(oklab.Oklab)
	_oklch := oklab.OklchModel.Convert(_rgba).(oklab.Oklch)

	COLORhsl := fmt.Sprintf("hsl(%v %v%% %v%%)", roundFloat(_hsl.H, 2), roundFloat(_hsl.S*100, 2), roundFloat(_hsl.L*100, 2))
	COLORhsla := fmt.Sprintf("hsla(%v %v%% %v%% 0.5)", roundFloat(_hsl.H, 2), roundFloat(_hsl.S*100, 2), roundFloat(_hsl.L*100, 2))
	COLORrgb := fmt.Sprintf("rgb(%v, %v, %v)", _rgba.R, _rgba.G, _rgba.B)
	COLORrgba := fmt.Sprintf("rgba(%v, %v, %v, %v)", _rgba.R, _rgba.G, _rgba.B, "0.5")
	COLORoklab := fmt.Sprintf("oklab(%v%% %v %v)", roundFloat(_oklab.L*100, 2), roundFloat(_oklab.A, 4), roundFloat(_oklab.B, 4))
	COLORoklch := fmt.Sprintf("oklch(%v%% %v%% %vdeg)", roundFloat(_oklch.L*100, 2), roundFloat(_oklch.C*250, 2), roundFloat(radToDeg(_oklch.H), 1))
	COLORclosestNamed, err := findClosestDistanceInRgbSpace(_rgb)
	if err != nil {
		return ColorStrings{}, err
	}
	return ColorStrings{
		HEX:          COLORhex,
		RGB:          COLORrgb,
		RGBA:         COLORrgba,
		HSL:          COLORhsl,
		HSLA:         COLORhsla,
		OKLAB:        COLORoklab,
		OKLCH:        COLORoklch,
		ClosestNamed: COLORclosestNamed,
	}, nil
}

func DrawTmpThumbnail(hex string) error {
	s := 5
	img := image.NewRGBA(image.Rect(0, 0, s, s))
	rgb, err := hexToRgb(hex)
	if err != nil {
		return err
	}

	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = rgb.R
		img.Pix[i+1] = rgb.G
		img.Pix[i+2] = rgb.B
		img.Pix[i+3] = 255
	}

	tempPath := fmt.Sprintf("%s/color_thumb.png", os.TempDir())

	f, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer f.Close()

	png.Encode(f, img)
	return nil
}

func OutputAsLines(colorStrings ColorStrings) {
	fmt.Printf("%v\n", colorStrings.HEX)
	fmt.Printf("%v\n", colorStrings.RGB)
	fmt.Printf("%v\n", colorStrings.RGBA)
	fmt.Printf("%v\n", colorStrings.HSL)
	fmt.Printf("%v\n", colorStrings.HSLA)
	fmt.Printf("%v\n", colorStrings.OKLAB)
	fmt.Printf("%v\n", colorStrings.OKLCH)
	fmt.Printf("%v\n", colorStrings.ClosestNamed)
}

func outputAsJsonArr() {}

func outputAsJsonObj() {}

func OutputShadesAndTints(shadesAndTints []string) {
	for _, v := range shadesAndTints {
		fmt.Printf("%v\n", v)
	}
}
