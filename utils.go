package maze

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/exp/constraints"
)

func WritePNG[T constraints.Integer](grid [][]T, filename string) error {
	img := image.NewRGBA(image.Rect(0, 0, len(grid[0]), len(grid)))
	for y, row := range grid {
		for x, cell := range row {
			var col color.Color
			if cell == 1 {
				col = color.RGBA{0, 0, 255, 255}
			} else {
				col = color.Gray{30}
			}
			img.Set(x, y, col)
		}
	}
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()
	return png.Encode(outFile, img)
}
