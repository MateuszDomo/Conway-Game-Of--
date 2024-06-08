package main

type ConwayGrid struct {
	height        int
	width         int
	window_cycles int
	conway_games  [][]ConwayGame
}

//func NewConwayGrid(grid_height int, grid_width int, game_height int, game_width int) *ConwayGrid {
//	conway_games := make([][]ConwayGame, grid_height)
//	conway_games[0] = nil
//	//return
//}
