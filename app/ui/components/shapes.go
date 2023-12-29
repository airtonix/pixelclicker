package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawBox(
	buffer *ebiten.Image,
	x, y, width, height float32,
) {
	vector.StrokeRect(
		buffer,
		x, y,
		width, height,
		2,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
		false,
	)
}
