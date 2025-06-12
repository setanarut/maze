package main

import "github.com/setanarut/maze"

func main() {

	cellSize, wallThickness := 32, 3
	m := maze.NewMaze(9, 5, cellSize, wallThickness)

	m.Generate(0, 1)
	maze.WritePNG(m.Grid, "examples/demo/maze1.png")

	m.Generate(0, 3)
	maze.WritePNG(m.Grid, "examples/demo/maze2.png")
}
