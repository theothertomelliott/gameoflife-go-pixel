package gameoflife

import (
	"math/rand"
	"time"
)

// Grid represents a grid on which the Game of Life is "played".
type Grid [][]bool

// New creates an empty grid with the specified width and height.
func New(w, h int) Grid {
	var newGrid Grid
	newGrid = make([][]bool, w)
	for index := range newGrid {
		newGrid[index] = make([]bool, h)
	}
	return Grid(newGrid)
}

// Width returns the width of the grid
func (grid Grid) Width() int {
	return len(grid)
}

// Height returns the height of the grid
func (grid Grid) Height() int {
	if len(grid) == 0 {
		return 0
	}
	// All heights should be equal
	return len(grid[0])
}

// Populate randomly sets live cells in the grid.
func (grid Grid) Populate() {
	sizeX := grid.Width()
	sizeY := grid.Height()
	rand.Seed(time.Now().Unix())
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			grid[i][j] = rand.Intn(2) == 1
		}
	}
}

// TurnCrank applies Conway's crank to the grid and returns the next state as a new grid.
func (grid Grid) TurnCrank() Grid {
	var newGrid Grid
	newGrid = make(Grid, len(grid))
	for index := range newGrid {
		newGrid[index] = make([]bool, len(grid[index]))
	}

	for i := range grid {
		for j := range grid[i] {
			neighbors := grid.countNeighbors(i, j)
			live := grid[i][j]
			switch {
			case neighbors < 2 && live:
				newGrid[i][j] = false
			case neighbors > 3 && live:
				newGrid[i][j] = false
			case neighbors == 3 && !live:
				newGrid[i][j] = true
			default:
				newGrid[i][j] = live
			}
		}
	}

	return newGrid
}

// Walk traverses all cells in the grid, calling the provided visit function for each.
// If the visit function returns an error for any cell, the error will be returned and the traversal will stop.
func (grid Grid) Walk(visit func(x, y int, value bool) error) error {
	sizeX := grid.Width()
	sizeY := grid.Height()
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			err := visit(i, j, grid[int(i)][int(j)])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// countNeighbors returns the number of live neighbors for the given position in the grid.
func (grid Grid) countNeighbors(x, y int) int {
	var neighbors = 0
	if y > 0 {
		if x > 0 && grid[x-1][y-1] { // Bottom left
			neighbors++
		}
		if grid[x][y-1] { // Bottom center
			neighbors++
		}
		if x < len(grid)-1 && grid[x+1][y-1] { // Bottom right
			neighbors++
		}
	}
	if x > 0 && grid[x-1][y] { // Middle left
		neighbors++
	}
	if x < len(grid)-1 && grid[x+1][y] { // Middle right
		neighbors++
	}
	if y < len(grid[x])-1 {
		if x > 0 && grid[x-1][y+1] { // Top left
			neighbors++
		}
		if grid[x][y+1] { // Top center
			neighbors++
		}
		if x < len(grid)-1 && grid[x+1][y+1] { // Top right
			neighbors++
		}
	}
	return neighbors
}
