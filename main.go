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
func drawProject() image.NRGBA {
	project := Project{
		Resolution: Vector2{1920, 1080},
		Main: Group{
			Size: Vector2{1920, 1080},
			Children: []Object{
				Rectangle{
					Min: Vector2{0, 0},
					Max: Vector2{500, 500},
					Color: color.NRGBA{
						R: 123,
						G: 126,
						B: 199,
						A: math.MaxUint8,
					},
				},
			},
		},
	}
	return project.Draw()
}

func main() {
	img := drawProject()

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, &img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
