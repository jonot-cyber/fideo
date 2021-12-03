package main

import "image"

// An interface for structs that can render something
type Renderer interface {
	Render(size Vector2) image.NRGBA // size is the size of what to render onto. Outputs the image
}
