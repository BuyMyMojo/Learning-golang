package main

import (
	"fmt"

	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
)

func main() {
	img, err := imgio.Open("pakemon2.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := blur.Gaussian(img, 128.0)

	if err := imgio.Save("pakemon2_edited.png", result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}
