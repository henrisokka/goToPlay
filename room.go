package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	id   string
	room Room
}

type Room struct {
	clients []Client
	id      string
}

var Actions = make([]action, 0)

var rooms = make([]Room, 0)

func registerClient(conn *websocket.Conn, roomId string) {
	var currentRoom *Room

	var currentRoomIndex int

	if len(rooms) > 0 {
		for _, rm := range rooms {
			if rm.id == roomId {
				currentRoom = &rm
			}
		}
	}

	if currentRoom == nil {
		fmt.Println("We need a new room")
		currentRoomIndex = newRoom(roomId)
		currentRoom = &rooms[currentRoomIndex]
	}
	fmt.Println("rooms:")
	fmt.Println(rooms)

	cl := Client{conn, "foo", *currentRoom}

	fmt.Println("currentRoomIndex: ", currentRoomIndex)

	rooms[currentRoomIndex].addClientToRoom(cl)

	currentRoom.id = "BOO"
	fmt.Println(currentRoom)
	fmt.Println(rooms)
}

func newRoom(id string) int {
	room := Room{
		clients: make([]Client, 0),
		id:      id,
	}
	rooms = append(rooms, room)

	fmt.Println("New room created")
	fmt.Println(len(rooms) - 1)

	return len(rooms) - 1
}

func (r *Room) addClientToRoom(cl Client) {
	fmt.Println("add client to room")
	r.clients = append(r.clients, cl)
	cl.conn.WriteMessage(1, []byte("You are registered to the room"))
}

func handleMessage(conn *websocket.Conn, message action) {
	fmt.Println("handleMessage")
	fmt.Println(message)
	fmt.Println(len(rooms))
	Actions = append(Actions, message)

	//update all the other sockets
	var sender Client
	var room Room

	for _, rm := range rooms {
		fmt.Println(rm.clients)
		for _, cl := range rm.clients {
			fmt.Println(cl.conn == conn)
			if cl.conn == conn {
				fmt.Println("We have find what we are looking for")
				room = rm
				sender = cl
			}
		}
	}

	for _, cl := range room.clients {
		fmt.Println(sender)
		fmt.Println(cl)
	}

}

func (r *Room) sendStateToClients() {

}
