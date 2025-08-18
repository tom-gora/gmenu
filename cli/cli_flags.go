package cli

import "flag"

// runtime storage for flag values
var (
	passedPickerCmdString  string
	passedClipManCmdString string
	isThumbnailRequired    bool
	isHelp                 bool
)

func setupFlags() {
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
	pickSubcmd.StringVar(&passedPickerCmdString, "p", defaultPickerCmd, "Command to use to call picker that must return hex color value to stdout. [string]")
	pickSubcmd.StringVar(&passedPickerCmdString, "picker-command", defaultPickerCmd, "Command to use to call picker that must return hex color value to stdout. [string]")
	shadesSubcmd.StringVar(&passedPickerCmdString, "p", defaultPickerCmd, "Command to use to call picker that must return hex color value to stdout. [string]")
	shadesSubcmd.StringVar(&passedPickerCmdString, "picker-command", defaultPickerCmd, "Command to use to call picker that must return hex color value to stdout. [string]")
	paletteSubcmd.StringVar(&passedPickerCmdString, "p", defaultPickerCmd, "Command to use to call picker that must return hex color value to stdout. [string]")
	paletteSubcmd.StringVar(&passedPickerCmdString, "picker-command", defaultPickerCmd, "Command to use to call picker that must return hex color value to stdout. [string]")

	// flags only for either shades or palette subcmds
	shadesSubcmd.StringVar(&passedClipManCmdString, "c", defaultClipCmd, "Command to use to retrieve clipboard history. [string]")
	shadesSubcmd.StringVar(&passedClipManCmdString, "clipman-command", defaultClipCmd, "Command to use to retrieve clipboard history. [string]")
	paletteSubcmd.StringVar(&passedClipManCmdString, "c", defaultClipCmd, "Command to use to retrieve clipboard history. [string]")
	paletteSubcmd.StringVar(&passedClipManCmdString, "clipman-command", defaultClipCmd, "Command to use to retrieve clipboard history. [string]")

	// optional flag, valid only for color picker
	pickSubcmd.BoolVar(&isThumbnailRequired, "d", true, "Should a temporarty PNG thumbnail filled with picked color be created in /tmp [boolean]")
	pickSubcmd.BoolVar(&isThumbnailRequired, "draw-thumbnail", true, "Should a temporarty PNG thumbnail filled with picked color be created in /tmp [boolean]")
}
