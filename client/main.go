package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/DiscoViking/rogue/shared"
)

func Run(server string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		dec := json.NewDecoder(conn)
		message := &shared.Message{}
		for {
			err := dec.Decode(message)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%v\n", message)
		}
	}()

	m := &shared.Message{
		Id:   1,
		Name: "Player1",
		Type: "LOGIN",
	}

	enc := json.NewEncoder(conn)

	enc.Encode(m)

	m = &shared.Message{
		Id:   1,
		Name: "Player1",
		Type: "MOVE",
	}

	enc.Encode(m)

	time.Sleep(1 * time.Second)
}
