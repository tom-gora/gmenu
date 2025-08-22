package gcolors

type RGB struct {
	R uint8
	G uint8
	B uint8
}

type HSL struct {
	H float64
	S float64
	L float64
}

type ColorStrings struct {
	HEX          string
	RGB          string
	RGBA         string
	HSL          string
	HSLA         string
	OKLAB        string
	OKLCH        string
	ClosestNamed string
}

type ShadesStrings struct {
	Dark10  string
	Dark9   string
	Dark8   string
	Dark7   string
	Dark6   string
	Dark5   string
	Dark4   string
	Dark3   string
	Dark2   string
	Dark1   string
	Input   string
	Light1  string
	Light2  string
	Light3  string
	Light4  string
	Light5  string
	Light6  string
	Light7  string
	Light8  string
	Light9  string
	Light10 string
}

type Palette struct {
	Color1 string
	Color2 string
	Color3 string
	Color4 string
	Color5 string
}
