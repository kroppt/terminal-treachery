package actor

import (
	tl "github.com/JoelOtter/termloop"
)

// Actor type that has the entity and available actions
type Actor interface {
	GetEntity() *tl.Entity
	SetEntity(*tl.Entity)
	Inspect() []string
}
