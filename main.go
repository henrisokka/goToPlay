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
	newHub()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		newClient("0", conn)
		//fmt.Println(conn)
		for {
			// Read message from browser

			msgType, msg, err := conn.ReadMessage()
			fmt.Println("msgType: ", msgType)
			if err != nil {
				return
			}

			fmt.Println(msg)

			a := action{}
			json.Unmarshal(msg, &a)

			ie := incomingEvent{conn, a}
			actionHandler(ie)

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

func music() []byte {
	// here we mimic the music information
	return []byte("music mister BBoy")
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
