package actor

import (
	tl "github.com/JoelOtter/termloop"
)

// Actor is an entity with available Actions.
type Actor interface {
	SetEntity(*tl.Entity)
	Inspect() []Action
}

// Action is a struct that fulfills actionable behavior.
type Action interface {
	GetActor() Actor
	String() string
	Do() error
}
