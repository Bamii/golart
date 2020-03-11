// golang package that prints numbers and strings on a gif.
package chars

import (
	_"fmt"
	"image"
)

func PrintSingleNumber(letter int, img *image.Paletted, x, y int, height, width int) {
	switch letter {
	case 0:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == 0 || v == height - 1 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if v != 0 || v != height - 1 {
					if b == 0 || b == width - 1 {
						img.SetColorIndex(b + x, y - v, 1)
					}
				}
			}
		}

	case 1:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 && b >= (width / 2) - 3 && b <= width / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if v == 0 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width / 2 && (v != 0 || v != height - 1) {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}
	
	case 2:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 || v == 0 || v == height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 && (v >= height / 2) {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == 0 && (v <= height / 2) {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}
	
	case 3:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 || v == 0 || v == height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}

	case 4:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height / 4 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == 0 && v >= (height / 4) {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width / 2 && v <= height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}

	case 5:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 || v == 0 || v == height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == 0 && (v >= height / 2) {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 && (v <= height / 2) {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}


	case 6:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 || v == 0 || v == height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == 0 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 && v <= height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}

	case 7:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}

	case 8:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 || v == 0 || v == height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == 0 {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}

	case 9:
		for v := 0; v < height; v++ {
			for b := 0; b < width; b++ {
				if v == height - 1 || v == 0 || v == height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == width - 1 {
					img.SetColorIndex(b + x, y - v, 1)
				} else if b == 0 && v >= height / 2 {
					img.SetColorIndex(b + x, y - v, 1)
				}
			}
		}
	}
}

// TODO:: BELOW THIS LINE
func PrintNumbers(array[] int) {}

func PrintLetters() {

}

/* 

A - 
[0,0,0,1,1,0,0,0]
[0,0,1,1,1,1,0,0]
[0,1,1,0,0,1,1,0]
[0,1,0,0,0,0,1,0]
[1,1,1,1,1,1,1,1]
[1,1,1,1,1,1,1,1]
[1,1,0,0,0,0,1,1]
[1,1,0,0,0,0,1,1]

B - 
[1,1,1,1,1,1,1,0]
[1,1,0,0,0,1,1,1]
[1,1,0,0,0,1,1,1]
[1,1,1,1,1,1,1,0]
[1,1,1,1,1,1,1,0]
[1,1,0,0,0,0,1,1]
[1,1,0,0,0,0,1,1]
[1,1,1,1,1,1,1,1]

C - 
[0,1,1,1,1,1,1,0]
[1,1,1,1,1,1,1,1]
[1,1,0,0,0,0,0,0]
[1,1,0,0,0,0,0,0]
[1,1,0,0,0,0,0,0]
[1,1,0,0,0,0,0,0]
[1,1,1,1,1,1,1,1]
[0,1,1,1,1,1,1,0]
*/
