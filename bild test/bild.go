package main

import (
	"fmt"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)


	go func() {
		blurImg("./pakemon1.png", "./pakemon1-blur.png")
		wg.Done()
	}()

	go func() {
		blurImg("./pakemon2.png", "./pakemon2-blur.png")
		wg.Done()
	}()

	go func() {
		blurImg("./pakemon3.png", "./pakemon3-blur.png")
		wg.Done()
	}()

	go func() {
		blurImg("./pakemon4.png", "./pakemon4-blur.png")
		wg.Done()
	}()

	go func() {
		blurImg("./pakemon5.png", "./pakemon5-blur.png")
		wg.Done()
	}()

	wg.Wait()
}

// blur images
func blurImg(source string, target string) {
	img, err := imgio.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := blur.Gaussian(img, 128.0)

	if err := imgio.Save(target, result, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

}
