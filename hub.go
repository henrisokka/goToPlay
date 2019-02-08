package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	id   string
}

type Hub struct {
	clients []Client
}

var activeHub = Hub{
	make([]Client, 0),
}

func (h *Hub) registerClient(cl Client) {
	h.clients = append(h.clients, cl)
	cl.conn.WriteMessage(1, []byte("Sinut on nyt rekisteröity"))

	for _, c := range h.clients {
		if c != cl {
			c.conn.WriteMessage(1, []byte("Joku rekisteröityi!"))
		}
	}
	fmt.Println("New client registered:")
	fmt.Println(h.clients)
}

func newHub() {

	fmt.Println("Hub registering?")
}

func newClient(id string, conn *websocket.Conn) {
	c := Client{conn, id}
	activeHub.registerClient(c)
}

func updateClients(s state, conn *websocket.Conn) {

}
