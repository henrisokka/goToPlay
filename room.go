package main

import (
	"encoding/json"
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
}

var Actions = make([]action, 0)

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
	fmt.Println("Rooms after client add")
	fmt.Println(rooms)

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

func handleMessage(conn *websocket.Conn, message action) {
	fmt.Println("handleMessage")
	fmt.Println(message)
	Actions = append(Actions, message)

	//update all the other sockets
	var sender Client
	var room Room

	fmt.Println("ranging rooms:")
	fmt.Println(rooms)

	for _, rm := range rooms {
		for _, cl := range rm.clients {
			if cl.conn == conn {
				fmt.Println("We have find what we are looking for")
				room = rm
				sender = cl
			}
		}
	}

	for _, cl := range room.clients {
		if cl != sender {
			json, _ := json.Marshal(message)
			cl.conn.WriteMessage(1, []byte("Someone sent a message to the room"))
			cl.conn.WriteJSON(string(json))
		}
	}

}

func (r *Room) sendStateToClients() {

}
