package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Message struct {
	Message string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: upgrading to websocket connection: %v", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	defer ws.Close()

	log.Println("Client connected")

	for {
		var msg Message
		// Read message as JSON
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: reading websocket message: %v", err)
			break
		}
		log.Printf("Received: %s", msg.Message)

		// Write message back as JSON
		err = ws.WriteJSON(msg)

		if err != nil {
			log.Printf("error: writing websocket message: %v", err)
			break
		}
	}

	log.Println("Client disconnected")
}
