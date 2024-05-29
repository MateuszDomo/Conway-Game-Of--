package main

import (
	"fmt"
	"time"
)

type Grid struct {
	height int
	width  int
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

	cells[4][2] = 1
	cells[4][3] = 1
	cells[4][4] = 1
	cells[3][4] = 1
	cells[2][3] = 1

	// Hide cursor
	fmt.Print("\033[?25l")

	for {
		cells_cp := make([][]int, height)
		for i := 0; i < height; i++ {
			cells_cp[i] = make([]int, width)
		}
		for r := 0; r < height; r++ {
			for c := 0; c < width; c++ {
				neighbors := calc_neighbors(r, c, height, width, &cells)
				state := calc_state(neighbors, cells[r][c])
				if state != cells[r][c] {
					if state == 1 {
						TerminalWrite(r, c, "0")
					} else {
						TerminalWrite(r, c, " ")
					}
				}
				cells_cp[r][c] = state
			}
		}

		cells = cells_cp
		time.Sleep(100 * time.Millisecond)
	}
}

func calc_neighbors(r int, c int, height int, width int, cells_ptr *[][]int) int {
	cells := *cells_ptr
	neighbors := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			row := ((r+i)%height + height) % height
			col := ((c+j)%width + width) % width
			if cells[row][col] == 1 {
				neighbors++
			}
		}
	}
	return neighbors
}

func calc_state(neighbors int, prev_state int) int {
	dead_to_alive := prev_state == 0 && neighbors == 3
	alive_to_alive := prev_state == 1 && (neighbors >= 2 && neighbors <= 3)

	if dead_to_alive || alive_to_alive {
		return 1
	}

	return 0
}
