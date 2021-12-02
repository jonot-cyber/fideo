package main

import "image"

type Object interface {
	Render() *image.NRGBA
	GetRect() image.Rectangle
}
