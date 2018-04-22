package gameoflife

import (
	"math/rand"
	"time"
)

type GameOfLife [][]bool

// New creates an empty grid with the specified width and height.
func New(w, h int) GameOfLife {
	var newGrid GameOfLife
	newGrid = make([][]bool, w)
	for index := range newGrid {
		newGrid[index] = make([]bool, h)
	}
	return GameOfLife(newGrid)
}

// TurnCrank applies Conway's crank to the grid and returns the next state as a new grid.
func (grid GameOfLife) TurnCrank() GameOfLife {
	var newGrid GameOfLife
	newGrid = make(GameOfLife, len(grid))
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

// countNeighbors returns the number of live neighbors for the given position in the grid.
func (grid GameOfLife) countNeighbors(x, y int) int {
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

// PopulateGrid randomly sets live cells in the grid.
func (grid GameOfLife) Populate(sizeX, sizeY int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			grid[i][j] = rand.Intn(2) == 1
		}
	}
}
