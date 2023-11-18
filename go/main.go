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

type User struct {
	client   *websocket.Conn
	Username string
}

type Message struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var users []User

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		var user User
		user.client = conn
		user.Username = ""
		users = append(users, user)

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				var quitUser string
				for i, user := range users {
					if user.client == conn {
						users = append(users[:i], users[i+1:]...)
						quitUser = user.Username
						break
					}
				}
				quitMsg := Message{
					Type:     "error",
					Username: quitUser,
					Message:  "Viens de se déco",
				}
				messageJSON, err := json.Marshal(quitMsg)
				if err != nil {
					fmt.Println(err)
					return
				}
				for _, user := range users {
					if err = user.client.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
						fmt.Println(err)
						return
					}
				}
				return
			}

			//fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			var msgToSend Message

			err = json.Unmarshal(msg, &msgToSend)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Message Type %s from %s\nmessage content : %s \n", msgToSend.Type, msgToSend.Username, msgToSend.Message)
			if msgToSend.Type == "success" {
				for i, user := range users {
					if user.client == conn {
						users[i].Username = msgToSend.Username
						fmt.Print(users)
						break
					}
				}
			}
			actualCl := conn.RemoteAddr().String()

			for _, user := range users {
				// Write message back to browser
				if user.client.RemoteAddr().String() != actualCl {
					if err = user.client.WriteMessage(msgType, msg); err != nil {
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

	http.ListenAndServe(":8081", nil)
}
