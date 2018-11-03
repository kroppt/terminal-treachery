package actor

import (
	"reflect"
	"testing"

	tl "github.com/JoelOtter/termloop"
)

func TestHealAction_GetActor(t *testing.T) {
	act := NewPlayer(tl.NewEntity(1, 1, 1, 1))
	nilAct := &HealAction{}
	tests := []struct {
		name string
		a    *HealAction
		want Actor
	}{
		{
			name: "normal use",
			a:    &HealAction{Player: act},
			want: act,
		},
		{
			name: "nil actor",
			a:    nilAct,
			want: nilAct.Player,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetActor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HealAction.GetActor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealAction_String(t *testing.T) {
	tests := []struct {
		name string
		a    *HealAction
		want string
	}{
		{
			name: "normal use",
			a:    &HealAction{Player: NewPlayer(tl.NewEntity(1, 1, 1, 1))},
			want: "heal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("HealAction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealAction_Heal(t *testing.T) {
	type args struct {
		h int32
	}
	before := HealAction{Player: NewPlayer(tl.NewEntity(1, 1, 1, 1))}
	after := before
	after.h = 50
	after.ready = true
	tests := []struct {
		name string
		a    *HealAction
		args args
		want *HealAction
	}{
		{
			name: "normal use",
			a:    &before,
			args: args{50},
			want: &after,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Heal(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HealAction.Heal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHealAction_Do(t *testing.T) {
	tests := []struct {
		name    string
		a       *HealAction
		wantErr bool
	}{
		{
			name:    "normal use",
			a:       &HealAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 50, true},
			wantErr: false,
		},
		{
			name:    "preemptive use",
			a:       &HealAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 0, false},
			wantErr: true,
		},
		{
			name:    "max restore",
			a:       &HealAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 0, true},
			wantErr: false,
		},
		{
			name:    "at max health",
			a:       &HealAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 5, true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Do(); (err != nil) != tt.wantErr {
				t.Errorf("HealAction.Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHitAction_GetActor(t *testing.T) {
	act := NewPlayer(tl.NewEntity(1, 1, 1, 1))
	nilAct := &HitAction{}
	tests := []struct {
		name string
		a    *HitAction
		want Actor
	}{
		{
			name: "normal use",
			a:    &HitAction{Player: act},
			want: act,
		},
		{
			name: "nil actor",
			a:    nilAct,
			want: nilAct.Player,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetActor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HitAction.GetActor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHitAction_String(t *testing.T) {
	tests := []struct {
		name string
		a    *HitAction
		want string
	}{
		{
			name: "normal use",
			a:    &HitAction{Player: NewPlayer(tl.NewEntity(1, 1, 1, 1))},
			want: "hit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("HitAction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHitAction_Hit(t *testing.T) {
	before := HitAction{Player: NewPlayer(tl.NewEntity(1, 1, 1, 1))}
	after := before
	after.d = 50
	after.ready = true
	type args struct {
		d int32
	}
	tests := []struct {
		name string
		a    *HitAction
		args args
		want *HitAction
	}{
		{
			name: "normal use",
			a:    &before,
			args: args{50},
			want: &after,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Hit(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HitAction.Hit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHitAction_Do(t *testing.T) {
	tests := []struct {
		name    string
		a       *HitAction
		wantErr bool
	}{
		{
			name:    "normal use",
			a:       &HitAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 50, true},
			wantErr: false,
		},
		{
			name:    "preemptive use",
			a:       &HitAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 0, false},
			wantErr: true,
		},
		{
			name:    "min health",
			a:       &HitAction{NewPlayer(tl.NewEntity(1, 1, 1, 1)), 0, true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Do(); (err != nil) != tt.wantErr {
				t.Errorf("HitAction.Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlayer_Entity(t *testing.T) {
	entity := tl.NewEntity(1, 1, 1, 1)
	tests := []struct {
		name string
		p    *Player
		want *tl.Entity
	}{
		{
			name: "normal use",
			p:    NewPlayer(entity),
			want: entity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Entity; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Player.Entity = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_SetEntity(t *testing.T) {
	entity := tl.NewEntity(1, 1, 1, 1)
	type args struct {
		e *tl.Entity
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		{
			name: "normal use",
			p:    &Player{},
			args: args{entity},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.SetEntity(tt.args.e)
		})
	}
}

func TestPlayer_Inspect(t *testing.T) {
	p := NewPlayer(tl.NewEntity(1, 1, 1, 1))
	tests := []struct {
		name string
		p    *Player
		want []Action
	}{
		{
			name: "normal use",
			p:    p,
			want: []Action{&HitAction{Player: p}, &HealAction{Player: p}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Inspect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Player.Inspect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_AddLevel(t *testing.T) {
	nl := tl.NewBaseLevel(tl.Cell{})
	type args struct {
		l *tl.BaseLevel
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		{
			name: "normal use",
			p:    &Player{},
			args: args{nl},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.AddLevel(tt.args.l)
		})
	}
}

func TestPlayer_Tick(t *testing.T) {
	player := NewPlayer(tl.NewEntity(1, 1, 1, 1))
	type args struct {
		event tl.Event
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		{
			name: "non-keyboard event",
			p:    player,
			args: args{tl.Event{Type: tl.EventInterrupt}},
		},
		{
			name: "right-arrow event",
			p:    player,
			args: args{tl.Event{Type: tl.EventKey, Key: tl.KeyArrowRight}},
		},
		{
			name: "left-arrow event",
			p:    player,
			args: args{tl.Event{Type: tl.EventKey, Key: tl.KeyArrowLeft}},
		},
		{
			name: "up-arrow event",
			p:    player,
			args: args{tl.Event{Type: tl.EventKey, Key: tl.KeyArrowUp}},
		},
		{
			name: "down-arrow event",
			p:    player,
			args: args{tl.Event{Type: tl.EventKey, Key: tl.KeyArrowDown}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Tick(tt.args.event)
		})
	}
}

func TestPlayer_Collide(t *testing.T) {
	rect := tl.NewRectangle(0, 0, 1, 1, 0)
	type args struct {
		collision tl.Physical
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		{
			name: "normal use",
			p:    NewPlayer(tl.NewEntity(1, 1, 1, 1)),
			args: args{rect},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Collide(tt.args.collision)
		})
	}
}

func TestPlayer_Draw(t *testing.T) {
	game := tl.NewGame()
	bl := tl.NewBaseLevel(tl.Cell{})
	type args struct {
		screen *tl.Screen
	}
	tests := []struct {
		name string
		p    *Player
		args args
	}{
		{
			name: "normal use",
			p:    &Player{Entity: tl.NewEntity(1, 1, 1, 1), level: bl},
			args: args{game.Screen()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Draw(tt.args.screen)
		})
	}
}

func TestNewPlayer(t *testing.T) {
	entity := tl.NewEntity(1, 1, 1, 1)
	type args struct {
		e *tl.Entity
	}
	tests := []struct {
		name string
		args args
		want *Player
	}{
		{
			name: "normal use",
			args: args{entity},
			want: &Player{Entity: entity},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlayer(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
