package main

import "image"

type Vector2 struct {
	X int
	Y int
}

type Project struct {
	Resolution Vector2 // How big the output image should be
	Main       Group   // The group for the project to render.
}

// Render the main group and return the full image
func (p Project) Draw() image.NRGBA {
	return p.Main.Render(p.Resolution)
}
