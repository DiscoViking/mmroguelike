package server

import "github.com/DiscoViking/rogue/shared"

var players map[string]*player

func Run(port string) {
	rcv := make(chan *shared.Message, 10)
	players = map[string]*player{}

	go start(port, rcv)

	for m := range rcv {
		receive(m)
	}
}

func receive(m *shared.Message) {
	if player, ok := players[m.Name]; ok {
		player.receive(m)
	} else {
		if m.Type == "LOGIN" {
			players[m.Name] = newPlayer(m.Name)
			m.Type = "LOGIN OK"
			send(m)
		}
	}
}
