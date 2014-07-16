package server

import "github.com/DiscoViking/rogue/shared"

type playerState int

const (
	player_idle = iota
	player_moved
)

type player struct {
	shared.Player
	state playerState
}

func (p *player) receive(m *shared.Message) {
	switch m.Type {
	case "MOVE":
		p.X++
		m.Type = "MOVE OK"
		m.X = p.X
		m.Y = p.Y
		send(m)
	}
}

func newPlayer(name string) *player {
	return &player{shared.Player{}, player_idle}
}
