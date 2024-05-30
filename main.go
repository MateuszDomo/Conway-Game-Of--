package main

import (
	"conway-v2/utils"
	"fmt"
	"time"
)

type Grid struct {
	height            int
	width             int
	window_interal_ms int
	grid_interval_ms  int
}

type Window struct {
	height      int
	width       int
	cells       [][]int
	gridOffsetH int
	gridOffsetW int
}

func main() {

	TerminalClear()

	height := 20
	width := 50
	cells := make([][]int, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]int, width)
	}

	utils.RandomlyPopulateCells(height, width, &cells)

	// Hide cursor
	fmt.Print("\033[?25l")

	for {
		cells_cp := make([][]int, height)
		for i := 0; i < height; i++ {
			cells_cp[i] = make([]int, width)
		}
		for r := 0; r < height; r++ {
			for c := 0; c < width; c++ {
				neighbors := utils.CalcNumNeighbors(r, c, height, width, &cells)
				state := utils.CalcCurrentState(neighbors, cells[r][c])
				if state != cells[r][c] {
					if state == 1 {
						TerminalWrite(r, c, "1")
						//TerminalWrite(r, c, "█")
					} else {
						TerminalWrite(r, c, " ")
					}
				}
				cells_cp[r][c] = state
			}
		}

		cells = cells_cp
		time.Sleep(85 * time.Millisecond)
	}
}
