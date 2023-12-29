package scanning

import (
	"image"

	"github.com/go-vgo/robotgo"
	"github.com/hajimehoshi/ebiten/v2"
)

/**
 *
 */
func GetPixels(
	x, y, width, height int,
) [][]string {
	var output [][]string

	for i := 0; i < height; i++ {
		var row []string
		for j := 0; j < width; j++ {

			pixel := robotgo.GetPixelColor(x+j, y+i)
			row = append(row, pixel)

		}
		output = append(output, row)
	}

	return output
}

func CaptureScreen() image.Image {
	x, y := ebiten.WindowPosition()
	screenWidth, screenHeight := ebiten.WindowSize()
	captured := robotgo.CaptureImg(x, y, screenWidth, screenHeight)
	return captured
}
