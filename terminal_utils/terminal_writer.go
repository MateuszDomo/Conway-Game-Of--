package terminal_utils

import (
	"fmt"
	"golang.org/x/term"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

func EnableDebug() {
	fmt.Print("\033c")  // Rest terminal
	fmt.Print("\033[H") // Send cursor to top left
}

type TerminalDimensions struct {
	width  int
	height int
}

func TerminalGetDimensions() TerminalDimensions {
	fd := int(os.Stdout.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		EnableDebug()
		fmt.Println("Error: ", err)
		return TerminalDimensions{}
	}

	return TerminalDimensions{
		width,
		height,
	}
}

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

var colors = [...]TerminalColor{Red, Green, Yellow, Blue, Magenta, Cyan}

func TerminalRandomColor() TerminalColor {
	return colors[rand.Intn(len(colors))]
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
