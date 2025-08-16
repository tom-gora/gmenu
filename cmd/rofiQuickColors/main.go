package main

import (
	"fmt"
	"os"
	"rofiQuickColors/cli"
	u "rofiQuickColors/utils"
)

func main() {
	conf, err := cli.ParseArgs()
	if err != nil {
		fmt.Printf("An error parsing the command arguments occured: %v", err)
		os.Exit(1)
	}

	colorStrings, err := u.GatherColorStrings(*conf.PickConf)
	if err != nil {
		fmt.Printf("Error collecting color strings: %v", err)
		os.Exit(1)
	}
	if *conf.PickConf.DrawThumb {
		err = u.DrawTmpThumbnail(colorStrings.HEX)
		if err != nil {
			fmt.Printf("Error writing thumbnail: %v", err)
			os.Exit(1)
		}
	}

	u.OutputAsLines(colorStrings)
}
