package cli

// Conf - complete configuration struct
type Conf struct {
	PickConf    *PickConf
	ShadesConf  *ShadesConf
	PaletteConf *PaletteConf
	MenuConf    *MenuConf
}

// MenuConf - config struct for menu subcommand
type MenuConf struct {
	MenuConfPath *string
	Result       *string
	DefaultExec  *string
	IsJoin       *bool
}

// PickConf - config struct for pick subcommand
type PickConf struct {
	Picker    *PickerCommand
	DrawThumb *bool
}

// ShadesConf - config struct for shades subcommand
type ShadesConf struct {
	Picker  *PickerCommand
	ClipMan *ClipManCommand
}

// PaletteConf - config struct for palette subcommand
type PaletteConf struct {
	Picker  *PickerCommand
	ClipMan *ClipManCommand
}

// PickerCommand - color picker command
type PickerCommand struct {
	Cmd  string
	Args []string
}

// ClipManCommand - clipboard manager command
type ClipManCommand struct {
	Cmd  string
	Args []string
}

type MenuEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Icon  string `json:"icon"`
}
