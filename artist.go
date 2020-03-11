// package utility to draw bar(s)
package main

import (
	"image"
	"tuts/lissajous/chars"
)

func drawBar(img *image.Paletted, x, y int, breadth, length int, present bool) {
	for v := 0; v < length; v++ {
		for b := 0; b < breadth; b++ {
			switch {
			case present:
				img.SetColorIndex(v + x, y - b, 1)
				continue
			case v == 0 || v == length - 1:
				fallthrough
			case (v != 0 || v != length - 1) && b == 0 || b == breadth - 1:
				img.SetColorIndex(v + x, y - b, 1)
			default:
			}
		}
	}
}

// TODO:: add an object?? for the configuration options.
func drawArray(array []int, img *image.Paletted, present int) {
	// vertical bar (drawBar) config
	barWidth := 20
	barGutter := 10
	lee := 0 // the width between the leftmost side of the first bar to the leftmost side of it's adjacent bar.
	xscale := 2 // the bar draws really small... xscale scales it up by n
	
	// typography (PrintLetters) config
	letterGutter := 2
	letterWidth := 4
	letterHeight := 8
	xPadding := 25
	yPadding := 280

	/* 
		gif structure (TODO::)
																		------------------------------- 0 
																		-		==       						      - 1
																		-		==				==					    - 2
																		-		==	 ==		==				==		- .
																		-		==	 ==		==	 ==		==		- .
		bars representing the digits -> -   ==   ==   ==   ==   ==    - .
		       number representation -> -   5    3    4    2    3     - 280 - 288 (full height of the number)
																		-------------------------------
	*/
		

	// loop through the array.
	for i, value := range array {
		drawBar(
			img, // image written
			lee + xPadding - 5, yPadding - 15, // x, y
			value * xscale, barWidth,
			i == present,
		)

		// find out the first and second value of the number...
		// NB: this function only works with two valued numbers.
		last := value % 10
		numberArray := [...]int{1: last, 0: (value - last) / 10}

		// "draw" the number for the nth bar
		for l, v := range numberArray {
			chars.PrintSingleNumber(
				v, img,
				lee + xPadding + (letterWidth + letterGutter) * l, yPadding, // x, y
				letterHeight, letterWidth, // height, width
			)
		}
		
		// increase the furthermost point for the next iteration
		lee += barGutter + barWidth
	}
}