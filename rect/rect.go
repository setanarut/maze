// rect is a subpackage for generating labyrinth walls of type []image.Rectangle
package rect

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

// Direction represents the four cardinal directions
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// Cell represents a single cell in the maze grid
type Cell struct {
	X, Y    int
	Visited bool
	Walls   [4]bool // North, East, South, West
}

// MazeGenerator holds the maze state
type MazeGenerator struct {
	WallThickness int
	Width, Height int
	CellSize      int
	Grid          [][]Cell
	rng           *rand.Rand
}

// NewMazeGenerator creates a new maze generator
func NewMazeGenerator(width, height, cellSize, wallThickness int) *MazeGenerator {
	mg := &MazeGenerator{
		WallThickness: wallThickness,
		Width:         width,
		Height:        height,
		CellSize:      cellSize,
		Grid:          make([][]Cell, height),
		rng:           rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// Initialize grid with all walls present
	for y := range height {
		mg.Grid[y] = make([]Cell, width)
		for x := range width {
			mg.Grid[y][x] = Cell{
				X:       x,
				Y:       y,
				Visited: false,
				Walls:   [4]bool{true, true, true, true}, // All walls initially present
			}
		}
	}

	return mg
}

// GenerateMaze creates a maze using depth-first search algorithm
func (mg *MazeGenerator) GenerateMaze() []image.Rectangle {
	// Start from top-left corner
	stack := []*Cell{&mg.Grid[0][0]}
	mg.Grid[0][0].Visited = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		// Get unvisited neighbors
		neighbors := mg.getUnvisitedNeighbors(current)

		if len(neighbors) > 0 {
			// Choose random neighbor
			next := neighbors[mg.rng.Intn(len(neighbors))]

			// Remove wall between current and next
			mg.removeWall(current, next)

			// Mark next as visited and add to stack
			next.Visited = true
			stack = append(stack, next)
		} else {
			// Backtrack - pop from stack
			stack = stack[:len(stack)-1]
		}
	}

	// Convert remaining walls to rectangles
	return mg.getWallRectangles()
}

// getUnvisitedNeighbors returns all unvisited neighboring cells
func (mg *MazeGenerator) getUnvisitedNeighbors(cell *Cell) []*Cell {
	var neighbors []*Cell
	x, y := cell.X, cell.Y

	// Check each direction
	directions := []struct {
		dx, dy int
		dir    Direction
	}{
		{0, -1, North}, // North
		{1, 0, East},   // East
		{0, 1, South},  // South
		{-1, 0, West},  // West
	}

	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < mg.Width && ny >= 0 && ny < mg.Height {
			if !mg.Grid[ny][nx].Visited {
				neighbors = append(neighbors, &mg.Grid[ny][nx])
			}
		}
	}

	return neighbors
}

// Size returns the size of the maze in pixels.
//
// The size is calculated as the number of cells multiplied by the cell size,
// plus the wall thickness to account for the walls around the maze.
func (mg *MazeGenerator) Size() image.Point {
	return image.Point{
		mg.Width*mg.CellSize + mg.WallThickness,
		mg.Height*mg.CellSize + mg.WallThickness,
	}
}

// removeWall removes the wall between two adjacent cells
func (mg *MazeGenerator) removeWall(current, next *Cell) {
	dx := next.X - current.X
	dy := next.Y - current.Y

	if dx == 1 { // Moving East
		current.Walls[East] = false
		next.Walls[West] = false
	} else if dx == -1 { // Moving West
		current.Walls[West] = false
		next.Walls[East] = false
	} else if dy == 1 { // Moving South
		current.Walls[South] = false
		next.Walls[North] = false
	} else if dy == -1 { // Moving North
		current.Walls[North] = false
		next.Walls[South] = false
	}
}

// getWallRectangles converts remaining walls to image.Rectangle slices
func (mg *MazeGenerator) getWallRectangles() []image.Rectangle {
	var walls []image.Rectangle
	for y := range mg.Height {
		for x := range mg.Width {
			cell := &mg.Grid[y][x]
			cellX := x * mg.CellSize
			cellY := y * mg.CellSize

			// North wall - extends full width including corners
			if cell.Walls[North] {
				walls = append(walls, image.Rect(
					cellX,
					cellY,
					cellX+mg.CellSize+mg.WallThickness,
					cellY+mg.WallThickness,
				))
			}

			// East wall - extends full height including corners
			if cell.Walls[East] {
				walls = append(walls, image.Rect(
					cellX+mg.CellSize,
					cellY,
					cellX+mg.CellSize+mg.WallThickness,
					cellY+mg.CellSize+mg.WallThickness,
				))
			}

			// South wall - extends full width including corners
			if cell.Walls[South] {
				walls = append(walls, image.Rect(
					cellX,
					cellY+mg.CellSize,
					cellX+mg.CellSize+mg.WallThickness,
					cellY+mg.CellSize+mg.WallThickness,
				))
			}

			// West wall - extends full height including corners
			if cell.Walls[West] {
				walls = append(walls, image.Rect(
					cellX,
					cellY,
					cellX+mg.WallThickness,
					cellY+mg.CellSize+mg.WallThickness,
				))
			}
		}
	}

	return walls
}

// DrawWallsToImage draws the wall rectangles onto an RGBA image
func DrawWallsToImage(walls []image.Rectangle, img *image.RGBA) {
	bounds := img.Bounds()

	for _, wall := range walls {
		// Ensure the wall rectangle is within image bounds
		clipped := wall.Intersect(bounds)
		if clipped.Empty() {
			continue
		}

		FillRectangle(img, wall, color.White)
	}
}

// FillRectangle is a utility function to fill a single rectangle on an RGBA image
func FillRectangle(img *image.RGBA, rect image.Rectangle, fillColor color.Color) {
	bounds := img.Bounds()
	clipped := rect.Intersect(bounds)

	if clipped.Empty() {
		return
	}

	for y := clipped.Min.Y; y < clipped.Max.Y; y++ {
		for x := clipped.Min.X; x < clipped.Max.X; x++ {
			img.Set(x, y, fillColor)
		}
	}
}
