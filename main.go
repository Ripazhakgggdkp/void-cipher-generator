package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/elgopher/pixiq/glfw"
	"github.com/elgopher/pixiq/goimage"
)

func main() {

	glfw.StartMainThreadLoop(func(loop *glfw.MainThreadLoop) {
		// Create OpenGL object.
		gl, err := glfw.NewOpenGL(loop)
		if err != nil {
			log.Panicf("NewOpenGL failed: %v", err)
		}
		// Destroy OpenGL when the function ends
		defer gl.Destroy()

		img := gl.NewImage(300, 300)

		wholeSelection := img.WholeImageSelection()
		drawBackground(&wholeSelection, Gray2)

		lowerHalf := wholeSelection.Selection(0, 90).WithSize(300, 300)
		drawBackground(&lowerHalf, Black)

		announcement := getDrawer(Announcement)
		trailer := getDrawer(Trailer)

		announcement.drawSentence(&wholeSelection, 290, 10, sentenceToCipher("JULKAISUMYOS/ITCHIOSSA/SAIMAA"))
		trailer.drawSentence(&wholeSelection, 290, 100, sentenceToCipher("WHATTHEFUCK/DIDYOUJUST/FUCKINGSAY"))

		newImage := goimage.FromSelection(wholeSelection, goimage.Zoom(5))
		existingIMage := image.NewRGBA(image.Rect(0, 0, wholeSelection.Width(), wholeSelection.Height()))
		goimage.FillWithSelection(existingIMage, wholeSelection)
		goimage.CopyToSelection(newImage, wholeSelection)
		save("output.png", newImage)

	})
}

func save(filePath string, img image.Image) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img)
}
