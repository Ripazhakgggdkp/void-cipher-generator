package main

import "github.com/elgopher/pixiq/image"

var (
	White = image.RGB(255, 255, 255)
	Gray  = image.RGB(192, 192, 192)
	Gray2 = image.RGB(128, 128, 128)
	Black = image.RGB(0, 0, 0)
	Empty = image.RGBA(0, 0, 0, 0)
)

type Pattern [][]image.Color

type Theme struct {
	BoxOnPattern    Pattern
	BoxOffPattern   Pattern
	BoxBlankPattern Pattern
}

var TrailerBoxOn = [][]image.Color{
	{White, White, White, Gray},
	{White, Gray, Gray, Gray2},
	{White, Gray, Gray, Gray2},
	{Gray, Gray2, Gray2, Gray2},
}

var TrailerBoxOff = [][]image.Color{
	{Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Gray2},
	{Empty, Empty, Empty, Gray2},
	{Empty, Gray2, Gray2, Empty},
}

var BoxEmpty = [][]image.Color{
	{Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Empty},
}

var AnnouncementBoxOn = [][]image.Color{
	{Empty, Black, Black, Empty},
	{Black, Black, Black, Black},
	{Black, Black, Black, Black},
	{Empty, Black, Black, Empty},
}

var TrailerTheme = Theme{BoxOnPattern: TrailerBoxOn, BoxOffPattern: TrailerBoxOff, BoxBlankPattern: BoxEmpty}
var AnnouncementTheme = Theme{BoxOnPattern: AnnouncementBoxOn, BoxOffPattern: BoxEmpty, BoxBlankPattern: BoxEmpty}
