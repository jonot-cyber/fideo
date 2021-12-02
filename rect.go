package main

import (
	"image"
	"image/color"
)

type Rectangle struct {
	Position Vector2     // The top-left corner of the rectangle
	Size     Vector2     // The height and the width of the rectangle
	Color    color.NRGBA // Fill color
}

func (r Rectangle) Render() *image.NRGBA {
	img := image.NewNRGBA(r.GetBounds())
	for y := 0; y < r.Size.Y; y++ {
		for x := 0; x < r.Size.X; x++ {
			img.Set(x, y, r.Color)
		}
	}
	return img
}

// Get the bounds of the rectangle
func (r Rectangle) GetBounds() image.Rectangle {
	return image.Rect(
		0,
		0,
		r.Size.X,
		r.Size.Y,
	)
}

func (r Rectangle) GetRect() image.Rectangle {
	return image.Rect(
		r.Position.X,
		r.Position.Y,
		r.Position.X+r.Size.X,
		r.Position.Y+r.Size.Y,
	)
}
