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

	id := string(len(rooms[currentRoomIndex].clients)) + "client"

	cl := Client{conn, id, roomId}

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
}

func handleMessage(e Event) {
	Actions = append(Actions, e.action)

	var sender Client
	//var room *Room
	var roomIndex int

	for i, rm := range rooms {
		for _, cl := range rm.clients {
			if cl.conn == e.conn {
				//room = &rm
				roomIndex = i
				sender = cl
			}
		}
	}

	rooms[roomIndex].sendActionToOtherClients(e.action, sender.conn)
	/*
		rooms[roomIndex].state.handleAction(e.action, sender.id)
		room.sendStateToClients(sender.conn)
	*/
}

func (r *Room) sendActionToOtherClients(a Action, sender *websocket.Conn) {
	for _, cl := range r.clients {
		if cl.conn != sender {
			stateJSON, _ := json.Marshal(a.Sound)
			if err := cl.conn.WriteJSON(string(stateJSON)); err != nil {
				fmt.Println("Error in JSONing: ", err)
			}
		}
	}
}

func (r *Room) sendStateToClients(sender *websocket.Conn) {
	fmt.Println("Send state to clients")
	for _, cl := range r.clients {
		if cl.conn != sender {
			stateJSON, _ := json.Marshal(r.state)
			if err := cl.conn.WriteJSON(string(stateJSON)); err != nil {
				fmt.Println("Error in JSONing: ", err)
			}
		}
	}
}
