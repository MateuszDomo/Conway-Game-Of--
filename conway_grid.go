package main

import (
	"conway-v2/terminal_utils"
	"time"
)

type ConwayGrid struct {
	rows             int
	cols             int
	iteration_cycles int
	conway_games     [][]*ConwayGame
}

func NewConwayGrid(rows int, cols int, iteration_cycles int) *ConwayGrid {

	terminal_utils.EnableDebug()

	game_height := 20
	game_width := 50

	conway_games := make([][]*ConwayGame, rows)
	for i := 0; i < rows; i++ {
		conway_games[i] = make([]*ConwayGame, cols)
		for j := 0; j < cols; j++ {
			conway_game := NewConwayGame(game_height, game_width, i, j, terminal_utils.TerminalRandomColor())
			conway_game.RandomlyPopulate()
			conway_games[i][j] = conway_game
		}
	}

	return &ConwayGrid{
		rows,
		cols,
		iteration_cycles,
		conway_games,
	}
}

func (conway_grid *ConwayGrid) Run() {

	for k := 0; k < conway_grid.iteration_cycles; k++ {
		for i := 0; i < conway_grid.rows; i++ {
			for j := 0; j < conway_grid.cols; j++ {
				conway_game := conway_grid.conway_games[i][j]
				conway_game.PlayCycle()
			}
		}
		time.Sleep(85 * time.Millisecond)
	}
}
