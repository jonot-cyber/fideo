package main

import "image"

type Group struct {
	Size     Vector2  // Size of the group render. Make sure to pass from project
	Children []Object // A list of objects to be rendered onto one layer
}

func (g Group) Render() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, g.Size.X, g.Size.Y))
	for _, child := range g.Children {
		child_image := child.Render()

		for y := 0; y < child.GetRect().Dy(); y++ {
			for x := 0; x < child.GetRect().Dx(); x++ {
				color := child_image.At(x, y)
				img.Set(x+child.GetRect().Min.X, y+child.GetRect().Min.Y, color)
			}
		}
	}
	return img
}
