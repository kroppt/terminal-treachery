package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

func main() {
	fmt.Println("test")
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	game.Screen().SetLevel(level)
	game.Start()
}
