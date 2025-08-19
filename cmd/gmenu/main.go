package main

import (
	"fmt"
	"gmenu/cli"
	"os"
)

func main() {
	conf, err := cli.ParseArgs()
	if err != nil {
		fmt.Printf("An error parsing the command arguments occured:\n%v\n", err)
		os.Exit(1)
	}

	switch {
	case conf.MenuConf != nil:
		handleMenu(conf.MenuConf)

	case conf.PickConf != nil:
		handlePick(conf.PickConf)
	}
}
