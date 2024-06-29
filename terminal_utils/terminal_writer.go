package terminal_utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/term"
)

func EnableDebug() {
	fmt.Print("\033c")  // Rest terminal
	fmt.Print("\033[H") // Send cursor to top left
}

func DebugPrint(row int, col int, val string) {
	// 1 based
	row_s := strconv.Itoa(row + 1)
	col_s := strconv.Itoa(col + 2)

	fmt.Print("\033[" + row_s + ";" + col_s + "H")
	fmt.Print(val)
}

type TerminalDimensions struct {
	Width  int
	Height int
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

	fmt.Print("\033c")     // Rest terminal
	fmt.Print("\033[?25l") // Hide cursor
}

func TerminalSetForegroundColor(color TerminalColor) {
	fmt.Print(color)
}

func TerminalSetBackgroundColor(active bool) {
	if active {
		fmt.Printf("\033[48;2;220;150;150m")
	} else {
		fmt.Printf("\033[49m")
	}
}

type TerminalArgs struct {
	Rows int
	Cols int
}

func TerminalExtractArgs(args []string) (TerminalArgs, error) {
	if len(args) != 3 {
		return TerminalArgs{}, errors.New("Incorrect arguments: Format => rows cols")
	}

	rows, err := strconv.Atoi(args[1])
	if err != nil {
		return TerminalArgs{}, errors.New("Incorrect arguments. rows must be an integer")
	}

	cols, err := strconv.Atoi(args[2])
	if err != nil {
		return TerminalArgs{}, errors.New("Incorrect arguments. cols must be an integer")
	}

	return TerminalArgs{rows, cols}, nil
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
