// Package cli
package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var errWithUsage = fmt.Errorf("usage printed")

type PickerCommand struct {
	Cmd  string
	Args []string
}

var (
	passedPickerCmdString string
	isThumbnailRequired   bool
)

type Conf struct {
	Picker    *PickerCommand
	DrawThumb *bool
}

func ParseArgs() (*Conf, error) {
	picker := &PickerCommand{}
	cfg := &Conf{}

	flag.Usage = func() {
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		fmt.Fprintf(os.Stderr, "\n%s\n\n", headerFmt("Usage of %s:", os.Args[0]))
		tbl := table.New("Flag", "Type", "Description", "Default")
		tbl.WithHeaderFormatter(color.New(color.FgYellow).SprintfFunc())
		tbl.AddRow("-p, --picker-command", "string", "Command to use to call picker that must return hex color value to stdout.", "(default: hyprpicker)")
		tbl.Print()
	}

	flag.StringVar(&passedPickerCmdString, "p", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	flag.StringVar(&passedPickerCmdString, "picker-command", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	flag.BoolVar(&isThumbnailRequired, "d", true, "Should a temporarty PNG thumbnail filled with picked color be created in /tmp [boolean]")
	flag.BoolVar(&isThumbnailRequired, "draw-thumbnail", true, "Should a temporarty PNG thumbnail filled with picked color be created in /tmp [boolean]")

	if passedPickerCmdString == "" {
		flag.Usage()
		return nil, errWithUsage
	}
	flag.Parse()

	parts := strings.Split(passedPickerCmdString, " ")

	picker.Cmd = parts[0]
	picker.Args = parts[1:]

	cfg.Picker = picker
	if isThumbnailRequired {
		cfg.DrawThumb = &isThumbnailRequired
	}

	return cfg, nil
}
