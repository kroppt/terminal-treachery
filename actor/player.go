package actor

import (
	tl "github.com/JoelOtter/termloop"
)

// Player is the user-controlled player
type Player struct {
	*tl.Entity
	Health    int32
	MaxHealth int32
	prevX     int
	prevY     int
	level     *tl.BaseLevel
}

// AddLevel adds the given level to the player for position centering
func (p *Player) AddLevel(l *tl.BaseLevel) {
	p.level = l
}

// Tick executes one game loop
func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		p.prevX, p.prevY = p.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			p.SetPosition(p.prevX+1, p.prevY)
		case tl.KeyArrowLeft:
			p.SetPosition(p.prevX-1, p.prevY)
		case tl.KeyArrowUp:
			p.SetPosition(p.prevX, p.prevY-1)
		case tl.KeyArrowDown:
			p.SetPosition(p.prevX, p.prevY+1)
		}
	}
}

// Collide stops the player in a collision.
func (p *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		p.SetPosition(p.prevX, p.prevY)
	}
}

// Draw the screen where the player is centered.
func (p *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := p.Position()
	p.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	p.Entity.Draw(screen)
}

// NewPlayer returns a pointer to the newly created player with the given entity
func NewPlayer(e *tl.Entity) *Player {
	var p Player
	p.SetEntity(e)
	return &p
}

// GetEntity returns the entity of the player
func (p *Player) GetEntity() *tl.Entity {
	return p.Entity
}

// SetEntity sets the entity of the player to the given entity
func (p *Player) SetEntity(e *tl.Entity) {
	p.Entity = e
}

// Heal increases the player's health
func (p *Player) Heal(h int32) {
	if h == 0 {
		p.Health = p.MaxHealth
		return
	}
	p.Health += h
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
}

// Hit lowers the player's health
func (p *Player) Hit(d int32) {
	p.Health -= d
	if p.Health < 0 {
		p.Health = 0
	}
}

// Inspect returns a list of available actions to take against the player
func (p *Player) Inspect() []string {
	return []string{"heal", "hit"}
}
