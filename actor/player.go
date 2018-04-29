package actor

import (
	"errors"

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

// HealAction is the action for healing
type HealAction struct {
	*Player
	h     int32
	ready bool
}

// HitAction is the action for being hit
type HitAction struct {
	*Player
	d     int32
	ready bool
}

// GetActor returns the actor associated with the action
func (a *HealAction) GetActor() Actor {
	return a.Player
}

// GetActor returns the actor associated with the action
func (a *HealAction) String() string {
	return "heal"
}

// Heal prepares the heal action
func (a *HealAction) Heal(h int32) {
	a.h = h
	a.ready = true
}

// Do executes the heal action
func (a *HealAction) Do() error {
	if !a.ready {
		return errors.New("struct HealAction not ready, use HealAction.Heal(int32) first")
	}
	if a.h == 0 {
		// 0 HP is code for max heal
		a.Health = a.MaxHealth
	}
	a.Health += a.h
	if a.Health > a.MaxHealth {
		// Wasted health points
		a.Health = a.MaxHealth
	}
	return nil
}

// GetActor returns the actor associated with the action
func (a *HitAction) GetActor() Actor {
	return a.Player
}

// GetActor returns the actor associated with the action
func (a *HitAction) String() string {
	return "hit"
}

// Hit lowers the player's health
func (a *HitAction) Hit(d int32) {
	a.d = d
	a.ready = true
}

// Do executes the hit action on its actor
func (a *HitAction) Do() error {
	if !a.ready {
		return errors.New("struct HitAction not ready, use HitAction.Hit(int32) first")
	}
	a.Health -= a.d
	if a.Health < 0 {
		a.Health = 0
	}
	return nil
}

// GetEntity returns the entity of the player
func (p *Player) GetEntity() *tl.Entity {
	return p.Entity
}

// SetEntity sets the entity of the player to the given entity
func (p *Player) SetEntity(e *tl.Entity) {
	p.Entity = e
}

// Inspect returns a list of available actions to take against the player
func (p *Player) Inspect() []Action {
	hit := &HitAction{p, 0, false}
	heal := &HealAction{p, 0, false}
	return []Action{hit, heal}
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
