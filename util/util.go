package util

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

// FatalLog closes the game out immediately and prints the given error.
func FatalLog(e error) {
	termbox.Flush()
	termbox.Close()
	log.Fatalln(e)
}
