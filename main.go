package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const sizeX, sizeY = 100, 100
const screenWidth, screenHeight = float64(1024), float64(768)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of Life",
		Bounds: pixel.R(0, 0, screenWidth, screenHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	grid := newGrid(sizeX, sizeY)
	populateGrid(grid)

	for !win.Closed() {
		grid = turnCrank(grid)

		win.Clear(colornames.Aliceblue)
		drawGrid(win, grid)
		win.Update()
		time.Sleep(time.Second / 25)
	}
}

// newGrid creates an empty grid with the specified width and height.
func newGrid(w, h int) [][]bool {
	var newGrid [][]bool
	newGrid = make([][]bool, w)
	for index := range newGrid {
		newGrid[index] = make([]bool, h)
	}
	return newGrid
}

// turnCrank applies Conway's crank to an existing grid and returns the next state as a new grid.
func turnCrank(grid [][]bool) [][]bool {
	var newGrid [][]bool
	newGrid = make([][]bool, len(grid))
	for index := range newGrid {
		newGrid[index] = make([]bool, len(grid[index]))
	}

	for i := range grid {
		for j := range grid[i] {
			neighbors := countNeighbors(i, j, grid)
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

// countNeighbors returns the number of live neighbors for the given position in the provided grid.
func countNeighbors(x, y int, grid [][]bool) int {
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

// populateGrid randomly sets live cells in the provided grid.
func populateGrid(grid [][]bool) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < sizeX; i++ {
		for j := 0; j < sizeY; j++ {
			grid[i][j] = rand.Intn(2) == 1
		}
	}
}

// drawGrid draws the provided grid to the specified window
func drawGrid(win *pixelgl.Window, grid [][]bool) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)
	screenWidth := win.Bounds().W()
	width, height := screenWidth/sizeX, screenWidth/sizeY
	for i := float64(0); i < sizeX; i++ {
		for j := float64(0); j < sizeY; j++ {
			if grid[int(i)][int(j)] {
				imd.Push(pixel.V(width*i, height*j))
				imd.Push(pixel.V(width*i+width, height*j+height))
				imd.Rectangle(0)
			}
		}
	}
	imd.Draw(win)
}
