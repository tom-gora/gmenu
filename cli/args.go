// Package cli
package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// Conf - complete configuration struct
type Conf struct {
	PickConf    *PickConf
	ShadesConf  *ShadesConf
	PaletteConf *PaletteConf
}

// PickConf - config struct for pick subcommand
type PickConf struct {
	Picker    *PickerCommand
	DrawThumb *bool
}

// ShadesConf - config struct for shades subcommand
type ShadesConf struct {
	Picker  *PickerCommand
	ClipMan *ClipManCommand
}

// PaletteConf - config struct for palette subcommand
type PaletteConf struct {
	Picker  *PickerCommand
	ClipMan *ClipManCommand
}

// PickerCommand - color picker command
type PickerCommand struct {
	Cmd  string
	Args []string
}

// ClipManCommand - clipboard manager command
type ClipManCommand struct {
	Cmd  string
	Args []string
}

var defaultPickerCommand = &PickerCommand{
	Cmd:  "hyprpicker",
	Args: []string{"-n", "-f", "hex"},
}

var defaultClipManCommand = &ClipManCommand{
	Cmd:  "cliphist",
	Args: []string{"list"},
}

var (
	errWithUsage  = fmt.Errorf("usage printed")
	subcmdWarnFmt = color.New(color.FgYellow, color.Bold).SprintfFunc()
)

