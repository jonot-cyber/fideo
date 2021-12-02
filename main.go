package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

// Draw the main image. Will be replaced with json later
func draw() *image.NRGBA {
	project := Project{
		Resolution: Vector2{1920, 1080},
		Main: Group{
			Size: Vector2{1920, 1080},
			Children: []Object{
				Rectangle{
					Position: Vector2{100, 100},
					Size:     Vector2{500, 500},
					Color:    color.NRGBA{R: 111, G: 188, B: 22, A: math.MaxUint8},
				},
			},
		},
	}
	return project.Main.Render()
}

func main() {
	img := draw()

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
