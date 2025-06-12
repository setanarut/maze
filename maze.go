package maze

import (
	"math/rand/v2"
)

type Maze struct {
	Grid          [][]int    // 0: path, 1: wall
	Visited       [][]bool   // visited cells for DFS
	Rnd           *rand.Rand // random number generator
	CellSize      int        // path width in pixels
	WallThickness int        // wall thickness in pixels
	Cols          int        // number of maze cells (width)
	Rows          int        // number of maze cells (height)
}

func NewMaze(w, h, cellSize, wallThickness int) *Maze {
	R := h*cellSize + (h+1)*wallThickness
	C := w*cellSize + (w+1)*wallThickness

	// Pre-allocate matrix
	mat := make([][]int, R)
	for i := range mat {
		mat[i] = make([]int, C)
	}

	// Pre-allocate visited matrix
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	return &Maze{
		Grid:          mat,
		Visited:       visited,
		CellSize:      cellSize,
		WallThickness: wallThickness,
		Cols:          w,
		Rows:          h,
	}

}

func (m *Maze) Generate(seed1 uint64, seed2 uint64) {
	m.Rnd = rand.New(rand.NewPCG(seed1, seed2))

	// Reset matrix to all walls
	for i := range m.Grid {
		for j := range m.Grid[i] {
			m.Grid[i][j] = 1
		}
	}

	// Reset visited matrix
	for i := range m.Visited {
		for j := range m.Visited[i] {
			m.Visited[i][j] = false
		}
	}

	// Start DFS from (0,0)
	m.dfs(0, 0)
}

func (m *Maze) dfs(r, c int) {
	m.Visited[r][c] = true

	// Fill cell area with path (0)
	startY := m.WallThickness + r*(m.CellSize+m.WallThickness)
	startX := m.WallThickness + c*(m.CellSize+m.WallThickness)
	for y := range m.CellSize {
		for x := range m.CellSize {
			wy := startY + y
			wx := startX + x
			if wy >= 0 && wy < len(m.Grid) && wx >= 0 && wx < len(m.Grid[0]) {
				m.Grid[wy][wx] = 0
			}
		}
	}

	dirs := m.Rnd.Perm(4)
	for _, dir := range dirs {
		var nr, nc int
		switch dir {
		case 0: // up
			nr, nc = r-1, c
			if nr >= 0 && !m.Visited[nr][nc] {
				// open wall above
				for x := range m.CellSize {
					for y := range m.WallThickness {
						wy := startY - m.WallThickness + y
						wx := startX + x
						if wy >= 0 && wy < len(m.Grid) && wx >= 0 && wx < len(m.Grid[0]) {
							m.Grid[wy][wx] = 0
						}
					}
				}
				m.dfs(nr, nc)
			}
		case 1: // left
			nr, nc = r, c-1
			if nc >= 0 && !m.Visited[nr][nc] {
				// open wall to the left
				for y := range m.CellSize {
					for x := range m.WallThickness {
						wy := startY + y
						wx := startX - m.WallThickness + x
						if wy >= 0 && wy < len(m.Grid) && wx >= 0 && wx < len(m.Grid[0]) {
							m.Grid[wy][wx] = 0
						}
					}
				}
				m.dfs(nr, nc)
			}
		case 2: // down
			nr, nc = r+1, c
			if nr < m.Rows && !m.Visited[nr][nc] {
				// open wall below
				for x := range m.CellSize {
					for y := range m.WallThickness {
						wy := startY + m.CellSize + y
						wx := startX + x
						if wy >= 0 && wy < len(m.Grid) && wx >= 0 && wx < len(m.Grid[0]) {
							m.Grid[wy][wx] = 0
						}
					}
				}
				m.dfs(nr, nc)
			}
		case 3: // right
			nr, nc = r, c+1
			if nc < m.Cols && !m.Visited[nr][nc] {
				// open wall to the right
				for y := range m.CellSize {
					for x := range m.WallThickness {
						wy := startY + y
						wx := startX + m.CellSize + x
						if wy >= 0 && wy < len(m.Grid) && wx >= 0 && wx < len(m.Grid[0]) {
							m.Grid[wy][wx] = 0
						}
					}
				}
				m.dfs(nr, nc)
			}
		}
	}
}
