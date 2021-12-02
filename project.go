package main

type Vector2 struct {
	X int
	Y int
}

type Project struct {
	Resolution Vector2
	Main       Group // The group for the project to render.
}
