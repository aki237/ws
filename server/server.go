package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var p []byte = []byte(`{ "jsonrpc": "2.0", "id": 1, "method": "call", "params": [ "00000000000000000000000000000000", "session", "login", { "username": "root", "password": "hello"  } ] }`)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		err := conn.WriteMessage(websocket.TextMessage, p)
		if err != nil {
			log.Println(err)
			return
		}
		_, mess, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Message : ", string(mess), " Error : ", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
