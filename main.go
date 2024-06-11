package main

import (
	"conway-v2/terminal_utils"
	"fmt"
	"time"
)

func main() {

	terminal_utils.TerminalClear()

	// Hide cursor
	fmt.Print("\033[?25l")

	game := NewConwayGame(20, 50, 0, 0, terminal_utils.TerminalRandomColor())
	game.RandomlyPopulate()

	for {
		game.PlayCycle()
		time.Sleep(85 * time.Millisecond)
	}
}
