package main

import (
	"fmt"
	"gmenu/cli"
	"os"
	"os/exec"
	"strings"

	gcolors "gmenu/internal/gmenu_colors"
)

func handleMenu(menuEntries []cli.MenuEntry, isJoin *bool) {
	for _, menuEntry := range menuEntries {
		if menuEntry.Name == "" {
			continue
		}

		if *isJoin {
			menuEntry.Name = fmt.Sprintf("%s %s", menuEntry.Value, menuEntry.Name)
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
	colorStrings, err := gcolors.GatherColorStrings(*conf)
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

	gcolors.OutputAsLines(colorStrings)
}
