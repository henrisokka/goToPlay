package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	id string
	room Room
}

type Room struct {
	clients []Client
	id string
}

var rooms = make([]Room, 0)

func newRoom(id string) *Room {
	room := Room{
		clients: make([]Client, 0),
		id: id,
	}
	rooms = append(rooms, room)

	fmt.Println("New room created")

	return &room
}

func (r *Room) addClientToRoom(cl Client) {
	r.clients = append(r.clients, cl)
	cl.conn.WriteMessage(1, []byte("You are registered to the room"))
}

func registerClient(conn *websocket.Conn, roomId string) {
	var currentRoom *Room
	
	if len(rooms) > 0 {
		for _, rm := range rooms {
			if rm.id == roomId {
				currentRoom = &rm
			}
		}
	}

	if currentRoom == nil {
		fmt.Println("We need a new room")
		currentRoom = newRoom(roomId)
	}
	

	cl := Client{conn, "foo", *currentRoom}

	currentRoom.addClientToRoom(cl)
}

func (r *Room) sendStateToClients() {

}