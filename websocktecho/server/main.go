package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "Connection closed")

	for {
		messageType, data, err := conn.Read(r.Context())
		if websocket.CloseStatus(err) != -1 {
			log.Println("Connection closed:", err)
			return
		}
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Println("Received message:", string(data))

		err = conn.Write(r.Context(), messageType, data)
		if err != nil {
			log.Println("Error sending response:", err)
			return
		}

	}
}
func main() {
	http.HandleFunc("/ws", echoHandler)
	fmt.Println("Echo Server started on ws://localhost:8080/ws")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
