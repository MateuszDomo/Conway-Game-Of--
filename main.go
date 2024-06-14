package main

import (
	"conway-v2/terminal_utils"
	"fmt"
)

func main() {

	terminal_utils.TerminalClear()

	// Hide cursor
	fmt.Print("\033[?25l")

	conway_grid := NewConwayGrid(1, 3, 10)

	for {
		conway_grid.Run()
	}
}
