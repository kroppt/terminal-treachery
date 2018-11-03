package actor

import (
	tl "github.com/JoelOtter/termloop"
)

// Actor type that has the entity and available actions
type Actor interface {
	SetEntity(*tl.Entity)
	Inspect() []Action
}

// Action type in order to fulfill generality of actor behavior
type Action interface {
	GetActor() Actor
	String() string
	Do() error
}
