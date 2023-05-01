package main

import "github.com/elgopher/pixiq/image"

type AnnouncementDrawer struct {
	Theme
}

func (d *AnnouncementDrawer) drawLetter(image *image.Selection, x int, y int, grid [3][3]State) {
	selection := image.Selection(x, y)
	kind := map[State]Pattern{On: d.BoxOnPattern, Off: d.BoxOffPattern, Blank: d.BoxBlankPattern}
	for i, row := range grid {
		for j, column := range row {
			drawBox(&selection, j*5, i*5, kind[column])
		}
	}
}

func (d *AnnouncementDrawer) drawSentence(image *image.Selection, x int, y int, cipher [][3][3]State) {
	line := 0
	caret := 0
	selection := image.Selection(x, y)
	for _, letter := range cipher {
		if letter[0][0] == Blank {
			line++
			caret = 0
			continue
		}

		d.drawLetter(&selection, -caret*16-14, 16*line, letter)
		caret++
	}
}
