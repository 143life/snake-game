package main

import (
	"github.com/nsf/termbox-go"
)

type Coordinate struct {
	x int
	y int
}

type Snake struct {
	length    int
	direction termbox.Key
	body      []Coordinate
}

func main() {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Set a cell in the back buffer
	for i := 1; i <= 60; i++ {
		termbox.SetCell(i, 1, '-', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(i, 30, '-', termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i <= 30; i++ {
		termbox.SetCell(1, i, '|', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(60, i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	// initialize snake
	snake := Snake{length: 3, direction: termbox.KeyArrowRight}
	for i := 2; i < 2+snake.length; i++ {
		termbox.SetBg(i, 2, termbox.ColorGreen)
	}

	// Flush the back buffer to the terminal
	termbox.Flush()

	// Wait for an ESC key press to exit
	for {
		switch snake.direction {
		case termbox.KeyArrowRight:

		case termbox.KeyArrowLeft:

		case termbox.KeyArrowDown:

		case termbox.KeyArrowUp:

		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// Handle keyboard input
			if ev.Key == termbox.KeyEsc {
				return // Exit on ESC key
			}
		}
	}
}
