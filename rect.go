package main

import (
	"image"
	"image/color"
)

type Rectangle struct {
	Min   Vector2     // The location of the top left corner of the rectangle
	Max   Vector2     // The location of the bottom right corner of the rectangle
	Color color.NRGBA // The fill color of the rectangle
}

// A simple render function. Just fills a rectangle with a color
func (r Rectangle) Render(size Vector2) image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, size.X, size.Y))
	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			img.Set(x, y, r.Color)
		}
	}
	return *img
}
