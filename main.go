package main

import (
	"conway-v2/conway_utils"
	"conway-v2/terminal_utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	terminal_utils.TerminalClear()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// Wait for a signal
		receivedSignal := <-sig
		terminal_utils.TerminalClear()
		fmt.Println("Terminating...\n", receivedSignal)

		os.Exit(0)
	}()

	args, err := terminal_utils.TerminalExtractArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := args.Rows
	cols := args.Cols

	conway_grid := conway_utils.NewConwayGrid(rows, cols, 10)
	conway_grid.Run()
}
