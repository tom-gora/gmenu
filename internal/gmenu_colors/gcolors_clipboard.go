package gcolors

import (
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strings"

	"github.com/crazy3lf/colorconv"
)

func extractColor(line string, re *regexp.Regexp) string {
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func GetLastColorFromClipboard() (string, error) {
	cmdStr := "cliphist list | head -n50"
	cmd := exec.Command("sh", "-c", cmdStr)

	output, _ := cmd.Output()

	patterns := []string{
		`#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})`,
		`rgb\(\s*\d+\s*,\s*\d+\s*,\s*\d+\s*\)`,
		`rgba\(\s*\d+\s*,\s*\d+\s*,\s*\d+\s*,\s*[\d.]+\s*\)`,
		`hsl\(\s*[\d.]+\s*[\s,]+[\d.]+%\s*[\s,]+[\d.]+%\s*\)`,
		`hsla\(\s*[\d.]+\s*[\s,]+[\d.]+%\s*[\s,]+[\d.]+%\s*[\s,]+[\d.]+\s*\)`,
		`oklab\(\s*[\d.]+%\s*[\s,]+[-]?[\d.]+\s*[\s,]+[-]?[\d.]+\s*\)`,
		`oklch\(\s*[\d.]+%\s*[\s,]+[-]?[\d.]+%\s*[\s,]+[-]?[\d.]+deg\s*\)`,
	}
	foundColorStr := ""
	scanner := bufio.NewScanner(bytes.NewReader(output))
	var colorRegexes []*regexp.Regexp
	for _, pattern := range patterns {
		colorRegexes = append(colorRegexes, regexp.MustCompile(pattern))
	}
	for scanner.Scan() {
		line := scanner.Text()
		for _, colorRegex := range colorRegexes {
			if colorRegex.MatchString(line) {
				foundColorStr = extractColor(line, colorRegex)
				if foundColorStr != "" { // inner loop
					break
				}
			}
		}
		if foundColorStr != "" { // outer loop
			switch {
			case strings.HasPrefix(foundColorStr, "#"):
				return foundColorStr, nil
			case strings.HasPrefix(foundColorStr, "rgb(") || strings.HasPrefix(foundColorStr, "rgba("):
				r, g, b := parseRgbStringBackToRGB(foundColorStr)
				hex := colorconv.RGBToHex(r, g, b)
				return strings.ReplaceAll(hex, "0x", "#"), nil
			case strings.HasPrefix(foundColorStr, "hsl(") || strings.HasPrefix(foundColorStr, "hsla("):
				r, g, b := parseHslStringBackToRGB(foundColorStr)
				hex := colorconv.RGBToHex(r, g, b)
				return strings.ReplaceAll(hex, "0x", "#"), nil
			case strings.HasPrefix(foundColorStr, "oklab("):
				r, g, b := parseOklabStringBackToRGB(foundColorStr)
				hex := colorconv.RGBToHex(r, g, b)
				return strings.ReplaceAll(hex, "0x", "#"), nil
			case strings.HasPrefix(foundColorStr, "oklch("):
				r, g, b := parseOklchStringBackToRGB(foundColorStr)
				hex := colorconv.RGBToHex(r, g, b)
				return strings.ReplaceAll(hex, "0x", "#"), nil
			default:
				return "", errors.New("no color / unsupported color format found in clipboard")
			}
		}
	}
	return "", errors.New("no color / unsupported color format found in clipboard")
}
