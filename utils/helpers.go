// Package utils
package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"regexp"
	"rofiQuickColors/cli"
	"strings"
)

func roundFloat(val float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	rounded := math.Round(val * shift)
	return rounded / shift
}

func ensureValidHexColor(out string) string {
	pattern := regexp.MustCompile(`#([A-Fa-f0-9]{6})`)
	scanner := bufio.NewScanner(strings.NewReader(out))

	for scanner.Scan() {
		lineText := scanner.Text()
		matches := pattern.FindAllStringSubmatch(lineText, -1)
		if len(matches) > 0 && len(matches[0]) > 1 {
			return matches[0][0]
		}
	}

	if scanner.Err() != nil {
		return ""
	}

	return ""
}

func captureColorFromPicker(p cli.PickConf) (string, error) {
	pickerCmd := p.Picker.Cmd
	pickerArgs := p.Picker.Args
	execPicker := exec.Command(pickerCmd, pickerArgs...)
	var stdout bytes.Buffer

	execPicker.Stdout = &stdout

	err := execPicker.Run()
	if err != nil {
		return "", err
	}
	outstring := stdout.String()

	// || strings.Contains(strings.ToLower(outstring), "null") {

	validatedHex := ensureValidHexColor(outstring)

	if validatedHex == "" {
		fmt.Println("Cancelled by user")
		os.Exit(130)
	}

	return validatedHex, nil
}
