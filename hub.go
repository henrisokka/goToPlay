package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients []*websocket.Conn
}

var activeHub Hub = Hub{
	make([]*websocket.Conn, 0),
}

func (h *Hub) registerClient(conn *websocket.Conn) {
	h.clients = append(h.clients, conn)
	conn.WriteMessage(1, []byte("Sinut on nyt rekisteröity"))

	for _, c := range h.clients {
		if c != conn {
			c.WriteMessage(1, []byte("Joku rekisteröityi!"))
		}
	}
	fmt.Println("New client registered:")
	fmt.Println(h.clients)
}

func newHub() {

	fmt.Println("Hub registering?")
}

func newClient(id int, conn *websocket.Conn) {
	activeHub.registerClient(conn)
}
