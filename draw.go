package main

import (
	"github.com/elgopher/pixiq/image"
)

type Style int

const (
	Trailer      Style = iota
	Announcement Style = iota
)

type Drawer interface {
	drawLetter(*image.Selection, int, int, [3][3]State)
	drawSentence(*image.Selection, int, int, [][3][3]State)
}

func drawBox(image *image.Selection, x int, y int, pattern [][]image.Color) {
	selection := image.Selection(x, y)
	for y, colors := range pattern {
		for x, color := range colors {
			if color != Empty {
				selection.SetColor(x, y, color)
			}
		}
	}
}

func drawBackground(image *image.Selection, color image.Color) {
	for i := 0; i < image.Width(); i++ {
		for j := 0; j < image.Height(); j++ {
			image.SetColor(i, j, color)
		}
	}
}

func getDrawer(theme Style) Drawer {
	if theme == Trailer {
		return newTrailerDrawer()
	} else {
		return newAnnouncementDrawer()
	}
}
