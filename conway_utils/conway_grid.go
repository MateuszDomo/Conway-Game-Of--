package conway_utils

import (
	"conway-v2/terminal_utils"
	"time"
)

type ConwayGrid struct {
	rows             int
	cols             int
	iteration_cycles int
	conway_games     [][]*ConwayGame
	states           [][]int // Alive/dead
}

func NewConwayGrid(rows int, cols int, iteration_cycles int) *ConwayGrid {

	terminal_utils.EnableDebug()

	game_height := 10
	game_width := 20

	states := make([][]int, rows)
	for i := 0; i < rows; i++ {
		states[i] = make([]int, cols)
	}

	conway_games := make([][]*ConwayGame, rows)
	for r := 0; r < rows; r++ {
		conway_games[r] = make([]*ConwayGame, cols)
		for c := 0; c < cols; c++ {
			states[r][c] = 1
			conway_game := NewConwayGame(game_height, game_width, r, c, terminal_utils.TerminalRandomColor())
			conway_game.RandomlyPopulate()
			conway_games[r][c] = conway_game
		}
	}

	return &ConwayGrid{
		rows,
		cols,
		iteration_cycles,
		conway_games,
		states,
	}
}

func (grid *ConwayGrid) Run() {
	for {
		grid.PlayCycle()
	}
}

func (grid *ConwayGrid) PlayCycle() {
	for k := 0; k < grid.iteration_cycles; k++ {
		for r := 0; r < grid.rows; r++ {
			for c := 0; c < grid.cols; c++ {
				if grid.states[r][c] == 0 {
					continue
				}

				conway_game := grid.conway_games[r][c]
				conway_game.PlayCycle()
			}
		}
		time.Sleep(85 * time.Millisecond)
	}

	states_cp := make([][]int, grid.rows)
	for r := 0; r < grid.rows; r++ {
		states_cp[r] = make([]int, grid.cols)
	}

	for r := 0; r < grid.rows; r++ {
		for c := 0; c < grid.cols; c++ {
			neighbors := CalcNumNeighbors(r, c, grid.rows, grid.cols, grid)
			state := CalcCurrentState(neighbors, grid.states[r][c])

			if state != grid.states[r][c] {
				if state == 1 {
					grid.conway_games[r][c].RandomlyPopulate()
					grid.conway_games[r][c].WriteGameTerminal(false)
				} else {
					grid.conway_games[r][c].WriteGameTerminal(true)
				}
			}

			states_cp[r][c] = state
		}
	}

	grid.states = states_cp
}
