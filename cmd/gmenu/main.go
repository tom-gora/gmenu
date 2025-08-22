package main

import (
	"fmt"
	"gmenu/cli"
	menu "gmenu/internal/gmenu_menu"
	"os"
	"reflect"
)

func main() {
	confPtr, err := cli.ParseArgsToConf()
	if err != nil {
		fmt.Printf("An error parsing the command arguments occured:\n%v\n", err)
		os.Exit(1)
	}

	switch {
	// case no result been passed. act default-construct and output menu entries
	case confPtr.MenuConf != nil &&
		reflect.TypeOf(confPtr.MenuConf).Elem() == reflect.TypeOf(cli.MenuConf{}) &&
		*confPtr.MenuConf.Result == "" &&
		*confPtr.MenuConf.MenuConfJSON != "":
		// ------------------------ conditions end
		entries, err := menu.GetMenuEntriesFromJSON(confPtr.MenuConf.MenuConfJSON)
		if err != nil {
			fmt.Printf("An error reading the menu configuration file '%s':\n%v\n", *confPtr.MenuConf.MenuConfJSON, err)
			os.Exit(1)
		}
		handleMenu(entries, confPtr.MenuConf.IsJoin, confPtr.MenuConf.ReturnValue)

	// case result passed by flag. handle result by using adequate command to execute with result str
	case confPtr.MenuConf != nil &&
		reflect.TypeOf(confPtr.MenuConf).Elem() == reflect.TypeOf(cli.MenuConf{}) &&
		*confPtr.MenuConf.Result != "" &&
		*confPtr.MenuConf.MenuConfJSON != "":
		// ------------------------ conditions end
		entries, err := menu.GetMenuEntriesFromJSON(confPtr.MenuConf.MenuConfJSON)
		if err != nil {
			fmt.Printf("An error reading the menu configuration file '%s':\n%v\n", *confPtr.MenuConf.MenuConfJSON, err)
			os.Exit(1)
		}
		handleMenuResult(*confPtr.MenuConf.Result, entries, confPtr.MenuConf.DefaultExec)

	// case pick color
	case confPtr.PickConf != nil &&
		reflect.TypeOf(confPtr.PickConf).Elem() == reflect.TypeOf(cli.PickConf{}):
		handlePick(confPtr.PickConf)

	case confPtr.ShadesConf != nil:
		handleShades(confPtr.ShadesConf)
	}
}
