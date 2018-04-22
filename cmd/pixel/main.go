package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	gameoflife "github.com/theothertomelliott/gameoflife-go-pixel"
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

	grid := gameoflife.New(sizeX, sizeY)
	grid.Populate(sizeX, sizeY)

	for !win.Closed() {
		grid = grid.TurnCrank()

		win.Clear(colornames.Aliceblue)
		drawGrid(win, grid)
		win.Update()
		time.Sleep(time.Second / 25)
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
