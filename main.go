package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"os/exec"
	"regexp"
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

func ensureHexColor(out string) (string, error) {
	pattern := regexp.MustCompile(`#([A-Fa-f0-9]{6})`)
	scanner := bufio.NewScanner(strings.NewReader(out))

	for scanner.Scan() {
		lineText := scanner.Text()
		matches := pattern.FindAllStringSubmatch(lineText, -1)
		if len(matches) > 0 && len(matches[0]) > 1 {
			return matches[0][0], nil
		}
	}

	if scanner.Err() != nil {
		return "", fmt.Errorf("error scanning lines: %v", scanner.Err())
	}

	return "", errors.New("no valid hex color code found in output")
}

func captureHexColorString() (string, error) {
	pickerCmd := "hyprpicker"
	picker := exec.Command(pickerCmd, "-n", "-f", "hex")
	var stdout bytes.Buffer

	picker.Stdout = &stdout

	err := picker.Run()
	if err != nil {
		return "", err
	}
	out := stdout.String()
	ensuredHexValue, err := ensureHexColor(out)
	if err != nil {
		return "", err
	}

	return ensuredHexValue, nil
}

func round(val float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	rounded := math.Round(val * shift)
	return rounded / shift
}

func radToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func main() {
	COLORhex, err := captureHexColorString()
	if err != nil {
		log.Fatalf("An error grabbing the color occured: %v", err)
		os.Exit(1)
	}
	_rgb, err := hexToRgb(COLORhex)
	if err != nil {
		log.Fatalf("An error converting the hex color occured: %v", err)
		os.Exit(1)
	}
	_rgba := color.RGBA{
		R: _rgb.R,
		G: _rgb.G,
		B: _rgb.B,
		A: 255,
	}

	_h, _s, _l := colorconv.RGBToHSL(_rgba.R, _rgba.G, _rgba.B)
	COLORhsl := fmt.Sprintf("hsl(%v %v%% %v%%)", round(_h, 2), round(_s*100, 2), round(_l*100, 2))

	COLORnamed, err := colorconv.HexToColor(COLORhex)
	if err != nil {
		log.Fatalf("An error converting the hex color occured: %v", err)
		os.Exit(1)
	}
	_oklab := oklab.OklabModel.Convert(_rgba).(oklab.Oklab)
	_oklch := oklab.OklchModel.Convert(_rgba).(oklab.Oklch)

	COLORrgb := fmt.Sprintf("rgb(%v, %v, %v)", _rgba.R, _rgba.G, _rgba.B)
	COLORrgba := fmt.Sprintf("rgba(%v, %v, %v, %v)", _rgba.R, _rgba.G, _rgba.B, "0.5")
	COLORoklab := fmt.Sprintf("oklab(%v%% %v %v)", round(_oklab.L*100, 2), round(_oklab.A, 4), round(_oklab.B, 4))
	COLORoklch := fmt.Sprintf("oklch(%v%% %v%% %vdeg)", round(_oklch.L*100, 2), round(_oklch.C*250, 2), round(radToDeg(_oklch.H), 1))
	fmt.Println(COLORhex)
	fmt.Println(COLORrgb)
	fmt.Println(COLORrgba)
	fmt.Println(COLORhsl)
	fmt.Println(COLORoklab)
	fmt.Println(COLORoklch)
	fmt.Println(COLORnamed)
}
