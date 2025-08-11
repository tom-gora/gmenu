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

var passedPickerCmdString string

type Conf struct {
	Picker PickerCommand
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

	flag.StringVar(&passedPickerCmdString, "p", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout.")
	flag.StringVar(&passedPickerCmdString, "picker-command", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout.")

	if passedPickerCmdString == "" {
		flag.Usage()
		return nil, errWithUsage
	}
	flag.Parse()

	parts := strings.Split(passedPickerCmdString, " ")

	picker.Cmd = parts[0]
	picker.Args = parts[1:]

	cfg.Picker = *picker

	return cfg, nil
}
