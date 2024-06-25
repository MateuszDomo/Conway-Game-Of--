package conway_utils

import ()

type ConwayObject interface {
	IsCellAlive(int, int) bool
	PopulateCell(int, int)
}

func (game *ConwayGame) IsCellAlive(row int, col int) bool {
	return game.cells[row][col] == 1
}

func (game *ConwayGame) PopulateCell(row int, col int) {
	game.cells[row][col] = 1
}

func (grid *ConwayGrid) IsCellAlive(row int, col int) bool {
	return grid.states[row][col] == 1
}

func (grid *ConwayGrid) PopulateCell(row int, col int) {
	grid.conway_games[row][col].RandomlyPopulate()
}
