package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	id     string
	roomId string
}

type Room struct {
	clients []Client
	id      string
	state   State
}

var Actions = make([]Action, 0)

var rooms = make([]Room, 0)

func registerClient(conn *websocket.Conn, roomId string) {
	fmt.Println("registerClient")
	var currentRoom *Room

	var currentRoomIndex int

	if len(rooms) > 0 {
		for i, rm := range rooms {
			if rm.id == roomId {
				currentRoom = &rm
				currentRoomIndex = i
			}
		}
	}

	if currentRoom == nil {
		currentRoomIndex = newRoom(roomId)
		currentRoom = &rooms[currentRoomIndex]
	}

	cl := Client{conn, "foo", roomId}

	rooms[currentRoomIndex].addClientToRoom(cl)
}

func newRoom(id string) int {
	room := Room{
		clients: make([]Client, 0),
		id:      id,
	}
	rooms = append(rooms, room)

	return len(rooms) - 1
}

func (r *Room) addClientToRoom(cl Client) {
	r.clients = append(r.clients, cl)
	cl.conn.WriteMessage(1, []byte("You are registered to the room"))
}

func handleMessage(e Event) {
	fmt.Println("handleMessage")
	fmt.Println(e.action)
	Actions = append(Actions, e.action)

	//update all the other sockets
	var sender Client
	var room *Room
	var roomIndex int

	for _, rm := range rooms {
		for i, cl := range rm.clients {
			if cl.conn == e.conn {
				fmt.Println("We have find what we are looking for")
				room = &rm
				roomIndex = i
				sender = cl
			}
		}
	}
	rooms[roomIndex].state.handleAction(e.action)
	room.sendStateToClients(sender.conn)
}

func (r *Room) sendStateToClients(sender *websocket.Conn) {
	for _, cl := range r.clients {
		if cl.conn != sender {
			cl.conn.WriteMessage(1, []byte("We should send you the new state here, just wait a minute!"))
		}
	}
}
