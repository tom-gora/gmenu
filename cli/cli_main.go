// Package cli
package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var (
	// for flag error handling
	errWithUsage = fmt.Errorf(uErr)
	// for formatting warning prints
	subcmdWarnFmt = color.New(color.FgYellow, color.Bold).SprintfFunc()
)

func ParseArgs() (*Conf, error) {
	// assign usage functions
	flag.Usage = topLevelUsage
	pickSubcmd.Usage = pickUsage
	shadesSubcmd.Usage = shadesUsage
	paletteSubcmd.Usage = paletteUsage
	menuSubcmd.Usage = menuUsage

	setupFlags()

	if len(os.Args) < 2 {
		fmt.Println(subcmdWarnFmt("Expected valid subcommand: [ pick | shades | palette ].\n"))
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "menu":
		e := menuSubcmd.Parse(os.Args[2:])
		if e != nil {
			fmt.Fprintf(os.Stderr, "error parsing menu subcommand: %v\n", e)
		}

		if isHelp {
			menuSubcmd.Usage()
			os.Exit(0)
		}

		if passedMenuConfigFile == "" {
			return nil, fmt.Errorf("for \"%v\" subcommand [ -m | --menu-file ] flag is mandatory", os.Args[1])
		}

		if _, err := os.Stat(passedMenuConfigFile); err != nil || filepath.Ext(passedMenuConfigFile) != ".json" {
			return nil, fmt.Errorf("passed config file does not exist or is not json")
		}

		return &Conf{
			MenuConf: &MenuConf{
				MenuConfPath: &passedMenuConfigFile,
				Result:       &passedResult,
				DefaultExec:  &passedDefaultExec,
				IsJoin:       &isJoin,
			},
		}, nil

	case "pick":
		e := pickSubcmd.Parse(os.Args[2:])
		if e != nil {
			return nil, fmt.Errorf("error parsing pick subcommand")
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
			return nil, fmt.Errorf("error parsing pick subcommand")
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
			return nil, fmt.Errorf("error parsing palette subcommand")
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
	return &Conf{}, errors.New("ERROR: Expected valid subcommand: [ pick | shades | palette ].\n")
}
