package main

import (
	"fmt"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func image(s string, tab []int) {
	// If file is in current directory. This can also be a URL to an image or gif.
	var filePath = s

	flags := aic_package.DefaultFlags()

	// This part is optional.
	// You can directly pass default flags variable to aic_package.Convert() if you wish.
	// There are more flags, but these are the ones shown for demonstration
	flags.Dimensions = tab
	flags.Colored = true
	//flags.SaveTxtPath = "."
	//flags.SaveImagePath = "."
	flags.CustomMap = " .-=+#@"
	//flags.FontFilePath = "./RobotoMono-Regular.ttf" // If file is in current directory
	//flags.SaveBackgroundColor = [3]int{50, 50, 50}

	// Note: For environments where a terminal isn't available (such as web servers), you MUST
	// specify atleast one of flags.Width, flags.Height or flags.Dimensions

	// Conversion for an image
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
}
