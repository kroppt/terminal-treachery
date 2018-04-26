package main

import (
	tl "github.com/JoelOtter/termloop"
	cfg "github.com/kroppt/terminal-treachery/cfg"
	"github.com/nsf/termbox-go"
)

// Player represents entity information for the player.
// Player holds Entity info, previous coordinates.
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

// Tick executes one game loop
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

// Collide stops the player in a collision.
func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

// Draw the screen where the player is centered.
func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}

func main() {
	// remove base colors and bold
	termbox.SetOutputMode(termbox.Output216)
	game := tl.NewGame()

	off := 1 // default color is 0
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.Attr(180 + off),
		Fg: tl.Attr(137 + off),
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)

	conf := cfg.GetConfig()

	game.Screen().SetFps(conf.FPS)
	game.Screen().SetLevel(level)
	game.Start()
}
