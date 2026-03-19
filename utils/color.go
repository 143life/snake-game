package utils

import (
	"fmt"
)

func ColorByCountByteInRune(r rune, size int) {
	// color depends on the number of bytes in rune
	switch size {
	case 2:
		fmt.Printf("\033[1;31m%c, %d\033[0m\n", r, r)
	case 1:
		fmt.Printf("\033[1;34m%c, %d\033[0m\n", r, r)
	default:
		fmt.Printf("%c, %d\n", r, r)
	}

}
