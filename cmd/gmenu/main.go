package main

import (
	"fmt"
	"gmenu/cli"
	menu "gmenu/internal/gmenu_menu"
	"os"
)

func main() {
	conf, err := cli.ParseArgs()
	if err != nil {
		fmt.Printf("An error parsing the command arguments occured:\n%v\n", err)
		os.Exit(1)
	}

	switch {
	case conf.MenuConf != nil && conf.MenuConf.Result == nil:
		entries, err := menu.ReadJSONFile(conf.MenuConf.MenuConfPath)
		if err != nil {
			fmt.Printf("An error reading the menu configuration file '%s':\n%v\n", *conf.MenuConf.MenuConfPath, err)
			os.Exit(1)
		}
		handleMenu(entries, conf.MenuConf.IsJoin)

	case conf.MenuConf != nil && conf.MenuConf.Result != nil:
		entries, err := menu.ReadJSONFile(conf.MenuConf.MenuConfPath)
		if err != nil {
			fmt.Printf("An error reading the menu configuration file '%s':\n%v\n", *conf.MenuConf.MenuConfPath, err)
			os.Exit(1)
		}
		handleMenuResult(*conf.MenuConf.Result, entries, conf.MenuConf.DefaultExec)
	case conf.PickConf != nil:

		handlePick(conf.PickConf)
	}
}
