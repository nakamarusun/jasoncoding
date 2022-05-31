package cool

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"jasoncoding.com/backendv2/config"
)

var cfg CoolOptions

type CoolOptions struct {
	WordList  []string
	ImgFormat string
	Colors    []string
	Quality   string
	W         int
	H         int
	FontList  []string
	ColRange  int // Collision range for the words
}

func New() CoolOptions {
	fontPath := config.Cfg.GetString("FONT_PATH")
	fontList := make([]string, 0)
	if fontPath != "" {
		// Gets the font list
		files, err := ioutil.ReadDir(fontPath)
		if err != nil {
			fmt.Printf("Error loading fonts %s", err)
		}

		for _, file := range files {
			filename := file.Name()
			if !file.IsDir() && (strings.HasSuffix(filename, ".ttf") ||
				strings.HasSuffix(filename, ".otf") ||
				strings.HasSuffix(filename, ".woff") ||
				strings.HasSuffix(filename, ".woff2")) {
				fontList = append(fontList, filepath.Join(fontPath, filename))
			}
		}
	}

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
		Quality:   "60%",

		// Protanopia and deuteranopia color blindness safe palette
		// https://www.nature.com/articles/nmeth.1618
		Colors: []string{
			"#000000",
			"#E69F00",
			"#56B4E9",
			"#009E73",
			// "#F0E442",
			"#0072B2",
			"#D55E00",
			"#CC79A7",
		},

		FontList: fontList,
		ColRange: 70,
	}
}
