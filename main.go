// websockets.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("main started")
	newHub()

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		newClient(1)
		fmt.Println(r)
		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			if string(msg) == "test" {
				fmt.Println("You send test message")
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, music()); err != nil {
				return
			}
		}
	})
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./front/index.html")
		})
	*/
	http.Handle("/", http.FileServer(http.Dir("./front")))

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening at port 8080")
}

func music() []byte {
	// here we mimic the music information
	return []byte("music mister BBoy")
}
