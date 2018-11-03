package main

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/kroppt/terminal-treachery/actor"
	"github.com/kroppt/terminal-treachery/cfg"
	"github.com/kroppt/terminal-treachery/terrain"
	"github.com/nsf/termbox-go"
)

func main() {
	// Remove base colors and bold.
	termbox.SetOutputMode(termbox.Output216)
	game := tl.NewGame()

	player := actor.NewPlayer(tl.NewEntity(1, 1, 1, 1))

	off := 1 // Default color is 0.
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.Attr(180 + off),
		Fg: tl.Attr(137 + off),
		Ch: 'v',
	})
	rect := tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue)
	w := &terrain.Wall{Rectangle: rect}
	level.AddEntity(w)

	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'o'})
	player.AddLevel(level)
	level.AddEntity(player)

	conf := cfg.GetConfig()

	game.Screen().SetFps(conf.FPS)
	game.Screen().SetLevel(level)
	game.Start()
}
