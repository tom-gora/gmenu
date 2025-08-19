package cli

import "flag"

// runtime storage for flag values
var (
	isHelp bool
	// menu
	passedMenuConfigFile string
	passedResult         string
	passedDefaultExec    string
	isJoin               bool
	// colors
	passedPickerCmdString  string
	passedClipManCmdString string
	isThumbnailRequired    bool
)

func setupFlags() {
	// help flags for any stage
	flag.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	flag.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")

	menuSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	pickSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	pickSubcmd.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")
	shadesSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	shadesSubcmd.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")
	paletteSubcmd.BoolVar(&isHelp, "h", false, "Print this message and exit [boolean]")
	paletteSubcmd.BoolVar(&isHelp, "help", false, "Print this message and exit [boolean]")

	// flags for menu subcmd
	menuSubcmd.StringVar(&passedMenuConfigFile, "m", "", "Path to menu config file in JSON format. [string]")
	menuSubcmd.StringVar(&passedMenuConfigFile, "menu-file", "", "Path to menu config file in JSON format. [string]")
	menuSubcmd.StringVar(&passedResult, "r", "", "Path to file where result of menu selection will be written. [string]")
	menuSubcmd.StringVar(&passedResult, "result", "", "Path to file where result of menu selection will be written. [string]")
	menuSubcmd.StringVar(&passedDefaultExec, "e", "", "Default command to execute if no selection is made. [string]")
	menuSubcmd.StringVar(&passedDefaultExec, "default-exec", "", "Default command to execute if no selection is made. [string]")
	menuSubcmd.BoolVar(&isJoin, "j", false, "Join menu items with new line character. [boolean]")
	menuSubcmd.BoolVar(&isJoin, "join", false, "Join menu items with new line character. [boolean]")

	// flags for color subcmds
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
