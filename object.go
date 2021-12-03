package main

import "image"

type Object interface {
	Render(size Vector2) image.NRGBA
}
