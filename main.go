package main

import (
	"conway-v2/conway_utils"
	"conway-v2/terminal_utils"
	"fmt"
)

func main() {

	terminal_utils.TerminalClear()

	// Hide cursor
	fmt.Print("\033[?25l")

	conway_grid := conway_utils.NewConwayGrid(5, 5, 10)

	conway_grid.Run()
}
