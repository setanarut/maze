package main

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/setanarut/maze/rect"
)

// Example usage
func main() {
	mg := rect.NewMazeGenerator(7, 5, 64, 9)
	walls := mg.GenerateMaze()
	img := image.NewRGBA(mg.Bounds())
	fmt.Println(img.Bounds())
	rect.FillRectangle(img, img.Bounds(), image.Black)
	rect.DrawWallsToImage(walls, img)
	imgio.Save("a.png", img, imgio.PNGEncoder())
}
