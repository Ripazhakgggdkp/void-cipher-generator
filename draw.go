package main

import (
	"github.com/elgopher/pixiq/image"
)

var (
	White = image.RGB(255, 255, 255)
	Gray  = image.RGB(192, 192, 192)
	Gray2 = image.RGB(128, 128, 128)
	Black = image.RGB(0, 0, 0)
)

func drawBox(image *image.Selection, x int, y int, on bool) {
	if on {
		drawBoxOn(image, x, y)
	} else {
		drawBoxOff(image, x, y)
	}
}

func drawBoxOn(image *image.Selection, x int, y int) {
	selection := image.Selection(x, y)
	selection.SetColor(0, 0, White)
	selection.SetColor(1, 0, White)
	selection.SetColor(2, 0, White)
	selection.SetColor(3, 0, Gray)
	selection.SetColor(0, 1, White)
	selection.SetColor(1, 1, Gray)
	selection.SetColor(2, 1, Gray)
	selection.SetColor(3, 1, Gray2)
	selection.SetColor(0, 2, White)
	selection.SetColor(1, 2, Gray)
	selection.SetColor(2, 2, Gray)
	selection.SetColor(3, 2, Gray2)
	selection.SetColor(0, 3, Gray)
	selection.SetColor(1, 3, Gray2)
	selection.SetColor(2, 3, Gray2)
	selection.SetColor(3, 3, Gray2)
}

func drawBoxOff(image *image.Selection, x int, y int) {
	selection := image.Selection(x, y)
	selection.SetColor(0, 0, Black)
	selection.SetColor(1, 0, Black)
	selection.SetColor(2, 0, Black)
	selection.SetColor(3, 0, Black)
	selection.SetColor(0, 1, Black)
	selection.SetColor(1, 1, Black)
	selection.SetColor(2, 1, Black)
	selection.SetColor(3, 1, Gray2)
	selection.SetColor(0, 2, Black)
	selection.SetColor(1, 2, Black)
	selection.SetColor(2, 2, Black)
	selection.SetColor(3, 2, Gray2)
	selection.SetColor(0, 3, Black)
	selection.SetColor(1, 3, Gray2)
	selection.SetColor(2, 3, Gray2)
	selection.SetColor(3, 3, Black)
}

func drawBoxBlack(image *image.Selection, x int, y int) {
	selection := image.Selection(x, y)
	selection.SetColor(0, 0, Black)
	selection.SetColor(1, 0, Black)
	selection.SetColor(2, 0, Black)
	selection.SetColor(3, 0, Black)
	selection.SetColor(0, 1, Black)
	selection.SetColor(1, 1, Black)
	selection.SetColor(2, 1, Black)
	selection.SetColor(3, 1, Black)
	selection.SetColor(0, 2, Black)
	selection.SetColor(1, 2, Black)
	selection.SetColor(2, 2, Black)
	selection.SetColor(3, 2, Black)
	selection.SetColor(0, 3, Black)
	selection.SetColor(1, 3, Black)
	selection.SetColor(2, 3, Black)
	selection.SetColor(3, 3, Black)
}

func drawVerticalPadding(image *image.Selection, x int, y int) {
	selection := image.Selection(x, y)
	selection.SetColor(0, 0, Black)
	selection.SetColor(0, 1, Black)
	selection.SetColor(0, 2, Black)
	selection.SetColor(0, 3, Black)
}

func drawHorizontalPadding(image *image.Selection, x int, y int) {
	selection := image.Selection(x, y)
	selection.SetColor(0, 0, Black)
	selection.SetColor(1, 0, Black)
	selection.SetColor(2, 0, Black)
	selection.SetColor(3, 0, Black)
	selection.SetColor(4, 0, Black)
	selection.SetColor(5, 0, Black)
	selection.SetColor(6, 0, Black)
	selection.SetColor(7, 0, Black)
	selection.SetColor(8, 0, Black)
	selection.SetColor(9, 0, Black)
	selection.SetColor(10, 0, Black)
	selection.SetColor(11, 0, Black)
	selection.SetColor(12, 0, Black)
	selection.SetColor(13, 0, Black)

}
func drawLetter(image *image.Selection, x int, y int, grid [8]bool) {
	selection := image.Selection(x, y)
	for i, p := range grid {
		if i < 2 {
			drawBox(&selection, i*5, 0, p)
			drawVerticalPadding(&selection, i*5+4, 0)
		}
		if i == 2 {
			drawBox(&selection, i*5, 0, p)
			drawHorizontalPadding(&selection, 0, 4)
		}
		if i == 3 {
			drawBox(&selection, 0, 5, p)
			drawVerticalPadding(&selection, 4, 5)
			drawBoxBlack(&selection, 5, 5)
			drawVerticalPadding(&selection, 9, 5)
		}
		if i == 4 {
			drawBox(&selection, 10, 5, p)
			drawHorizontalPadding(&selection, 0, 9)
		}
		if i > 4 && i < 7 {
			drawBox(&selection, (i-5)*5, 10, p)
			drawVerticalPadding(&selection, (i-5)*5+4, 10)
		}
		if i == 7 {
			drawBox(&selection, (i-5)*5, 10, p)
		}
	}
}

func drawSentence(image *image.Selection, x int, y int, cipher [][8]bool) {
	line := 0
	caret := 0
	selection := image.Selection(x, y)
	for _, letter := range cipher {
		if letter[0] == false {
			line++
			caret = 0
			continue
		}

		drawLetter(&selection, -caret*15-14, 15*line, letter)
		caret++
	}
}
