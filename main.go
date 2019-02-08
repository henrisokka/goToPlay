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

type soundInfo struct {
	Vel    int
	Freq   int
	Length int
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
			sound := &soundInfo{
				Vel:    20,
				Freq:   440,
				Length: 30,
			}
			jsMsg, err := json.Marshal(sound)
			fmt.Println(string(jsMsg))

			msgType, msg, err := conn.ReadMessage()
			fmt.Println("msgType: ", msgType)
			if err != nil {
				return
			}

			if string(msg) == "test" {
				fmt.Println("You send test message")
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteJSON(string(jsMsg)); err != nil {
				return
			}
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