func ParseArgs() (*Conf, error) {
	// TODO: TO TRACK - as flags and functionality is built adjust the usage print func
	// old prototype impl kept for now
	flag.Usage = func() {
		headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
		fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s:\n", os.Args[0]))
		tbl := table.New("Subcommand", "Description", "Valid flags")
		tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
		tbl.AddRow("  pick", "Call color picker on the system to get a set of various formats for this color.", "[ -p | --picker-command ] [ -d | --draw-thumbnail ] [ -h | --help ]")
		tbl.AddRow(" ", "Outputs: hex, rgb, rgba, hsl, hsla, oklab, oklch, closest CSS named color (in RGB colorspace)", " ")
		tbl.AddRow("  shades", "Create shades of color coming either from picker or most recent stored in system clipboard manager.", "[ -p | --picker-command ] [ -c | --clipman-command ] [ -h | --help ]")
		tbl.AddRow("  palette", "Create a color palette from color coming either from picker or most recent stored in system clipboard manager.", "[ -p | --picker-command ] [ -c | --clipman-command ] [ -h | --help ]")
		tbl.Print()
	}

	// storage for flag values
	var (
		passedPickerCmdString  string
		passedClipManCmdString string
		isThumbnailRequired    bool
		isHelp                 bool
	)

	// declare subcommands
	pickSubcmd := flag.NewFlagSet("pick", flag.ExitOnError)
	shadesSubcmd := flag.NewFlagSet("shades", flag.ExitOnError)
	paletteSubcmd := flag.NewFlagSet("palette", flag.ExitOnError)

	pickSubcmd.Usage = func() {
		headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
		fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s pick:\n", os.Args[0]))
		tbl := table.New("Flag", "Type", "Description", "Default")
		tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
		tbl.AddRow("[ -h | --help          ", "boolean", "Print this message.", "")
		tbl.AddRow("[ -p | --picker-command", "string", "Command to use to call picker that must return hex color value to stdout.", "hyprpicker -n -f hex")
		tbl.AddRow("[ -d | --draw-thumbnail", "boolean", "Should a temporarty PNG thumbnail filled with picked color be created in /tmp/color_thumb.png", "true")
		tbl.Print()
	}

	shadesSubcmd.Usage = func() {
		headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
		fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s shades:\n", os.Args[0]))
		tbl := table.New("Flag", "Type", "Description", "Default")
		tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
		tbl.AddRow("[ -h | --help            ]", "boolean", "Print this message.", "")
		tbl.AddRow("[ -p | --picker-command  ]", "string", "Command to use to call picker that must return hex color value to stdout.", "hyprpicker -n -f hex")
		tbl.AddRow("[ -c | --clipman-command ]", "string", "Command to use to retrieve clipboard history.", "cliphist list")
		tbl.Print()
	}

	paletteSubcmd.Usage = func() {
		headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
		fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s palette:\n", os.Args[0]))
		tbl := table.New("Flag", "Type", "Description", "Default")
		tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
		tbl.AddRow("[ -h | --help            ]", "boolean", "Print this message.", "")
		tbl.AddRow("[ -p | --picker-command  ]", "string", "Command to use to call picker that must return hex color value to stdout.", "hyprpicker -n -f hex")
		tbl.AddRow("[ -c | --clipman-command ]", "string", "Command to use to retrieve clipboard history.", "cliphist list")
		tbl.Print()
	}

	// help flags for any stage
	flag.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	flag.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")

	pickSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	pickSubcmd.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")
	shadesSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	shadesSubcmd.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")
	paletteSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	paletteSubcmd.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")

	// flags for any subcmd
	pickSubcmd.StringVar(&passedPickerCmdString, "p", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	pickSubcmd.StringVar(&passedPickerCmdString, "picker-command", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	shadesSubcmd.StringVar(&passedPickerCmdString, "p", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	shadesSubcmd.StringVar(&passedPickerCmdString, "picker-command", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	paletteSubcmd.StringVar(&passedPickerCmdString, "p", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")
	paletteSubcmd.StringVar(&passedPickerCmdString, "picker-command", "hyprpicker -n -f hex", "Command to use to call picker that must return hex color value to stdout. [string]")

	// flags only for either shades or palette subcmds
	shadesSubcmd.StringVar(&passedClipManCmdString, "c", "cliphist list", "Command to use to retrieve clipboard history. [string]")
	shadesSubcmd.StringVar(&passedClipManCmdString, "clipman-command", "cliphist list", "Command to use to retrieve clipboard history. [string]")
	paletteSubcmd.StringVar(&passedClipManCmdString, "c", "cliphist list", "Command to use to retrieve clipboard history. [string]")
	paletteSubcmd.StringVar(&passedClipManCmdString, "clipman-command", "cliphist list", "Command to use to retrieve clipboard history. [string]")

	// optional flag, valid only for color picker
	pickSubcmd.BoolVar(&isThumbnailRequired, "d", true, "Should a temporarty PNG thumbnail filled with picked color be created in /tmp [boolean]")
	pickSubcmd.BoolVar(&isThumbnailRequired, "draw-thumbnail", true, "Should a temporarty PNG thumbnail filled with picked color be created in /tmp [boolean]")

	if len(os.Args) < 2 {
		fmt.Println(subcmdWarnFmt("Expected valid subcommand: [ pick | shades | palette ].\n"))
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "pick":
		e := pickSubcmd.Parse(os.Args[2:])
		if e != nil {
			fmt.Fprintf(os.Stderr, "error parsing pick subcommand: %v\n", e)
		}

		if isHelp {
			pickSubcmd.Usage()
			os.Exit(0)
		}

		parts := strings.Split(passedPickerCmdString, " ")

		return &Conf{
			PickConf: &PickConf{
				Picker: &PickerCommand{
					Cmd:  parts[0],
					Args: parts[1:],
				},
				DrawThumb: &isThumbnailRequired,
			},
		}, nil
	case "shades":
		e := shadesSubcmd.Parse(os.Args[2:])
		if e != nil {
			fmt.Fprintf(os.Stderr, "error parsing pick subcommand: %v\n", e)
		}

		if isHelp {
			shadesSubcmd.Usage()
			os.Exit(0)
		}

		pickerParts := strings.Split(passedPickerCmdString, " ")
		clipManParts := strings.Split(passedClipManCmdString, " ")

		return &Conf{
			ShadesConf: &ShadesConf{
				Picker: &PickerCommand{
					Cmd:  pickerParts[0],
					Args: pickerParts[1:],
				},
				ClipMan: &ClipManCommand{
					Cmd:  clipManParts[0],
					Args: clipManParts[1:],
				},
			},
		}, nil
	case "palette":
		e := paletteSubcmd.Parse(os.Args[2:])
		if e != nil {
			return nil, fmt.Errorf("error parsing palette subcommand: %w", e)
		}

		if isHelp {
			paletteSubcmd.Usage()
			os.Exit(0)
		}

		pickerParts := strings.Split(passedPickerCmdString, " ")
		clipManParts := strings.Split(passedClipManCmdString, " ")

		return &Conf{
			PaletteConf: &PaletteConf{
				Picker: &PickerCommand{
					Cmd:  pickerParts[0],
					Args: pickerParts[1:],
				},
				ClipMan: &ClipManCommand{
					Cmd:  clipManParts[0],
					Args: clipManParts[1:],
				},
			},
		}, nil

	default:
		fmt.Println(subcmdWarnFmt("Expected valid subcommand: [ pick | shades | palette ].\n"))
		flag.Usage()
		os.Exit(1)
	}
	return &Conf{}, errors.New("subcommand not implemented yet, please use 'pick' or 'shades' for now")
}

func main() {
	conf, err := ParseArgs()
	if err != nil {
		if errors.Is(err, errWithUsage) {
			os.Exit(0)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Parsed configuration: %+v\n", conf)
}
