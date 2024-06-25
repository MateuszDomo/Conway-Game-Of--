package main

import (
	"conway-v2/conway_utils"
	"conway-v2/terminal_utils"
	"fmt"
)

func main() {

	terminal_utils.TerminalClear()

	// Hide cursor
	fmt.Print("\033c") // Rest terminal
	fmt.Print("\033[?25l")

	rows := 5
	cols := 5
	conway_grid := conway_utils.NewConwayGrid(rows, cols, 10)
	conway_grid.Run()
}
