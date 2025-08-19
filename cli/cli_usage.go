package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func topLevelUsage() {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s:\n", os.Args[0]))
	tbl := table.New("Subcommand", "Description")
	tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
	tbl.AddRow("  menu", "Build arbitrary menu from config defined as json.")
	tbl.AddRow("", "[ -m | --menu-file ] [ -r | --result ] [ -e | --default-exec ] [ -j | --join ] [ -h | --help ]")
	tbl.AddRow("", "")
	tbl.AddRow("  pick", "Call color picker on the system to get a set of various formats for this color.")
	tbl.AddRow(" ", "Outputs: hex, rgb, rgba, hsl, hsla, oklab, oklch, closest CSS named color (in RGB colorspace)", " ")
	tbl.AddRow("", "[ -p | --picker-command ] [ -d | --draw-thumbnail ] [ -h | --help ]")
	tbl.AddRow("", "")
	tbl.AddRow("  shades", "Create shades of color coming either from picker or most recent stored in system clipboard manager.")
	tbl.AddRow("", "[ -p | --picker-command ] [ -c | --clipman-command ] [ -h | --help ]")
	tbl.AddRow("", "")
	tbl.AddRow("  palette", "Create a color palette from color coming either from picker or most recent stored in system clipboard manager.")
	tbl.AddRow("", "[ -p | --picker-command ] [ -c | --clipman-command ] [ -h | --help ]")

	tbl.Print()
}

func pickUsage() {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s pick:\n", os.Args[0]))
	tbl := table.New("Flag", "Type", "Description", "Default")
	tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
	tbl.AddRow("[ -h | --help           ]", "boolean", "Print this message.", "")
	tbl.AddRow("[ -p | --picker-command ]", "string", "Command to use to call picker that must return hex color value to stdout.", defaultPickerCmd)
	tbl.AddRow("[ -d | --draw-thumbnail ]", "boolean", "Should a temporarty PNG thumbnail filled with picked color be created in /tmp/color_thumb.png", "true")
	tbl.Print()
}

func shadesUsage() {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s shades:\n", os.Args[0]))
	tbl := table.New("Flag", "Type", "Description", "Default")
	tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
	tbl.AddRow("[ -h | --help            ]", "boolean", "Print this message.", "")
	tbl.AddRow("[ -p | --picker-command  ]", "string", "Command to use to call picker that must return hex color value to stdout.", defaultPickerCmd)
	tbl.AddRow("[ -c | --clipman-command ]", "string", "Command to use to retrieve clipboard history.", defaultClipCmd)
	tbl.Print()
}

func paletteUsage() {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s palette:\n", os.Args[0]))
	tbl := table.New("Flag", "Type", "Description", "Default")
	tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
	tbl.AddRow("[ -h | --help            ]", "boolean", "Print this message.", "")
	tbl.AddRow("[ -p | --picker-command  ]", "string", "Command to use to call picker that must return hex color value to stdout.", defaultPickerCmd)
	tbl.AddRow("[ -c | --clipman-command ]", "string", "Command to use to retrieve clipboard history.", defaultClipCmd)
	tbl.Print()
}

func menuUsage() {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	fmt.Fprintf(os.Stderr, "%s", headerFmt("Usage of %s menu:\n", os.Args[0]))
	tbl := table.New("Flag", "Type", "Description", "Default")
	tbl.WithHeaderFormatter(color.New(color.FgCyan).SprintfFunc())
	tbl.AddRow("[ -h | --help          ]", "boolean", "Print this message.", "")
	tbl.AddRow("[ -m | --menu-file     ]", "string", "Path to JSON file with menu configuration.", "")
	tbl.AddRow("[ -r | --result        ]", "string", "Use this to pass result string from your fuzzy picker.", "")
	tbl.AddRow("[ -e | --default-exec  ]", "string", "Optionally pass command to prepend to result before running if it is the same in every case.", "wl-copy")
	tbl.AddRow("[ -j | --join          ]", "boolean", "Join items with this string in the result file.", "")
	tbl.Print()
}
