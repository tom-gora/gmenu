package main

import (
	"fmt"
	"gmenu/cli"
	"os"
	"os/exec"
	"strings"

	u "gmenu/internal"
	gcolors "gmenu/internal/gmenu_colors"
)

func handleMenu(menuEntries []cli.MenuEntry, isJoin *bool, returnValue *bool) {
	for _, menuEntry := range menuEntries {
		if menuEntry.Name == "" {
			continue
		}

		if *isJoin {
			menuEntry.Name = fmt.Sprintf("%s %s", menuEntry.Value, menuEntry.Name)
		}

		if *returnValue {
			menuEntry.Name = menuEntry.Value
		}

		icon := menuEntry.Icon
		if icon == "" {
			fmt.Printf("%s\n", menuEntry.Name)
		} else {
			fmt.Printf("%s\x00icon\x1f%s\n", menuEntry.Name, icon)
		}
	}
	os.Exit(0)
}

func handleMenuResult(result string, menuEntries []cli.MenuEntry, defaultExec *string) {
	var selectedEntry *cli.MenuEntry

	for i := range menuEntries {
		if strings.Contains(result, menuEntries[i].Name) {
			selectedEntry = &menuEntries[i]
			break
		}
	}

	if selectedEntry == nil {
		fmt.Fprintf(os.Stderr, "Error: Task '%s' not found in config file. No command to execute.\n", result)
		os.Exit(1)
	}

	var commandToExecute string
	if *defaultExec != "" {
		commandToExecute = fmt.Sprintf("%s %s", *defaultExec, selectedEntry.Value)
	} else {
		commandToExecute = selectedEntry.Value
	}

	cmd := exec.Command("bash", "-c", commandToExecute)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting command '%s': %v\n", commandToExecute, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func handlePick(conf *cli.PickConf) {
	colorStrings, err := gcolors.GatherColorStringsFromPicker(*conf.Picker)
	if err != nil {
		fmt.Printf("Error collecting color strings: %v", err)
		os.Exit(1)
	}
	if *conf.DrawThumb {
		err = gcolors.DrawTmpThumbnail(colorStrings.HEX)
		if err != nil {
			fmt.Printf("Error writing thumbnail: %v", err)
			os.Exit(1)
		}
	}
	u.PrintLines(colorStrings)
}

func handleShades(conf *cli.ShadesConf) {
	var c string
	var e error
	if *conf.UseClipMan {
		c, e = gcolors.GetLastColorFromClipboard()
		if e != nil {
			fmt.Printf("Error getting last color from clipboard: %v", e)
			os.Exit(1)
		}
	} else {
		c, e = gcolors.GetHexColorFromPicker(*conf.Picker)
		if e != nil {
			fmt.Printf("Error getting hex color from picker: %v", e)
			os.Exit(1)
		}
	}

	shadesStrings, err := gcolors.GetShadesStrings(c)
	if err != nil {
		fmt.Printf("Error getting shades strings: %v", err)
		os.Exit(1)
	}

	u.PrintLines(shadesStrings)
}
