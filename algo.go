package main

import (
	"os"
	_"fmt"
	"image"
	"image/gif"
	"image/color"
	_"strings"
)

func BubbleSort (array []int) {
	// gif controls
	anim := gif.GIF{LoopCount: -1}
	rect := image.Rect(0, 0, 3*100+1, 3*100+1)

	// bubble sort.
	for i := range array {
		if i == 0 { continue }
		for x := i; x >= 1 && array[x] < array[x-1]; {
			// selecting the xth bar
			appendToGif(&anim, rect, array, x);

			// doing the swap;;
			array[x], array[x-1] = array[x-1], array[x]
			x -= 1
			
			// drawing the result.
			appendToGif(&anim, rect, array, x);
		}
	}
	gif.EncodeAll(os.Stdout, &anim)
}

func appendToGif(seq *gif.GIF, rect image.Rectangle, array []int, x int) {
	palette := []color.Color{
		color.White,
		color.Black,
		color.RGBA{0xcc, 0x00, 0x88, 0xff},
	}
	img := image.NewPaletted(rect, palette)
	drawArray(array, img, x);
	seq.Delay = append(seq.Delay, 40)
	seq.Image = append(seq.Image, img)
}
