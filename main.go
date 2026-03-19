package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/143life/snake-game/utils"
	"golang.org/x/term"
)

func main() {
	fmt.Print("\033[1;31mThis is my first terminal util\033[0m\n")

	// raw mode to our terminal
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Print("error while putting terminal into raw mode")
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)
	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			fmt.Println("error while reading rune from stdin")
		}
		if r == 'q' {
			return
		}
		// 001 lab
		utils.ColorByCountByteInRune(r, size)
	}
}
