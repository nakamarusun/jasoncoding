package cool

var cfg CoolOptions

type CoolOptions struct {
	WordList  []string
	ImgFormat string
	Colors    []string
	Quality   string
	W int
	H int
}

func New() CoolOptions {
	return CoolOptions{
		W: 480,
		H: 240,
		WordList: []string{
			"about",
			"above",
			"across",
			"act",
			"active",
			"activity",
			"add",
			"afraid",
			"after",
			"again",
		},
		ImgFormat: "jpg",
		Quality:   "30%",

		// Protanopia and deuteranopia color blindness safe palette
		// https://www.nature.com/articles/nmeth.1618
		Colors: []string{
			"#000000",
			"#E69F00",
			"#56B4E9",
			"#009E73",
			"#F0E442",
			"#0072B2",
			"#D55E00",
			"#CC79A7",
		},
	}
}
