package server

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/DiscoViking/rogue/shared"
)

var connections map[string]net.Conn

func start(port string, rcv chan *shared.Message) {
	connections = map[string]net.Conn{}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		// Accept connection and handle.
		c, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleConnection(c, rcv)
	}
}

func handleConnection(c net.Conn, rcv chan *shared.Message) {
	defer c.Close()

	buf := make([]byte, 1024)
	message := shared.Message{}
	for {
		// Read the next message.
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Print(err)
		}
		data := buf[:n]

		// Unmarshal into the message struct.
		err = json.Unmarshal(data, &message)
		if err != nil {
			log.Print(err)
		}

		// If this connection isn't the one we expect messages
		// from this player on, error out.
		if match, ok := connections[message.Name]; ok {
			if match != c {
				log.Printf("Received message for player %v on wrong connection\n", message.Name)
				continue
			}
		} else {
			// Put us in the player connection list.
			connections[message.Name] = c
		}

		// Send message on to core.
		rcv <- &message
	}
}

func send(m *shared.Message) {
	buf, err := json.Marshal(m)
	if err != nil {
		log.Print(err)
	}

	conn, ok := connections[m.Name]
	if !ok {
		log.Printf("No connection to player %v\n", m.Name)
	}

	_, err = conn.Write(buf)
	if err != nil {
		log.Print(err)
	}
}
