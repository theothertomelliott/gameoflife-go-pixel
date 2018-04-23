package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	gameoflife "github.com/theothertomelliott/gameoflife-go-pixel"
)

const sizeX, sizeY = 100, 100

var grid gameoflife.Grid

func update(screen *ebiten.Image) error {

	// Clear the screen
	screen.Fill(color.Black)

	screenWidth, _ := screen.Size()
	width, height := screenWidth/sizeX, screenWidth/sizeY
	cell, err := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if err != nil {
		return err
	}
	err = cell.Fill(color.White)
	if err != nil {
		return err
	}

	grid.Walk(func(x, y int, value bool) error {
		if !value {
			return nil
		}
		offset := ebiten.GeoM{}
		offset.Translate(float64(width)*float64(x), float64(height)*float64(y))
		err = screen.DrawImage(cell, &ebiten.DrawImageOptions{
			GeoM: offset,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func main() {
	grid = gameoflife.New(sizeX, sizeY)
	grid.Populate()

	go func() {
		for true {
			grid = grid.TurnCrank()
			time.Sleep(time.Second / 25)
		}
	}()

	ebiten.Run(update, 320, 240, 2, "Hello world!")
}
