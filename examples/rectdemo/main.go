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
	img := image.NewRGBA(image.Rectangle{Max: mg.Size()})
	fmt.Println(img.Bounds())
	rect.FillRectangle(img, img.Bounds(), image.Black)
	rect.DrawWallsToImage(walls, img)
	imgio.Save("rect.png", img, imgio.PNGEncoder())
}
