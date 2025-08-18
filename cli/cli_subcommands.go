package cli

import "flag"

// declare subcommands
var (
	pickSubcmd    *flag.FlagSet = flag.NewFlagSet("pick", flag.ExitOnError)
	shadesSubcmd  *flag.FlagSet = flag.NewFlagSet("shades", flag.ExitOnError)
	paletteSubcmd *flag.FlagSet = flag.NewFlagSet("palette", flag.ExitOnError)
	menuSubcmd    *flag.FlagSet = flag.NewFlagSet("menu", flag.ExitOnError)
)
