package terminal_utils

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

func TerminalSetColor(color TerminalColor) {
	fmt.Println(color)
}

type TerminalColor string

const (
	Red     TerminalColor = "\033[31m"
	Green   TerminalColor = "\033[32m"
	Yellow  TerminalColor = "\033[33m"
	Blue    TerminalColor = "\033[34m"
	Magenta TerminalColor = "\033[35m"
	Cyan    TerminalColor = "\033[36m"
	Reset   TerminalColor = "\033[0m"
)
