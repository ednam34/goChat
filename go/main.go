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
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var clients []*websocket.Conn

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		clients = append(clients, conn)

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				// Remove the disconnected client
				for i, client := range clients {
					if client == conn {
						clients = append(clients[:i], clients[i+1:]...)
						break
					}
				}
				return
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			var msgToSend Message

			err = json.Unmarshal(msg, &msgToSend)
			if err != nil {
				fmt.Println(err)
				continue
			}

			actualCl := conn.RemoteAddr().String()

			for _, client := range clients {
				// Write message back to browser
				if client.RemoteAddr().String() != actualCl {
					if err = client.WriteMessage(msgType, msg); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}
