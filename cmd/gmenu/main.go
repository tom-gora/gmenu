package main

import (
	"fmt"
	"gmenu/cli"
	gcolors "gmenu/internal/gmenu_colors"
	"os"
)

func main() {
	conf, err := cli.ParseArgs()
	if err != nil {
		fmt.Printf("An error parsing the command arguments occured: %v", err)
		os.Exit(1)
	}

	colorStrings, err := gcolors.GatherColorStrings(*conf.PickConf)
	if err != nil {
		fmt.Printf("Error collecting color strings: %v", err)
		os.Exit(1)
	}
	if *conf.PickConf.DrawThumb {
		err = gcolors.DrawTmpThumbnail(colorStrings.HEX)
		if err != nil {
			fmt.Printf("Error writing thumbnail: %v", err)
			os.Exit(1)
		}
	}

	gcolors.OutputAsLines(colorStrings)
}
