package utils

import (
	"math/rand"
)

func CalcNumNeighbors(r int, c int, height int, width int, cells_ptr *[][]int) int {
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

func CalcCurrentState(neighbors int, prev_state int) int {
	dead_to_alive := prev_state == 0 && neighbors == 3
	alive_to_alive := prev_state == 1 && (neighbors >= 2 && neighbors <= 3)

	if dead_to_alive || alive_to_alive {
		return 1
	}

	return 0
}

func RandomlyPopulateCells(height int, width int, cells_ptr *[][]int) {
	cells := *cells_ptr
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			random := rand.Intn(10)

			if random <= 3 {
				cells[r][c] = 1
			} else {
				cells[r][c] = 0
			}
		}
	}
}
