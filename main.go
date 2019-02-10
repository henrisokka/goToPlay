// websockets.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type incomingEvent struct {
	c *websocket.Conn
	a action
}

func main() {
	fmt.Println("main started")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		registerClient(conn, "1")

		for {
			// Read message from browser

			_, msg, err := conn.ReadMessage()

			if err != nil {
				return
			}

			a := action{}
			json.Unmarshal(msg, &a)

			//ie := incomingEvent{conn, a}
			fmt.Println("Message received")

			handleMessage(conn, a)
			//actionHandler(ie)

			// Write message back to browser
			/*if err = conn.WriteJSON(string(jsMsg)); err != nil {
				return
			}
			*/
		}
	})

	http.Handle("/", http.FileServer(http.Dir("./front")))

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening at port 8080")

}

func jsonHandler() {
	sound := &soundInfo{
		Vel:    20,
		Freq:   440,
		Length: 30,
	}
	jsMsg, _ := json.Marshal(sound)
	fmt.Println(string(jsMsg))
}
