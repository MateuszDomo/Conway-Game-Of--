package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func TerminalWrite(row int, col int, val string) {
	// 1 based
	row_s := strconv.Itoa(row + 1)
	col_s := strconv.Itoa(col + 1)

	// Move terminal cursor to row and column using ANSI escape sequences
	fmt.Print("\033[" + row_s + ";" + col_s + "H")
	fmt.Print(val)
}

func TerminalClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
