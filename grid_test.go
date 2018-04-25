package gameoflife

import (
	"reflect"
	"testing"
)

func TestSizing(t *testing.T) {
	var tests = []struct {
		name           string
		grid           Grid
		expectedWidth  int
		expectedHeight int
	}{
		{
			name: "empty",
		},
		{
			name: "1x1",
			grid: Grid(
				[][]bool{
					{false},
				},
			),
			expectedWidth:  1,
			expectedHeight: 1,
		},
		{
			name: "3x2",
			grid: Grid(
				[][]bool{
					{false, false, false},
					{false, false, false},
				},
			),
			expectedWidth:  2,
			expectedHeight: 3,
		},
	}
	for _, test := range tests {
		w := test.grid.Width()
		if w != test.expectedWidth {
			t.Errorf("width: expected %d, got %d", test.expectedWidth, w)
		}
		h := test.grid.Height()
		if h != test.expectedHeight {
			t.Errorf("height: expected %d, got %d", test.expectedHeight, h)
		}
	}
}

func TestConwaysCrank(t *testing.T) {
	var tests = []struct {
		name         string
		gridIn       [][]bool
		gridExpected [][]bool
	}{
		{
			name:         "single empty cell",
			gridIn:       [][]bool{{false}},
			gridExpected: [][]bool{{false}},
		},
		// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
		{
			name: "death by underpopulation",
			gridIn: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
			gridExpected: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
		// Any live cell with two or three live neighbours lives on to the next generation.
		// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
		// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		{
			name: "line flip",
			gridIn: [][]bool{
				{false, true, false},
				{false, true, false},
				{false, true, false},
			},
			gridExpected: [][]bool{
				{false, false, false},
				{true, true, true},
				{false, false, false},
			},
		},
		// Any live cell with more than three live neighbours dies, as if by overpopulation.
		{
			name: "overcrowding",
			gridIn: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			},
			gridExpected: [][]bool{
				{true, false, true},
				{false, false, false},
				{true, false, true},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid := Grid(test.gridIn)
			grid = grid.TurnCrank()
			if !reflect.DeepEqual(Grid(test.gridExpected), grid) {
				t.Log("Expected:", test.gridExpected)
				t.Log("Got:", grid)
				t.Errorf("Grid doesn't match")
			}
		})
	}
}
