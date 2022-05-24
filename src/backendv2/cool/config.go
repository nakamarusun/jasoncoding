package cool

var cfg CoolOptions

type CoolOptions struct {
	WordList  []string
	ImgFormat string
	Colors    []string
}

func New() CoolOptions {
	return CoolOptions{
		ImgFormat: "jpg",

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
