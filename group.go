package main

import (
	"image"
	"image/draw"
)

type Group struct {
	Size     Vector2  // Size of the group render. Make sure to pass from project
	Children []Object // A list of objects to be rendered onto one layer
}

func (g Group) Render(size Vector2) image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, size.X, size.Y))
	for _, child := range g.Children {
		childImg := child.Render(size)
		draw.Draw(img, image.Rect(0, 0, size.X, size.Y), &childImg, image.Point{}, draw.Over)
	}
	return *img
}
