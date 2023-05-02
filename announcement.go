package main

import (
	"github.com/elgopher/pixiq/image"
)

type AnnouncementDrawer struct {
	Drawer
	Theme
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

func newAnnouncementDrawer() Drawer {
	return &AnnouncementDrawer{
		Theme:  AnnouncementTheme,
		Drawer: newDefaultDrawer(AnnouncementTheme),
	}
}
